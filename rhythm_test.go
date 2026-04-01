package main

import (
	"container/ring"
	"fmt"
	"testing"
)

// ringToBools converts a flattened ring of bools to a slice for easy comparison.
func ringToBools(r *ring.Ring) []bool {
	if r == nil {
		return []bool{}
	}
	var result []bool
	r.Do(func(v any) {
		result = append(result, v.(bool))
	})
	return result
}

// euclidean calls rhythm(beats, steps-beats) and optionally rotates by shift.
func euclidean(beats, steps, shift int) []bool {
	r := rhythm(beats, steps-beats)
	if r == nil {
		return []bool{}
	}
	r = r.Move(shift)
	return ringToBools(r)
}

func countBeats(pattern []bool) int {
	n := 0
	for _, v := range pattern {
		if v {
			n++
		}
	}
	return n
}

// TestKnownPatterns checks classic Euclidean rhythms from the literature.
// Reference: Toussaint, "The Euclidean Algorithm Generates Traditional Musical Rhythms" (2005)
func TestKnownPatterns(t *testing.T) {
	tests := []struct {
		beats, steps int
		want         []bool
		name         string
	}{
		// E(1,2): [x .]
		{1, 2, []bool{true, false}, "E(1,2)"},
		// E(1,3): [x . .]
		{1, 3, []bool{true, false, false}, "E(1,3)"},
		// E(2,3): [x . x] -> actually [x x .] before rotate; canonical is [x . x]
		// Bjorklund without shift gives [x x .], shift by 1 gives [x . x]
		{2, 3, []bool{true, true, false}, "E(2,3) no shift"},
		// E(3,4): [x x x .]
		{3, 4, []bool{true, true, true, false}, "E(3,4) no shift"},
		// E(2,5): [x . x . .] — bossa nova / tresillo-adjacent
		{2, 5, []bool{true, false, true, false, false}, "E(2,5)"},
		// E(3,8): [x . . x . . x .] — Cuban tresillo
		{3, 8, []bool{true, false, false, true, false, false, true, false}, "E(3,8)"},
		// E(4,12): [x . . x . . x . . x . .] — standard 4/4 subdivision
		{4, 12, []bool{true, false, false, true, false, false, true, false, false, true, false, false}, "E(4,12)"},
		// E(5,8): [x . x x . x x .] — common 5-in-8 pattern
		{5, 8, []bool{true, false, true, true, false, true, true, false}, "E(5,8)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := euclidean(tt.beats, tt.steps, 0)
			if len(got) != len(tt.want) {
				t.Fatalf("length %d, want %d: got %v", len(got), len(tt.want), got)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("mismatch at index %d: got %v, want %v\n  full: %v", i, got[i], tt.want[i], got)
					break
				}
			}
		})
	}
}

// TestBeatCount verifies the output always contains exactly `beats` true values.
func TestBeatCount(t *testing.T) {
	cases := [][2]int{
		{1, 4}, {2, 4}, {3, 4},
		{3, 8}, {5, 8}, {7, 8},
		{4, 16}, {5, 16}, {11, 16},
		{1, 1}, {5, 5},
	}
	for _, c := range cases {
		beats, steps := c[0], c[1]
		got := euclidean(beats, steps, 0)
		if len(got) != steps {
			t.Errorf("E(%d,%d): length %d, want %d", beats, steps, len(got), steps)
			continue
		}
		n := countBeats(got)
		if n != beats {
			t.Errorf("E(%d,%d): got %d beats, want %d — pattern: %v", beats, steps, n, beats, got)
		}
	}
}

// TestEvenDistribution checks that beats are as evenly spaced as possible.
// For a valid Euclidean rhythm, the gap sizes between beats differ by at most 1.
func TestEvenDistribution(t *testing.T) {
	cases := [][2]int{
		{3, 8}, {5, 8}, {4, 12}, {7, 16}, {3, 7}, {5, 13},
	}
	for _, c := range cases {
		beats, steps := c[0], c[1]
		pattern := euclidean(beats, steps, 0)

		// Collect gap lengths between consecutive beats (wrapping around).
		var gaps []int
		for i, v := range pattern {
			if v {
				gap := 1
				for j := 1; j < steps; j++ {
					if pattern[(i+j)%steps] {
						break
					}
					gap++
				}
				gaps = append(gaps, gap)
			}
		}

		if len(gaps) != beats {
			t.Errorf("E(%d,%d): expected %d gaps, got %d", beats, steps, beats, len(gaps))
			continue
		}

		min, max := gaps[0], gaps[0]
		for _, g := range gaps {
			if g < min {
				min = g
			}
			if g > max {
				max = g
			}
		}
		if max-min > 1 {
			t.Errorf("E(%d,%d): gaps differ by %d (min=%d max=%d), want ≤1 — pattern: %v",
				beats, steps, max-min, min, max, pattern)
		}
	}
}

// TestAllBeats checks that rhythm(n, 0) produces all beats.
func TestAllBeats(t *testing.T) {
	for n := 1; n <= 8; n++ {
		got := euclidean(n, n, 0)
		if len(got) != n {
			t.Errorf("E(%d,%d): length %d, want %d", n, n, len(got), n)
			continue
		}
		for i, v := range got {
			if !v {
				t.Errorf("E(%d,%d): index %d is false, want all true", n, n, i)
			}
		}
	}
}

// BenchmarkRhythm measures performance at various sizes.
func BenchmarkRhythm(b *testing.B) {
	cases := [][2]int{
		{3, 8},
		{5, 16},
		{13, 100},
		{100, 1000},
		{1000, 10000},
	}
	for _, c := range cases {
		beats, steps := c[0], c[1]
		b.Run(fmt.Sprintf("E(%d,%d)", beats, steps), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rhythm(beats, steps-beats)
			}
		})
	}
}

// TestShiftIsRotation verifies that shift simply rotates the pattern.
func TestShiftIsRotation(t *testing.T) {
	beats, steps := 3, 8
	base := euclidean(beats, steps, 0)
	for s := 1; s < steps; s++ {
		shifted := euclidean(beats, steps, s)
		for i := range base {
			if shifted[i] != base[(i+s)%steps] {
				t.Errorf("shift=%d: index %d got %v, want %v", s, i, shifted[i], base[(i+s)%steps])
				break
			}
		}
	}
}

package main

func rhythmFast(f, b int) []bool {
	// Degenerate case: no beats and no rests yields an empty pattern.
	if f == 0 && b == 0 {
		return []bool{}
	}
	// All beats: every step is a hit.
	if b == 0 {
		result := make([]bool, f)
		for i := range result {
			result[i] = true
		}
		return result
	}
	// All rests: every step is silent (bool zero value).
	if f == 0 {
		return make([]bool, b)
	}

	// Bjorklund seed: start with f copies of the single-beat pattern [true]
	// and b copies of the single-rest pattern [false].
	frontCount := f
	backCount := b
	frontPat := []bool{true}
	backPat := []bool{false}

	// Bjorklund merge loop: repeatedly fold one back-pattern into each
	// front-pattern to form a new, longer front-pattern, consuming the
	// smaller group's count. When the back group becomes smaller than the
	// front group, swap so the leftover (old front) becomes the new back.
	// Loop ends once either group has only one copy left.
	for frontCount > 1 && backCount > 1 {
		merged := append(append([]bool{}, frontPat...), backPat...)
		if backCount >= frontCount {
			backCount -= frontCount
			frontPat = merged
		} else {
			newBackCount := frontCount - backCount
			frontCount = backCount
			backCount = newBackCount
			backPat = frontPat // old frontPat becomes the new remainder
			frontPat = merged
		}
	}

	// Final assembly: concatenate frontCount copies of frontPat followed by
	// backCount copies of backPat to produce the Euclidean rhythm.
	result := make([]bool, 0, frontCount*len(frontPat)+backCount*len(backPat))
	for i := 0; i < frontCount; i++ {
		result = append(result, frontPat...)
	}
	for i := 0; i < backCount; i++ {
		result = append(result, backPat...)
	}
	return result
}

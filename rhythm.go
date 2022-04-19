package main

import (
	"container/ring"
)

func rhythm(f, b int) *ring.Ring {

	front := ring.New(f)

	for i := 0; i < front.Len(); i++ {
		front.Value = &ring.Ring{Value: true}
		front = front.Next()
	}

	back := ring.New(b)

	for i := 0; i < back.Len(); i++ {
		back.Value = &ring.Ring{Value: false}
		back = back.Next()
	}

	return recurse(front, back)

}

func recurse(front, back *ring.Ring) *ring.Ring {

	// Base Case
	if back.Len() <= 1 || front.Len() <= 1 {

		var builder, merged *ring.Ring

		if front != nil {
			merged = front.Prev().Link(back)
		} else {
			merged = back
		}

		merged.Do(func(value any) {

			linked := value.(*ring.Ring)

			if builder != nil {
				builder = builder.Prev().Link(linked)
			} else {
				builder = linked
			}

		})

		return builder

	}

	// Recursive Step
	var builder, leftover *ring.Ring

	for f, b := front, back; true; f, b = f.Next(), b.Next() {

		linked := &ring.Ring{Value: f.Value.(*ring.Ring).Prev().Link(b.Value.(*ring.Ring))}

		if builder != nil {
			builder = builder.Prev().Link(linked)
		} else {
			builder = linked
		}

		if b == back.Prev() {
			leftover = f.Unlink(f.Len() - b.Len())
			break
		}

		if f == front.Prev() {
			leftover = b.Unlink(b.Len() - f.Len())
			break
		}

	}

	return recurse(builder, leftover)

}

package main

func rhythmFast(f, b int) []bool {
	if f == 0 && b == 0 {
		return []bool{}
	}
	if b == 0 {
		result := make([]bool, f)
		for i := range result {
			result[i] = true
		}
		return result
	}
	if f == 0 {
		return make([]bool, b)
	}

	frontCount := f
	backCount := b
	frontPat := []bool{true}
	backPat := []bool{false}

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

	result := make([]bool, 0, frontCount*len(frontPat)+backCount*len(backPat))
	for i := 0; i < frontCount; i++ {
		result = append(result, frontPat...)
	}
	for i := 0; i < backCount; i++ {
		result = append(result, backPat...)
	}
	return result
}

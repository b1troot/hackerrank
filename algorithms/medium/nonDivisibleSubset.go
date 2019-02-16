package main


func nonDivisibleSubset(k int32, S []int32) int32 {

// this id the map of module values and its count
	modMap := make(map[int32]int32)
	var result int32

	// populate modMap for future comparision
	for _, v := range S {
		m := v % k
		_, exist := modMap[m]
		if exist {
			if m > 0 {
				modMap[m]++
			}
		} else {
			modMap[m] = 1
		}
	}

// iterate over modMap to determine which module have higher count
	for m, c := range modMap {
		// there can be only 1 module with 0 value in the set
		// also check for situations where for example k = 2 m = 1 k-m = 1
		// to make sure you count it only once
		if m == (k-m) || m == 0 {
			result++
			// choose between modules that can't be in the same set based on ther count
		} else if c > modMap[k-m] {
			result += c
			// if both have the same count, choose first and zero the second
		}else if c == modMap[k-m]{
			result +=c
			modMap[k-m] = 0
		}
	}

	return result
}

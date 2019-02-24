package easy

import (
	"math"
)

func viralAdvertising(n int32) int32 {
	var shared int32 = 5
	var liked int32 = 2
	var cumulative int32 = 2
	var i int32 = 2
	for ; i <= n; i++ {

		shared = liked * 3
		liked = int32(math.Floor(float64(shared) / 2))
		cumulative += liked
	}
	return cumulative
}

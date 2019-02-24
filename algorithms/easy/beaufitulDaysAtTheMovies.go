package easy

import "math"

func reverseNumber(n int32) int32 {
	var reversed int32
	var tens int32
	var copy = n
	for copy >= 10 {
		copy /= 10
		tens++
	}
	for n >= 1 {

		reminder := n % 10

		n /= 10
		if reminder > 0 {
			if tens >= 0 {
				reversed += reminder * int32((math.Pow(float64(10), float64(tens))))
			}
		}
		tens--
	}
	return reversed
}

func beautifulDays(i int32, j int32, k int32) int32 {
	var result int32 = 0
	start := i
	stop := j
	for start <= stop {
		if (start-reverseNumber(start))%k == 0 {
			result++
		}
		start++
	}
	return result
}

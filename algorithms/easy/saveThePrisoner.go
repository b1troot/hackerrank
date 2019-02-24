package easy

func saveThePrisoner(n int32, m int32, s int32) int32 {
	poisoned := (s + m - 1) % n
	if poisoned == 0 {
		poisoned = n
	}
	return poisoned
}

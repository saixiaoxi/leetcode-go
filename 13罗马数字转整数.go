package main

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)
	for i := 0; i < n; i++ {
		value := romanMap[s[i]]
		if i < n-1 && value < romanMap[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}

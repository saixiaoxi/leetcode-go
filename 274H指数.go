package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	max := 0
	for _, citation := range citations {
		if (len(citations) - citation) >= citation {
			max = citation
		}
	}
	return max
}

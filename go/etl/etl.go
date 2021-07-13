package etl

import "strings"

func Transform(old map[int][]string) map[string]int {
	new := map[string]int{}

	for s, l := range old {
		for _, lr := range l {
			new[strings.ToLower(lr)] = s
		}
	}

	return new
}

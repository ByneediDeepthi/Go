package reverse

func Rev(s string) string {
	if len(s) == 0 {
		return ""
	}
	revs := []string(s)
	j := len(revs) - 1
	i := 0
	for i < j {
		revs[i], revs[j] = revs[j], revs[i]
		j = j - 1
		i = i + 1
	}
	return revs
}

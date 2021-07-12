package series

func All(n int, s string) []string {
	a := []string{}
	j := 0
	for i := n; i <= len(s); i++ {
		a = append(a, s[j:i])
		j++
	}
	return a
}
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}

package sublist

import "reflect"

func Sublist(A []int, B []int) string {
	if len(A) == len(B) {
		for i := 0; i < len(A); i++ {
			if A[i] != B[i] {
				return "unequal"
			}
		}
		return "equal"
	}
	if len(B) > len(A) {
		for i := 0; i <= len(B)-len(A); i++ {
			res_val := reflect.DeepEqual(B[i:i+len(A)], A[:])
			if res_val == true {
				return "sublist"
			}
		}
	}
	for i := 0; i <= len(A)-len(B); i++ {
		res_val := reflect.DeepEqual(A[i:i+len(B)], B[:])
		if res_val == true {
			return "superlist"
		}
	}
	return "unequal"
}

package perfect

import "errors"

type Classification string

const ClassificationDeficient Classification = "ClassificationDeficient"
const ClassificationPerfect Classification = "ClassificationPerfect"
const ClassificationAbundant Classification = "ClassificationAbundant"

var ErrOnlyPositive = errors.New("input is not a positive integer")

func Classify(n int64) (class Classification, err error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	var sum int64
	for i := 1; int64(i) <= n/2; i++ {
		if int(n)%i == 0 {
			sum += int64(i)
		}
	}
	if sum < n {
		return ClassificationDeficient, nil
	}

	if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationPerfect, nil
}

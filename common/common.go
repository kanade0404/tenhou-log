package common

import (
	"strconv"
	"strings"
)

func SplitByCommaAsInt(str string) ([]int, error) {
	var res []int
	for _, s := range SplitByComma(str) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func SplitByComma(str string) []string {
	return strings.Split(str, ",")
}

func JoinByCommaFromInts(ints []int) string {
	var res []string
	for _, i := range ints {
		res = append(res, strconv.Itoa(i))
	}
	return JoinByComma(res)
}

func JoinByComma(strs []string) string {
	return strings.Join(strs, ",")
}

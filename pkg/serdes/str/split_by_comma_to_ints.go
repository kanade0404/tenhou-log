package str

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

func JoinByComma(strs []string) string {
	return strings.Join(strs, ",")
}

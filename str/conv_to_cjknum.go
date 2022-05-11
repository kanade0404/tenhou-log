package str

var cjkNum = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}

func ConvertToCjkNum(num uint) string {
	if num < 1 || num > 9 {
		return ""
	} else {
		return cjkNum[num]
	}
}

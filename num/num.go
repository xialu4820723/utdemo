package num

func transToStr(num int) string {
	n := num
	result := ""
	for n = num; n > 0; {
		var char string
		switch n % 10 {
		case 0: char = "零"
		case 1: char = "一"
		case 2: char = "二"
		case 3: char = "三"
		case 4: char = "四"
		case 5: char = "五"
		case 6: char = "六"
		case 7: char = "七"
		case 8: char = "八"
		case 9: char = "九"
		}
		n = n / 10
		result = char + result
	}
	return result
}

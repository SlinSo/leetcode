package p0012

// I = 1
// V = 5
// X = 10
// L = 50
// C = 100
// D = 500
// M = 1000
func intToRoman(num int) string {
	s := ""
	m := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	n := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, v := range m {
		for num >= v {
			num -= v
			s += n[i]
		}
	}
	return s
}

func intToRoman2(num int) string {
	s := ""
	m := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	n := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(m); i++ {
		j := num / m[i]
		num %= m[i]
		for ; j > 0; j-- {
			s += n[i]
		}
	}
	return s
}

func intToRomanStupid(num int) string {
	s := ""

	for num/1000 >= 1 {
		num -= 1000
		s += "M"
	}

	if num/900 >= 1 {
		num -= 900
		s += "CM"
	}

	for num/500 >= 1 {
		num -= 500
		s += "D"
	}

	if num/400 >= 1 {
		num -= 400
		s += "CD"
	}

	for num/100 >= 1 {
		num -= 100
		s += "C"
	}

	if num/90 >= 1 {
		num -= 90
		s += "XC"
	}

	for num/50 >= 1 {
		num -= 50
		s += "L"
	}

	if num/40 >= 1 {
		num -= 40
		s += "XL"
	}

	for num/10 >= 1 {
		num -= 10
		s += "X"
	}

	if num/9 >= 1 {
		num -= 9
		s += "IX"
	}

	for num/5 >= 1 {
		num -= 5
		s += "V"
	}

	if num/4 >= 1 {
		num -= 4
		s += "IV"
	}

	for num/1 >= 1 {
		num -= 1
		s += "I"
	}

	return s
}

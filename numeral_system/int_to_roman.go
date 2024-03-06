package numeral_system

import "bytes"

var (
	numMap = map[int]string{
		1:    `I`,
		5:    `V`,
		10:   `X`,
		50:   `L`,
		100:  `C`,
		500:  `D`,
		1000: `M`,
	}

	baseMap = map[int]string{
		0:  ``,
		1:  `I`,
		2:  `II`,
		3:  `III`,
		4:  `IV`,
		5:  `V`,
		6:  `VI`,
		7:  `VII`,
		8:  `VIII`,
		9:  `IX`,
		10: `X`,
	}
)

// Использует рекурсию
func intToRomanRecur(num int) string {
	if v, ok := baseMap[num]; ok {
		return v
	}

	var b bytes.Buffer

	if num < 40 {
		for num >= 10 {
			b.WriteString(numMap[10])
			num -= 10
		}
		b.WriteString(intToRomanRecur(num))
	} else if num >= 40 && num < 50 {
		b.WriteString(numMap[10])
		b.WriteString(numMap[50])
		num -= 40

		b.WriteString(intToRomanRecur(num))
	} else if num >= 50 && num < 90 {
		b.WriteString(numMap[50])
		num -= 50
		b.WriteString(intToRomanRecur(num))
	} else if num >= 90 && num < 100 {
		b.WriteString(numMap[10])
		b.WriteString(numMap[100])
		num -= 90
		b.WriteString(intToRomanRecur(num))
	} else if num >= 100 && num < 400 {
		for num >= 100 {
			b.WriteString(numMap[100])
			num -= 100
		}
		b.WriteString(intToRomanRecur(num))
	} else if num >= 400 && num < 500 {
		b.WriteString(numMap[100])
		b.WriteString(numMap[500])
		num -= 400
		b.WriteString(intToRomanRecur(num))
	} else if num >= 500 && num < 900 {
		b.WriteString(numMap[500])
		num -= 500
		b.WriteString(intToRomanRecur(num))
	} else if num >= 900 && num < 1000 {
		b.WriteString(numMap[100])
		b.WriteString(numMap[1000])

		num -= 900
		b.WriteString(intToRomanRecur(num))
	} else if num >= 1000 {
		b.WriteString(numMap[1000])
		num -= 1000
		b.WriteString(intToRomanRecur(num))
	}

	return b.String()
}

// использует итерации
func intToRomanIter(num int) string {
	var b bytes.Buffer

	for num > 0 {
		if v, ok := baseMap[num]; ok {
			b.WriteString(v)

			return b.String()
		}

		if num < 40 {
			for num >= 10 {
				b.WriteString(numMap[10])
				num -= 10
			}
		} else if num >= 40 && num < 50 {
			b.WriteString(numMap[10])
			b.WriteString(numMap[50])
			num -= 40
		} else if num >= 50 && num < 90 {
			b.WriteString(numMap[50])
			num -= 50
		} else if num >= 90 && num < 100 {
			b.WriteString(numMap[10])
			b.WriteString(numMap[100])
			num -= 90
		} else if num >= 100 && num < 400 {
			for num >= 100 {
				b.WriteString(numMap[100])
				num -= 100
			}
		} else if num >= 400 && num < 500 {
			b.WriteString(numMap[100])
			b.WriteString(numMap[500])
			num -= 400
		} else if num >= 500 && num < 900 {
			b.WriteString(numMap[500])
			num -= 500
		} else if num >= 900 && num < 1000 {
			b.WriteString(numMap[100])
			b.WriteString(numMap[1000])

			num -= 900
		} else if num >= 1000 {
			b.WriteString(numMap[1000])
			num -= 1000
		}

	}
	return b.String()
}

func intToRomanModule(num int) string {
	var b bytes.Buffer

	units := map[int]string{
		0: ``,
		1: `I`,
		2: `II`,
		3: `III`,
		4: `IV`,
		5: `V`,
		6: `VI`,
		7: `VII`,
		8: `VIII`,
		9: `IX`,
	}
	dec := map[int]string{
		1: `X`,
		2: `XX`,
		3: `XXX`,
		4: `XL`,
		5: `L`,
		6: `LX`,
		7: `LXX`,
		8: `LXXX`,
		9: `XC`,
	}
	cent := map[int]string{
		1: `C`,
		2: `CC`,
		3: `CCC`,
		4: `CD`,
		5: `D`,
		6: `DC`,
		7: `DCC`,
		8: `DCCC`,
		9: `CM`,
	}
	mill := map[int]string{
		1: `M`,
		2: `MM`,
		3: `MMM`,
	}

	mill_num := num / 1000
	if v, ok := mill[mill_num]; ok {
		b.WriteString(v)

		num -= mill_num * 1000
	}

	cent_num := num / 100
	if v, ok := cent[cent_num]; ok {
		b.WriteString(v)

		num -= cent_num * 100
	}

	dec_num := num / 10
	dec_rem := num % 10
	if v, ok := dec[dec_num]; ok {
		b.WriteString(v)
	}

	b.WriteString(units[dec_rem])

	return b.String()
}

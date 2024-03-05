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

func intToRoman(num int) string {
	if v, ok := baseMap[num]; ok {
		return v
	}

	var b bytes.Buffer

	if num < 40 {
		for num >= 10 {
			b.WriteString(numMap[10])
			num -= 10
		}
		b.WriteString(intToRoman(num))
	} else if num >= 40 && num < 50 {
		b.WriteString(numMap[10])
		b.WriteString(numMap[50])
		num -= 40

		b.WriteString(intToRoman(num))
	} else if num >= 50 && num < 90 {
		b.WriteString(numMap[50])
		num -= 50
		b.WriteString(intToRoman(num))
	} else if num >= 90 && num < 100 {
		b.WriteString(numMap[10])
		b.WriteString(numMap[100])
		num -= 90
		b.WriteString(intToRoman(num))
	} else if num >= 100 && num < 400 {
		for num >= 100 {
			b.WriteString(numMap[100])
			num -= 100
		}
		b.WriteString(intToRoman(num))
	} else if num >= 400 && num < 500 {
		b.WriteString(numMap[100])
		b.WriteString(numMap[500])
		num -= 400
		b.WriteString(intToRoman(num))
	} else if num >= 500 && num < 900 {
		b.WriteString(numMap[500])
		num -= 500
		b.WriteString(intToRoman(num))
	} else if num >= 900 && num < 1000 {
		b.WriteString(numMap[100])
		b.WriteString(numMap[1000])

		num -= 900
		b.WriteString(intToRoman(num))
	} else if num >= 1000 {
		b.WriteString(numMap[1000])
		num -= 1000
		b.WriteString(intToRoman(num))
	}

	return b.String()
}

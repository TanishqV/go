package stralign

import "strings"

func pad(str string, left int, right int, fill string) string {
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}

	if left == 0 && right == 0 {
		return str
	}

	var strBuff strings.Builder

	if left >= 0 {
		strBuff.WriteString(strings.Repeat(fill, left))
	}

	strBuff.WriteString(str)

	if right >= 0 {
		strBuff.WriteString(strings.Repeat(fill, right))
	}

	return strBuff.String()
}

// Works like fmt.Printf("%-{width}s", {str}) when fill is " "
func Ljust(str string, width int, fill string) string {
	if len(str) >= width {
		return str
	}

	return pad(str, 0, width - len(str), fill)
}

// Works like fmt.Printf("%{width}s", {str}) when fill is " "
func Rjust(str string, width int, fill string) string {
	if len(str) >= width {
		return str
	}

	return pad(str, width - len(str), 0, fill)
}

func Center(str string, width int, fill string) string {
	if len(str) >= width {
		return str
	}

	marg := width - len(str)
	left := marg/2 + (marg & width & 1)

	return pad(str, left, marg - left, fill)
}

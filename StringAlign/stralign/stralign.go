package stralign

import "strings"

func pad(str string, left int32, right int32, fill string) (string, error) {
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}

	if left == 0 && right == 0 {
		return str, nil
	}

	var strBuff strings.Builder
	var err error

	if left >= 0 {
		if _, err = strBuff.WriteString(strings.Repeat(fill, int(left))); err != nil {
			return "", err
		}
	}

	if _, err = strBuff.WriteString(str); err != nil {
		return "", err
	}

	if right >= 0 {
		if _, err = strBuff.WriteString(strings.Repeat(fill, int(right))); err != nil {
			return "", err
		}
	}

	return strBuff.String(), err
}

// Works like fmt.Printf("%-{width}s", {str}) when fill is " "
func Ljust(str string, width int32, fill string) (string, error) {
	lenStr := int32(len(str))

	if lenStr >= width {
		return str, nil
	}

	return pad(str, 0, width-lenStr, fill)
}

// Works like fmt.Printf("%{width}s", {str}) when fill is " "
func Rjust(str string, width int32, fill string) (string, error) {
	lenStr := int32(len(str))

	if lenStr >= width {
		return str, nil
	}

	return pad(str, width-lenStr, 0, fill)
}

func Center(str string, width int32, fill string) (string, error) {
	lenStr := int32(len(str))

	if lenStr >= width {
		return str, nil
	}

	marg := width - lenStr
	left := marg/2 + (marg & width & 1)

	return pad(str, left, marg-left, fill)
}

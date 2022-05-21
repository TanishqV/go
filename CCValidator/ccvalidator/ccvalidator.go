package ccvalidator

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/magiconair/properties"
)

var manufMap map[string]string
var CCErrors map[string]error

func init() {
	// load manufacturers' Card regex from ini file
	manufMap = properties.MustLoadFile("../ccdistributors.ini", properties.UTF8).Map()

	CCErrors = map[string]error{
		"NoNumber":      errors.New("No card number provided"),
		"NoRegManuf":    errors.New("Card does not have a registered manufacturer"),
		"InvalidNumber": errors.New("Card number is not valid"),
		"InvalidLength": errors.New("Digits in card number should be between 13 and 16"),
	}
}

func Validate(cardNumber string) (manufacturer string, ok bool, err error) {
	cardNumber = removeSpaces(cardNumber)

	if cardNumber == "" {
		return "", false, CCErrors["NoNumber"]
	}

	if len(cardNumber) >= 13 && len(cardNumber) <= 16 {
		even2xSum := 0
		oddSum := 0
		isSecond := false
		for i := len(cardNumber) - 1; i >= 0; i-- {
			tmp := int(cardNumber[i] - '0')

			if isSecond == true {
				tmp := 2 * tmp
				if tmp > 9 {
					tmp -= 9
				}
				even2xSum += tmp
			} else {
				oddSum += tmp
			}
			isSecond = !isSecond
		}

		if (oddSum+even2xSum)%10 == 0 {
			var matchFound bool
			var manuf, pattern string

			for manuf, pattern = range manufMap {
				matchFound, _ = regexp.MatchString(pattern, cardNumber)
				if matchFound == true {
					manufacturer = manuf
					ok = true
					err = nil
					break
				}
			}
			if matchFound == false {
				manufacturer = ""
				err = CCErrors["NoRegManuf"]
				ok = false
			}
		} else {
			manufacturer = ""
			ok = false
			err = CCErrors["InvalidNumber"]
		}
	} else {
		manufacturer = ""
		ok = false
		err = CCErrors["InvalidLength"]
	}
	return manufacturer, ok, err
}

func removeSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

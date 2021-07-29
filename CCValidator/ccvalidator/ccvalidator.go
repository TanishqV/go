package ccvalidator

import (
	"regexp"
	"errors"
)

func Validate(cardNumber string) (manufacturer string, ok bool, err error) {
	if cardNumber == "" {
		return "", false, errors.New("No card number provided")
	}

	m := map[string]string {
		"Visa"		:	`(4[0-9]{12}(?:[0-9]{3}?))`,
		"MasterCard"	:	`(5[1-5][0-9]{14})`,
		"AMEX"		:	`(3[47][0-9]{13})`,
		"Discover"	:	`(6(?:011|5[0-9]{2})[0-9]{12})`,
		"Diners Club"	:	`(3:(?:0[0-5]|[68][0-9])[0-9]{11})`,
		"JCB"		:	`((?:2131|1800|35[0-9]{3})[0-9]{11})`,
	}
	even2xSum := 0
	oddSum := 0
	for i:=0; i<len(cardNumber); i++ {
		if i%2 == 0 {
			tmp := 2*int(cardNumber[i] - '0')
			if tmp > 9 {
				tmp = tmp%10 + tmp/10
				// OR
				// tmp -= 9
			}
			even2xSum += tmp
		} else {
			oddSum += int(cardNumber[i] - '0')
		}
	}
	if (len(cardNumber) >= 13 && len(cardNumber) <= 16) &&
	(oddSum + even2xSum) % 10 == 0 {
		ok = true
		var matchFound bool
		var manuf, pattern string

		for manuf,pattern = range m {
			matchFound, _ = regexp.MatchString(pattern, cardNumber)
			if matchFound == true {
				break
			}
		}
		if matchFound == false {
			err = errors.New("New manufacturer, but satisfies algo")
		} else {
			manufacturer = manuf
			err = nil
		}
		ok = matchFound

	} else {
		ok = false
		err = errors.New("New manufacturer")
	}
	return manufacturer, ok, err
}

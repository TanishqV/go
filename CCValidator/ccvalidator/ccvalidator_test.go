package ccvalidator

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Random String generation using: https://stackoverflow.com/a/31832326
const letterBytes = "12345678909876543210"
const (
	letterIdxBits = 6                    // 6 bits to represent letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func TestValidate(t *testing.T) {
	testCases := []struct {
		name         string
		cardNumber   string
		manufacturer string
		ok           bool
		err          error
	}{
		{
			name:         "Empty Card",
			cardNumber:   "",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["NoNumber"],
		},
		{
			name:         "Short numbered Card",
			cardNumber:   "12",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["InvalidLength"],
		},
		{
			name:         "Long numbered Card",
			cardNumber:   "4012 8888 8888 1881 9901",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["InvalidLength"],
		},
		{
			name:         "Invalid Card-13",
			cardNumber:   "8451 432 325 829",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["InvalidNumber"],
		},
		{
			name:         "Invalid Card-14",
			cardNumber:   "9451 2476 436 472",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["InvalidNumber"],
		},
		{
			name:         "Invalid Card-15",
			cardNumber:   "8515 6610 1291 183",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["InvalidNumber"],
		},
		{
			name:         "Invalid Card-16",
			cardNumber:   "8451 2476 4364 7246",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["InvalidNumber"],
		},
		{
			name:         "Unregistered Manufacturer-13",
			cardNumber:   "1800 941 478 424",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["NoRegManuf"],
		},
		{
			name:         "Unregistered Manufacturer-14",
			cardNumber:   "3253 6745 627 666",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["NoRegManuf"],
		}, {
			name:         "Unregistered Manufacturer-15",
			cardNumber:   "3352 6547 6672 667",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["NoRegManuf"],
		}, {
			name:         "Unregistered Manufacturer-16",
			cardNumber:   "3352 6547 6672 6626",
			manufacturer: "",
			ok:           false,
			err:          CCErrors["NoRegManuf"],
		},
		{
			name:         "VISA-1",
			cardNumber:   "4012 8888 8888 1881",
			manufacturer: "VISA",
			ok:           true,
			err:          nil,
		},
		{
			name:         "VISA-2",
			cardNumber:   "4222 222 222 222",
			manufacturer: "VISA",
			ok:           true,
			err:          nil,
		},
		{
			name:         "Diner's Club",
			cardNumber:   "3852 0000 0232 37",
			manufacturer: "DINERS_CLUB",
			ok:           true,
			err:          nil,
		},
		{
			name:         "Amex",
			cardNumber:   "3714 496353 98431",
			manufacturer: "AMEX",
			ok:           true,
			err:          nil,
		},
		{
			name:         "MasterCard",
			cardNumber:   "5506 9208 0924 3667",
			manufacturer: "MASTERCARD",
			ok:           true,
			err:          nil,
		},
		{
			name:         "JCB",
			cardNumber:   "3566 0020 2036 0505",
			manufacturer: "JCB",
			ok:           true,
			err:          nil,
		},
		{
			name:         "Discover",
			cardNumber:   "6011 0009 9013 9424",
			manufacturer: "DISCOVER",
			ok:           true,
			err:          nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			// res, _ := generateNums("8451", 15)
			// fmt.Println(res)
			manufacturer, ok, err := Validate(tc.cardNumber)

			require.Equal(t, tc.ok, ok)
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.manufacturer, manufacturer)
		})
	}
}

func generateNums(part string, length int) (res []string, err error) {
	if length == 0 {
		return []string{""}, nil
	}
	num, err := randStringGenerator(length - len(part))
	if err != nil {
		return []string{}, err
	}

	num = part + num
	res = append(res, num)
	return res, nil
}

func randStringGenerator(n int) (string, error) {
	sb := strings.Builder{}
	var err error
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			err = sb.WriteByte(letterBytes[idx])
			if err != nil {
				return "", err
			}
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String(), nil
}

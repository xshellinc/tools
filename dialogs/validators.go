package dialogs

import (
	"fmt"
	"regexp"
	"strings"
)

// Validator type will be called on user input
// return true if valid input, false otherwise
type ValidatorFn func(input string) bool

// return true if string is not empty and valid, false otherwise
func EmptyStringValidator(inp string) bool {
	if inp == "" {
		fmt.Print("[-] Empty input, please repeat:")
		return false
	}
	return true
}

// return true if parsed IP is not nil, false otherwise
func IpAddressValidator(inp string) bool {
	fmt.Println("[+] Validating IP address:", inp)

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(inp) {
		return true
	}
	fmt.Print("[-] Not valid IP address, please repeat:")
	return false
}

func YesNoValidator(inp string) bool {
	if strings.EqualFold(inp, "y") || strings.EqualFold(inp, "yes") ||
		strings.EqualFold(inp, "n") || strings.EqualFold(inp, "no") {
		return true
	} else {
		fmt.Print("[-] Unknown user input. Please enter (\x1b[33my/yes\x1b[0m OR \x1b[33mn/no\x1b[0m):")
		return false
	}
}

func CreateValidatorFn(fn func(string) error) ValidatorFn {
	return func(inp string) bool {
		err := fn(inp)
		if err != nil {
			fmt.Println("[-] ", err)
			return false
		}

		return true
	}
}

func SpecialCharacterValidator(str string, cond bool) ValidatorFn {
	return func(inp string) bool {
		r, err := regexp.Compile(`[` + str + `]`)

		if err != nil {
			fmt.Println("[-] ", err)
			return false
		}

		c := r.Match([]byte(inp))
		if (c || cond) && !(c && cond) {
			return true
		}

		return false
	}
}

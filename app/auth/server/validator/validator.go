package validator

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	Email       string `validate:"required,email"`
	Name        string `validate:"required,alphaunicode,max=32"`
	Password    string `validate:"required,password"`
	PhoneNumber string `validate:"required,phonenumber"`
}

var (
	passwordMinLen    = 8
	passwordMaxLen    = 72
	prefixPhoneNumber = [...]string{"010", "011", "016", "018", "019"}
)

func IsValidAccount(account *Account) error {
	validate := validator.New()
	validate.RegisterValidation("phonenumber", isValidPhoneNumber)
	validate.RegisterValidation("password", isValidPassword)

	return validate.Struct(account)
}

func isValidPassword(fl validator.FieldLevel) bool {
	pw := fl.Field().String()

	// repeating same characters <3
	cnt := 0
	var old_v rune
	for _, v := range pw {
		if old_v != v {
			cnt = 0
			old_v = v
			continue
		}
		cnt++
		if cnt >= 2 {
			return false
		}
	}
	// The regexp package uses re2syntax. It doesn't support backreference to keep linear time.
	// match, _ := regexp.MatchString(`(.)\1\1+`, pw)
	// if match {
	// 	fmt.Println("same character")
	// 	return false
	// }
	if len(pw) < passwordMinLen || len(pw) > passwordMaxLen {
		return false
	}

	// alphabet|number|special combination >= 2
	// This includes all characters àb, 😀 ...
	// consider changing to below regex.
	// ~`!@#$%^&*()-_+={}[]|\/:;"'<>,.?₩
	cnt = 0
	if match, _ := regexp.MatchString(`[^A-Za-z0-9]`, pw); match {
		cnt++
	}
	if match, _ := regexp.MatchString(`[A-Za-z]`, pw); match {
		cnt++
	}
	if match, _ := regexp.MatchString(`[0-9]`, pw); match {
		cnt++
	}
	if cnt < 2 {
		return false
	}

	return true
}

// phonenumber's length: 11, remove '-'(hyphen), prefix: 010|011|016|017|018|019
func isValidPhoneNumber(fl validator.FieldLevel) bool {
	number := strings.Replace(fl.Field().String(), "-", "", -1)

	if len(number) != 11 {
		return false
	}

	_, err := strconv.Atoi(number)
	if err != nil {
		return false
	}

	prefix := number[:3]
	for _, val := range prefixPhoneNumber {
		if val == prefix {
			return true
		}
	}
	return false
}

package utility

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"infopack.co.in/offybox/app/common/constants"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func MapToJSON(inputMap map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// JSONToStruct converts JSON string to a struct
func JSONToStruct(jsonData string, resultStruct interface{}) error {
	err := json.Unmarshal([]byte(jsonData), resultStruct)
	return err
}

func ValidatePlatform(platform string) string {
	switch platform {
	case constants.SystemEmployeeAPI:
		return constants.UserTypeEmployee
	case constants.SystemDistributorAPI:
		return constants.UserTypeDistributor
	default:
		return platform
	}
}

func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// HashPassword returns a hashed password
func HashPassword(password string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func StringToInt(value string) int {
	id, _ := strconv.Atoi(value)
	/*if err != nil {
		fmt.Printf("string conversion failedL %v\n", err)
	}*/
	return id
}

func IsValidColumn(columnsStruct interface{}, column string) bool {
	v := reflect.ValueOf(columnsStruct)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == column {
			return true
		}
	}
	return false
}

func ExtractFirstAndLastName(fullName string) (string, string) {
	// Split the full name into words
	words := strings.Fields(fullName)

	// If there are no words or only one word, consider it as the first name
	if len(words) == 0 {
		return "", ""
	} else if len(words) == 1 {
		return words[0], ""
	}

	// If there are multiple words, the last word is the last name
	// and the rest are part of the first name
	lastName := words[len(words)-1]
	firstName := strings.Join(words[:len(words)-1], " ")

	return firstName, lastName
}

/*

CODE_STATE_AUTO:6_START:5
CODE_YYYY_AUTO:6_START:0
CODE_YYMM_AUTO:6_START:0
CODE_PX_LOAN_AUTO:6_START:0

*/

func GenerateCode(pseudo string, prefix string, current *string, masterId *int) (generatedCode string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic %v", r)
		}
	}()
	fmt.Println("pseudo :: ", pseudo)
	code := strings.Split(pseudo, "_")
	currentTime := time.Now()

	var result string
	var auto string
	var start string

	for _, s := range code {

		if strings.HasPrefix(s, "AUTO:") {
			_auto := strings.Split(s, ":")
			auto = _auto[1]
		} else if strings.HasPrefix(s, "START:") {
			_start := strings.Split(s, ":")
			start = _start[1]
		} else {
			switch s {
			case "CODE":
				result += prefix
				break
			case "YYYY":
				result += currentTime.Format("2006")
				break
			case "YYMM":
				result += currentTime.Format("0601")
				break
			case "YYDD":
				result += currentTime.Format("0602")
			default:
				result += s
				break
			}
		}
	}

	if start == "" || auto == "" {
		return "", errors.New(fmt.Sprintf("invalid pseudo %s", pseudo))
	}

	var series string
	if *current == "" || !strings.Contains(*current, "-") {
		series = fmt.Sprintf("%s%s%s", start, strings.Repeat("0", StringToInt(auto)-(1+len(start))), "1")
	} else {
		fmt.Println("*current :: ", *current)
		last := strings.Split(*current, "-")
		fmt.Println("last  :: ", last)
		series = GetNewRunningSeries(last[1])
	}

	return fmt.Sprintf("%s-%s", result, series), nil
}

func GetNewRunningSeries(str string) (nextStr string) {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	num++
	strLen := len(str)
	nextStr = fmt.Sprintf("%0*d", strLen, num)
	return
}

func IsExpired(expirationTime time.Time) bool {
	// Compare the current time with the expiration time
	return time.Now().After(expirationTime)
}

func IsValidPassword(password string) bool {
	// Check minimum length
	if len(password) < 9 {
		return false
	}

	// Check for at least one special character
	specialCharPattern := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !specialCharPattern.MatchString(password) {
		return false
	}

	// Check for at least one letter, one number, and one uppercase letter
	var hasLetter, hasDigit, hasUppercase bool
	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
			if unicode.IsUpper(char) {
				hasUppercase = true
			}
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasLetter || !hasDigit || !hasUppercase {
		return false
	}

	// Check for spaces
	if strings.Contains(password, " ") {
		return false
	}

	return true
}

func ParseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}

	return t
}

func StringToInt64(value string) int64 {
	finalValue, _ := strconv.ParseInt(value, 10, 64)
	return finalValue
}
func StringToFloat64(value string) float64 {
	finalValue, _ := strconv.ParseFloat(value, 64)
	return finalValue
}

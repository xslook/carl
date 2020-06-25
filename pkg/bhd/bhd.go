package bhd

import (
	"fmt"
	"strconv"
	"strings"
)

func hex2Decimal(input int64) (int64, error) {
	return input, nil
}

// Result for bhd convert
type Result struct {
	Hex string `json:"hexadecimal"`
	Dec string `json:"decimal"`
	Oct string `json:"octal"`
	Bin string `json:"binary"`

	Value int64 `json:"value"`
}

func isBinary(input string) bool {
	if input == "" {
		return false
	}
	if !strings.HasPrefix(input, "0b") {
		return false
	}
	input = string([]byte(input)[2:])
	if len(input) > 64 {
		return false
	}
	for _, char := range input {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

func isOctal(input string) bool {
	if input == "" {
		return false
	}
	if !strings.HasPrefix(input, "0o") && !strings.HasPrefix(input, "0") {
		return false
	}
	if strings.HasPrefix(input, "0o") {
		input = string([]byte(input)[2:])
	}
	if len(input) > 22 {
		return false
	}
	for _, char := range input {
		if char < '0' || char > '7' {
			return false
		}
	}
	return true
}

func isDecimal(input string) bool {
	if input == "" {
		return false
	}
	if len(input) > maxDecimalLength {
		return false
	}
	// 0XXX consider as octal number
	if strings.HasPrefix(input, "0") {
		return false
	}
	isNegative := strings.HasPrefix(input, "-")
	if isNegative {
		input = strings.TrimLeft(input, "-")
	}
	for _, char := range input {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

const (
	maxDecimalLength = 20

	minHexNumberLength = 3  // minimum number: 0x0
	maxHexNumberLength = 18 // maximum number: 0xFFFFFFFFFFFFFFFF
)

func isHexadecimal(input string) bool {
	if !strings.HasPrefix(input, "0x") && !strings.HasPrefix(input, "0X") {
		return false
	}
	if len(input) < minHexNumberLength {
		return false
	}
	if len(input) > maxHexNumberLength {
		return false
	}
	for i := 2; i < len(input); i++ {
		char := input[i]
		if char >= '0' && char <= '9' {
			continue
		}
		if char >= 'A' && char <= 'F' {
			continue
		}
		if char >= 'a' && char <= 'f' {
			continue
		}
		return false
	}
	return true
}

func convert(input string) (*Result, error) {
	if input == "" {
		return nil, fmt.Errorf("Empty input")
	}

	if !isDecimal(input) && !isBinary(input) && !isOctal(input) && !isHexadecimal(input) {
		return nil, fmt.Errorf("Input is not a valid number: %s", input)
	}

	res, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		return nil, err
	}

	result := &Result{
		Bin: fmt.Sprintf("%b", res),
		Oct: fmt.Sprintf("%o", res),
		Dec: fmt.Sprintf("%d", res),
		Hex: fmt.Sprintf("%x", res),

		Value: res,
	}
	return result, nil
}

// Convert numbers
func Convert(input string) (*Result, error) {
	return convert(input)
}

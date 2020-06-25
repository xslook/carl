package bhd

import (
	"testing"
)

func TestIsHexadecimal(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"10000a", false},
		{"0X", false},
		{"0XFF", true},
		{"0xABCDEF", true},
		{"0xABCDEFG", false},
		{"0xFFFFFFFFFFFFFFFF", true},
		{"0xFFFFFFFFFFFFFFFF0", false},
	}

	for _, c := range cases {
		ret := isHexadecimal(c.input)
		if ret != c.expect {
			t.Errorf("Verify %s failed, expect: %t, got: %t", c.input, c.expect, ret)
		}
	}
}

func TestIsBinary(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"0b1010", true},
		{"0b1101", true},
		{"010000", false},
		{"0o1000", false},
		{"0o2342", false},
		{"0o1876", false},
		{"123456", false},
		{"10000a", false},
	}

	for _, c := range cases {
		val := isBinary(c.input)
		if val != c.expect {
			t.Errorf("Verify binanry %s failed, expect: %t, got: %t", c.input, c.expect, val)
		}
	}
}

func TestIsOctal(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"0b1010", false},
		{"0b1101", false},
		{"010000", true},
		{"0o1000", true},
		{"0o2342", true},
		{"0o1876", false},
		{"123456", false},
		{"10000a", false},
	}

	for _, c := range cases {
		val := isOctal(c.input)
		if val != c.expect {
			t.Errorf("Verify binanry %s failed, expect: %t, got: %t", c.input, c.expect, val)
		}
	}
}

func TestIsDeciaml(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"0b1010", false},
		{"0b1101", false},
		{"010000", false},
		{"0o1000", false},
		{"0o2321", false},
		{"0o1876", false},
		{"123456", true},
		{"10000a", false},
	}

	for _, c := range cases {
		val := isDecimal(c.input)
		if val != c.expect {
			t.Errorf("Verify binanry %s failed, expect: %t, got: %t", c.input, c.expect, val)
		}
	}
}
func TestConvert(t *testing.T) {
	cases := []struct {
		input  string
		expect int64
	}{
		{"0b100", 4},
		{"0b10111010110", 1494},
		{"0100", 64},
		{"0o23456", 10030},
		{"0o100101", 32833},
		{"0xFF", 255},
		{"0x666FFF", 6713343},
	}

	for _, c := range cases {
		v, err := convert(c.input)
		if err != nil {
			t.Errorf("Convert %s error: %v", c.input, err)
			return
		}
		if v.Value != c.expect {
			t.Errorf("Convert %s failed, expect: %d, got: %d", c.input, c.expect, v.Value)
		}
	}
}

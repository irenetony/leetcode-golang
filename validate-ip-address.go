package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	IPV4    = "IPv4"
	IPV6    = "IPv6"
	Neither = "Neither"
)

func isValidIPV4(IP string) bool {
	array := strings.Split(IP, ".")
	if array == nil || len(array) != 4 {
		return false
	}
	for _, val := range array {
		if !isValidSingleIPV4(val) {
			return false
		}
	}
	return true
}

func isValidSingleIPV4(IP string) bool {
	if len(IP) == 0 || IP[0] == '-' {
		return false
	}

	// check that IP whether has leading zero or not
	if len(IP) >= 2 && IP[0] == '0' {
		return false
	}
	res, err := strconv.Atoi(IP)
	if err != nil {
		return false
	}
	if res >= 0 && res <= 255 {
		return true
	} else {
		return false
	}
}

func isValidIPV6(IP string) bool {
	array := strings.Split(IP, ":")
	if array == nil || len(array) != 8 {
		return false
	}
	for _, val := range array {
		if !isValidSingleIPV6(val) {
			return false
		}
	}

	return true
}

func isValidSingleIPV6(IP string) bool {
	if len(IP) == 0 || IP[0] == '-' {
		return false
	}

	if len(IP) > 4 {
		return false
	}
	for _, val := range IP {
		if !((val >= '0' && val <= '9') || (val >= 'a' && val <= 'f') || (val >= 'A' && val <= 'F')) {
			return false
		}
	}
	return true
}

func ValidIPAddress(IP string) string {
	res := Neither
	if isValidIPV4(IP) {
		res = IPV4
	} else if isValidIPV6(IP) {
		res = IPV6
	}
	return res
}

func TestValidIPAddress(IP string) {
	output := ValidIPAddress(IP)
	fmt.Printf("Input: \"%s\"\n", IP)
	fmt.Printf("Output: \"%s\"\n", output)
	var explanation string
	if IPV4 == output {
		explanation = "This is a valid IPv4 address, return \"IPv4\"."
	} else if IPV6 == output {
		explanation = "This is a valid IPv6 address, return \"IPv6\"."
	} else if Neither == output {
		explanation = "This is neither a IPv4 address nor a IPv6 address."
	} else {
		panic(explanation)
	}
	fmt.Printf("Explanation: %s\n", explanation)
}

func main() {
	TestValidIPAddress("1.0.1.")
	fmt.Println()
	TestValidIPAddress("172.16.254.1")
	fmt.Println()
	TestValidIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334")
	fmt.Println()
	TestValidIPAddress("256.256.256.256")
	fmt.Println()
	TestValidIPAddress("2001:db8:85a3:0:0:8A2E:0370:7334")
	fmt.Println()
	TestValidIPAddress("2001:0db8:85a3::8A2E:0370:7334")
	fmt.Println()
	TestValidIPAddress("02001:0db8:85a3:0000:0000:8a2e:0370:7334")
}

/**
 * @Author: liuzhiqi
 * @Data: 2021/11/9 19:24
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if validIPv6Address(IP) {
		return "IPv6"
	}
	if validIPv4Address(IP) {
		return "IPv4"
	}
	return "Neither"
}

func validIPv4Address(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil || num > 255 || num < 0 {
			return false
		}
		newStr := fmt.Sprint(num)
		if str != newStr {
			return false
		}
	}
	return true
}

func validIPv6Address(IP string) bool {
	strArr := strings.Split(IP, ":")
	if len(strArr) != 8 {
		return false
	}
	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}
		for i := 0; i < len(str); i++ {
			if !(str[i] >= '0' && str[i] <= '9') &&
				!(str[i] >= 'a' && str[i] <= 'f') &&
				!(str[i] >= 'A' && str[i] <= 'F') {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "0123"
	a, _ := strconv.Atoi(s)
	b := fmt.Sprintf(s)
	fmt.Println(a)
	fmt.Println(b)
}

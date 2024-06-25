package core

import "strings"

func GenerateString(s1, s2 string) string{
	newStr := strings.Fields(s2)
	return s1[:2] + newStr[1]
}
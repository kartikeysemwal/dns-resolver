package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 00160100000100000000000003646e7306676f6f676c6503636f6d0000010001
// 00160100000100000000000003646e7306676f6f676c6503636f6d0000010001

// 0016
// 0100
// 0001
// 0000
// 0000
// 0000
// 03646e7306676f6f676c6503636f6d: dns.google.com
// 00: reserved bits
// 0001
// 0001

// ID: 22
// Flags: 256
// Number of questions: 0
// Number of answer resource records: 0
// Number of authority resource records: 0
// Number of additional resource records: 0
// Question: "dns.google.com"
// reserved bit 0
// Query type: 1 (A record)
// Query class: 1 (Internet class)

func getByteHexString(input int, byteLen int64) string {
	var resultString string = strconv.FormatInt(int64(input), 16)

	len := len(resultString)

	for i := 0; i < (int(byteLen*2) - len); i++ {
		resultString = "0" + resultString
	}

	return resultString
}

func convertURLToHexString(input string) string {
	var resultString string

	splittedString := strings.Split(input, ".")

	for _, str := range splittedString {
		var len = len(str)

		resultString = resultString + getByteHexString(len, 1)

		for _, x := range str {
			resultString = resultString + getByteHexString(int(x), 1)
		}
	}

	return resultString
}

func main() {
	var id int = 22
	var flags int = 256
	var nQuestions int = 1
	var nAnswers int = 0
	var nAuthority int = 0
	var nAdditionalResource int = 0
	var urlString = "dns.google.com"
	var nReservedBit int = 0
	var nQueryType int = 1
	var nQueryClass int = 1

	// one hex digit is 4 bit, so for 2 byte values 4 * 4 bit should be there

	var resultStr string
	resultStr = resultStr + getByteHexString(id, 2)
	resultStr = resultStr + getByteHexString(flags, 2)
	resultStr = resultStr + getByteHexString(nQuestions, 2)
	resultStr = resultStr + getByteHexString(nAnswers, 2)
	resultStr = resultStr + getByteHexString(nAuthority, 2)
	resultStr = resultStr + getByteHexString(nAdditionalResource, 2)
	resultStr = resultStr + convertURLToHexString(urlString)
	resultStr = resultStr + getByteHexString(nReservedBit, 1)
	resultStr = resultStr + getByteHexString(nQueryType, 2)
	resultStr = resultStr + getByteHexString(nQueryClass, 2)

	fmt.Println(resultStr)
}

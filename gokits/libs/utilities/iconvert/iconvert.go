package iconvert

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(s)
	return int32(i), err
}

func StringToUint32(s string) (uint32, error) {
	i, err := strconv.Atoi(s)
	return uint32(i), err
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func Int8ToBool(b int8) bool {
	if b == 1 {
		return true
	} else {
		return false
	}
}

func HexToInt64(hexString string) int64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexString, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return int64(result)
}

func Int64ToBin(value int64) string {
	return strconv.FormatInt(value, 2) // base 2 for binary
}

func BinToInt64(binString string) int64 {
	result, _ := strconv.ParseInt(binString, 2, 64)
	return result
}

func BinTo2sCompleteInt32(bin string) int32 {
	if strings.HasPrefix(bin, "0") {
		return int32(BinToInt64(bin))
	}
	myBin := ""
	for _, r := range bin {
		myBin += string(r)
	}
	return int32(BinToInt64(myBin))
}

func HexTo2sCompleteInt32(hexString string) int32 {
	return BinTo2sCompleteInt32(Int64ToBin(HexToInt64(hexString)))
}

func StringToIntArray(input string) []int {
	// Split the input string by commas
	values := strings.Split(input, ",")

	// Initialize an empty integer array
	var result []int

	// Convert and append each substring to the result array
	for _, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Error converting '%s' to int: %v\n", value, err)
			continue // Skip invalid values
		}
		result = append(result, num)
	}

	return result
}

func TernaryOperatorFunc[T any](cond bool, value1, value2 T) T {
	if cond {
		return value1
	}
	return value2
}

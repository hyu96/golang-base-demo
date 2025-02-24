package islice

import "strconv"

type Int interface {
	int64 | int32
}

func ConvertToInt[I Int](vals []I) []int {
	result := make([]int, len(vals))
	for i := range vals {
		result[i] = int(vals[i])
	}

	return result
}

func ConvertToInt64[I int](vals []I) []int64 {
	result := make([]int64, len(vals))
	for i := range vals {
		result[i] = int64(vals[i])
	}

	return result
}

func SliceStringToInt(vals []string) ([]int, error) {
	result := make([]int, len(vals))
	var err error
	for i, item := range vals {
		result[i], err = strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

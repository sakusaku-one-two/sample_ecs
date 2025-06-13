package util

import (
	"encoding/json"
	"strconv"
)

// ----------------------JSON---------------------//

func StructToJson[M any](target M) []byte {
	byte_array, err := json.Marshal(target)
	if err != nil {
		return make([]byte, 0)
	}
	return byte_array
}

func MapToJson[Key_Type comparable, Value_Type any](target map[Key_Type]Value_Type) []byte {
	byte_array, err := json.Marshal(target)
	if err != nil {
		return make([]byte, 0)
	}
	return byte_array
}

func ByteToStruct[M any](target []byte) (M, error) {
	var reuslt M
	err := json.Unmarshal(target, &reuslt)
	return reuslt, err
}

func ByteToMap[Key_Type comparable](target []byte) map[Key_Type]any {
	var result = make(map[Key_Type]any)
	json.Unmarshal(target, &result)
	return result
}

// --------------------ConvertType--------------------//

func ToInt(from_string string) int {
	int_data, err := strconv.Atoi(from_string)
	if err != nil {
		return 0
	}
	return int_data
}

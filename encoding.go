package functional

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"strings"
)

func ReadAsCSV(val string) ([]string, error) {
	if val == "" {
		return []string{}, nil
	}
	stringReader := strings.NewReader(val)
	csvReader := csv.NewReader(stringReader)
	return csvReader.Read()
}

func ReadAsMap(val string) (map[string]string, error) {
	var newMap = make(map[string]string)
	Slice, err := ReadAsCSV(val)
	if err != nil {
		return nil, err
	}
	for _, str := range Slice { // iterating over each tab in the csv
		//map k:v are seperated by either = or : and then a comma
		strings.TrimSpace(str)
		if strings.Contains(str, "=") {
			newSlice := strings.Split(str, "=")
			newMap[newSlice[0]] = newSlice[1]
		}
		if strings.Contains(str, ":") {
			newSlice := strings.Split(str, ":")
			newMap[newSlice[0]] = newSlice[1]
		}
	}
	if newMap == nil {
		return nil, errors.New("cannot conver string to map[string]string- detected a nil map output")
	}
	return newMap, nil
}

// toJson encodes an item into a JSON string
func ToJson(v interface{}) string {
	output, _ := json.Marshal(v)
	return string(output)
}

// toPrettyJson encodes an item into a pretty (indented) JSON string
func ToPrettyJson(v interface{}) string {
	output, _ := json.MarshalIndent(v, "", "  ")
	return string(output)
}

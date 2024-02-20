package jsontoenv

import (
	"log"
	"strconv"
)

// parsePrimitives is a helper function to parse primitives
func parsePrimitives(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		log.Println("Unknown type")
		return ""
	}
}

// parseValuesWithObjects is a helper function to parse values with objects
func parseValuesWithObjects(key string, value interface{}) (map[string]string, error) {
	values := make(map[string]string)

	switch v := value.(type) {
	case map[string]interface{}:
		return parseObject(v)
	case []interface{}:
		values[key] = parseArray(v)
		return values, nil
	default:
		values[key] = parsePrimitives(v)
		return values, nil
	}
}

// parseObject is a helper function to parse objects
func parseObject(object map[string]interface{}) (map[string]string, error) {
	values := make(map[string]string)

	for key, value := range object {
		parsed, err := parseValuesWithObjects(key, value)
		if err != nil {
			return nil, err
		}

		for k, v := range parsed {
			values[k] = v
		}
	}

	return values, nil
}

// parseArray is a helper function to parse arrays
func parseArray(array []interface{}) string {
	var result string
	for _, value := range array {
		result += parsePrimitives(value) + ","
	}
	return result[:len(result)-1]
}

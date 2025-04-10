package conv

import (
	"encoding/json"
	"strings"
)

func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

func mapKeysToCamelCase(input any) any {
	switch v := input.(type) {
	case map[string]any:
		converted := make(map[string]any)
		for key, val := range v {
			converted[snakeToCamel(key)] = mapKeysToCamelCase(val)
		}
		return converted
	case []any:
		for i, item := range v {
			v[i] = mapKeysToCamelCase(item)
		}
	}
	return input
}

func SnakeJsonToCamelStruct(data []byte, out any) error {
	var raw any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	converted := mapKeysToCamelCase(raw)

	normalizedJSON, err := json.Marshal(converted)
	if err != nil {
		return err
	}

	return json.Unmarshal(normalizedJSON, out)
}

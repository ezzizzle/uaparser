package useragent

import "strings"

func csvRowToMap(values []string, headers []string) map[string]string {
	resultMap := map[string]string{}
	for index := range headers {
		resultMap[headers[index]] = strings.TrimSpace(values[index])
	}
	return resultMap
}

package core

// Create map from string array
func CreateMapFromStrings(sourceData []string) map[string]string {

	length := len(sourceData) //len() dont fail even with nil
	resultMap := make(map[string]string, length)

	for i := 0; i < length; i++ {
		itStr := sourceData[i]
		if itStr != "" {
			resultMap[itStr] = ""
		}
	}

	return resultMap
}

// Create string array from map keys
func CreateStringsFromMapKeys(sourceMap *map[string]string) []string {

	resultKeys := make([]string, 0, len(*sourceMap))
	for k := range *sourceMap {
		resultKeys = append(resultKeys, k)
	}

	return resultKeys
}

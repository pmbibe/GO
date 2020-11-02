func AnalystString(s1 string) map[string]int {
	duplicateFrequency := make(map[string]int)
	for _, item := range s1 {
		item := string(item)
		_, exist := duplicateFrequency[item] //Check if exist or not
		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}

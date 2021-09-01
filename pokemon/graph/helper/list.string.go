package helper

func getIndex(strList []string, find string) int {
	for idx, word := range strList {
		if word == find {
			return idx
		}
	}
	return -1
}

func delWord(strList []string, obj string) []string {
	idx := getIndex(strList, obj)
	if idx == -1 {
		return strList
	}
	return append(strList[:idx], strList[idx+1:]...)
}

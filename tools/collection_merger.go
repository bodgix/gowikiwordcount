package tools

func mergeTopN(wc []WordWithCount, n int) []WordWithCount {
	result := make([]WordWithCount, 0, n)
	result = append(result, wc[0])
	n--
	for _, curWord := range wc[1:] {
		if n == 0 {
			break
		}
		if result[len(result)-1].Count == curWord.Count {
			result[len(result)-1].Word = result[len(result)-1].Word + ", " + curWord.Word
		} else {
			result = append(result, curWord)
			n--
		}
	}
	return result
}

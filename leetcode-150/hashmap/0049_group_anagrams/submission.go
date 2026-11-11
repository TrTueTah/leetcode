package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		charArr := []rune(str)

		sort.Slice(charArr, func(i, j int) bool {
			return charArr[i] < charArr[j]
		})

		sortedKey := string(charArr)
		anagramGroups[sortedKey] = append(anagramGroups[sortedKey], str)
	}
	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

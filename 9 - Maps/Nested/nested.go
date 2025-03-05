package main

/* Assignment
Because Textio is a glorified customer database,
 we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names and returns a nested map.
 The parent map's keys are all the unique first characters
 (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.*/

func getNameCounts(names []string) map[rune]map[string]int {
	output := make(map[rune]map[string]int)
	for _, n := range names {
		firstRune := []rune(n)[0]
		if output[firstRune] == nil {
			output[firstRune] = make(map[string]int)

		}
		output[firstRune][n]++

	}
	return output
}

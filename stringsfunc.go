package main

func Split(s string, delimiter rune) []string {
	var result []string
	phrase := ""
	for idx, char := range s {
		if char == delimiter {
			result = append(result, phrase)
			phrase = ""
			continue
		}
		phrase += string(char)
		if idx == len(s)-1 {
			result = append(result, phrase)
		}
	}
	return result
}

func Join(s []string, delimiter string) string {
	var result string
	for i, r := range s {
		if i == len(s)-1 {
			result += r
			break
		}
		result += r + delimiter
	}
	return result
}

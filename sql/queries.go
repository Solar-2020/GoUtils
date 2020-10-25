package sql

import "strings"

func CreateIN(count int) (queryIN string) {
	queryIN = "("
	for i := 0; i < count; i++ {
		queryIN += "?, "
	}
	queryINRune := []rune(queryIN)
	queryIN = string(queryINRune[:len(queryINRune)-2])
	queryIN += ")"
	return
}

func CreateInsertQuery(sliceLen int, structLen int) (query string) {
	query = ""
	for i := 0; i < sliceLen; i++ {
		query += "("
		for j := 0; j < structLen; j++ {
			query += "?,"
		}
		// delete last comma
		query = strings.TrimRight(query, ",")
		query += "),"
	}
	// delete last comma
	query = strings.TrimRight(query, ",")

	return
}

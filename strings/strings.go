package strings

import "strconv"

func MustParseInt(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return id
}

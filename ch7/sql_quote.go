package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sqlQuote(v interface{}) string {
	switch v := v.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case string:
		return fmt.Sprintf(`"%s"`, v)
	case nil:
		return "NULL"
	default:
		return ""
	}
}

func FormatQuery(query string, args ...interface{}) string {
	for _, arg := range args {
		query = strings.Replace(query, "?", sqlQuote(arg), 1)
	}
	return query
}

func main() {
	fmt.Println(FormatQuery("SELECT * FROM t WHERE id = ? AND name = ? AND disabledAt = ?", 10, "yuki", nil))
}

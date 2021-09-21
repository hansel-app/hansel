package repositories

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

var qb = New()

func New() goqu.DialectWrapper {
	goqu.SetColumnRenameFunction(camelToSnake)
	var qb = goqu.Dialect("psql")
	return qb
}

// From https://stackoverflow.com/a/56616250
func camelToSnake(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	fmt.Printf("%s %s\n", s, snake)
	return strings.ToLower(snake)
}

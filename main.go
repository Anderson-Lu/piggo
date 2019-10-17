package main

import (
	"io"
	"log"
	"strings"

	"github.com/xwb1989/sqlparser"
)

func main() {
	r := strings.NewReader("INSERT INTO table1 VALUES (1, 'a'); INSERT INTO table2 VALUES (3, 4);")

	tokens := sqlparser.NewTokenizer(r)
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		switch stmt := stmt.(type) {
		case *sqlparser.Select:
			_ = stmt
		case *sqlparser.Insert:
			log.Println((*sqlparser.Insert)(stmt))
		}
	}
}

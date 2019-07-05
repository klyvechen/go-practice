package service

import (
	"database/sql"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func chgItf2Row(srcs []interface{}) []interface{} {
	result := []interface{}{}
	for _, src := range srcs {
		result = append(result, chgSql2Native(src))
	}
	return result
}

func chgSql2Native(src interface{}) interface{} {
	switch v := src.(type) {
	case *int:
		return *v
	case *sql.NullString:
		return (*v).String
	case *sql.NullInt64:
		return (*v).Int64
	default:
		return v
	}
}

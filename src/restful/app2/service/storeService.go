package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful/app2/model/table"

	_ "github.com/go-sql-driver/mysql"
)

func GetStoreData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	fmt.Println(connectionString)
	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	result := table.TableData{
		Thead: table.Thead{Ths: []string{"#", "First Name", "Last Name", "Username"}},
		Tbody: table.Tbody{Trs: []table.Tr{
			table.Tr{Tds: []interface{}{"1", "Mark", "Otto", "@mdo"}},
			table.Tr{Tds: []interface{}{"1", "Mark", "Otto", "@TwBootstrap"}},
			table.Tr{Tds: []interface{}{"2", "Jacob", "Thornton", "@fat"}},
			table.Tr{Tds: []interface{}{"3", "Larry the Bird", "", "@twitter"}},
		}},
	}

	jsonresult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintln(w, string(jsonresult))
}

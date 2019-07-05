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

const (
	host     = "127.0.0.1"
	database = "LianJunDb"
	user     = "root"
	password = ""
)

var (
	result = table.TableData{
		Thead: table.Thead{Ths: []string{"#", "First Name", "Last Name", "Username"}},
		Tbody: table.Tbody{Trs: []table.Tr{
			table.Tr{Tds: []interface{}{"1", "Mark", "Otto", "@mdo"}},
			table.Tr{Tds: []interface{}{"1", "Mark", "Otto", "@TwBootstrap"}},
			table.Tr{Tds: []interface{}{"2", "Jacob", "Thornton", "@fat"}},
			table.Tr{Tds: []interface{}{"3", "Larry the Bird", "", "@twitter"}},
		}},
	}
	row1  = table.Tr{Tds: []interface{}{"", "", "", ""}}
	row2  = table.Tr{Tds: []interface{}{"", "", ""}}
	tbody = table.Tbody{Trs: []table.Tr{row1, row2}}
	thead = table.Thead{}

	result1 table.TableData = table.TableData{
		Thead: thead,
		Tbody: tbody,
	}
)

type Product struct {
	id       int
	prodNo   sql.NullString // prod_no
	prodName sql.NullString // prod_name
	amount   sql.NullInt64  // amount
	pic      sql.NullString // pic
	size     sql.NullString // size
}

func GetInventoryData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Initialize connection string.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	fmt.Println(connectionString)
	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for reading data
	rows, err := db.Query("SELECT * FROM Product")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	result := table.TableData{
		Thead: table.Thead{Ths: []string{"#", "商品代號", "商品名稱", "庫存", "照片", "尺寸"}},
		Tbody: table.Tbody{Trs: []table.Tr{}},
	}

	for rows.Next() {

		dest := []interface{}{ // Standard MySQL columns
			new(int),            // id
			new(sql.NullString), // prod_no
			new(sql.NullString), // prod_name
			new(sql.NullInt64),  // amount
			new(sql.NullString), // pic
			new(sql.NullString), // size
		}
		// product := Product{}
		// if err := rows.Scan(&product.id, &product.prodNo, &product.prodName, &product.amount, &product.pic, &product.size); err != nil {
		// 	log.Fatal(err)
		// }
		// row := []interface{}{product.id, product.prodNo.String, product.prodName.String, product.size.String, product.amount.Int64}
		// result.Tbody.Trs = append(result.Tbody.Trs, table.Tr{Tds: row})
		if err := rows.Scan(dest...); err != nil {
			log.Fatal(err)
		}
		row := chgItf2Row(dest)
		result.Tbody.Trs = append(result.Tbody.Trs, table.Tr{Tds: row})
		fmt.Printf("%s\n", dest)
	}

	jsonresult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintln(w, string(jsonresult))
}

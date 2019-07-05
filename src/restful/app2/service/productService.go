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

func GetProductManagement(w http.ResponseWriter, r *http.Request) {
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

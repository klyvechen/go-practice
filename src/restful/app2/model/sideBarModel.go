package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SideBar struct {
	Type     string    `json:"type"`
	Href     string    `json:"href"`
	Icon     string    `json:"icon"`
	Name     string    `json:"name"`
	Children []SideBar `json:"children"`
}

func GetSideBars(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	result := []SideBar{
		{"active", "index", "Icon_house_alt", "Example", []SideBar{}},
		{"sub-menu", "javascript:;", "icon_document_alt", "採購", []SideBar{
			{"", "form_component.html", "", "採購單預購", []SideBar{}},
			{"", "form_validation.html", "", "採購單分析", []SideBar{}},
		}},
		{"sub-menu", "javascript:;", "Icon_document_alt", "廠商管理", []SideBar{
			{"active", "supplier", "", "供應商管理", []SideBar{}},
			{"active", "store", "", "店家管理", []SideBar{}},
		}},
		{"active", "supplier", "", "供應商管理", []SideBar{}},
		{"sub-menu", "javascript:;", "Icon_desktop", "商品", []SideBar{
			{"active", "inventory", "", "分類一", []SideBar{}},
			{"", "productManagement", "", "分類二", []SideBar{}},
			{"", "grids.html", "", "分類三", []SideBar{}},
		}},
		{"active", "widgets.html", "Icon_genius", "庫存管理", []SideBar{}},
		{"active", "topology", "Icon_genius", "Topology", []SideBar{}},
		{"active", "chart-chartjs.html", "Icon_piechart", "Charts", []SideBar{}},
		{"active", "basic_table.html", "Icon_table", "Basic Table", []SideBar{}},
		// {"sub-menu", "javascript:;", "icon_document_alt", "採購", nil},
		// {"sub-menu", "javascript:;", "icon_document_alt", "採購", nil},
		// {"sub-menu", "javascript:;", "icon_document_alt", "採購", nil},
	}

	//    SideBar {Type: "active", Href: "widgets.html", Icon: "Icon_genius", Name: "Widgets", Children: []},
	//    SideBar {Type: "active", Href: "chart-chartjs.html", Icon: "Icon_piechart", Name: "Charts", Children: []},
	//    SideBar {Type: "sub-menu", Href: "basic_table.html", Icon: "Icon_table", Name: "Basic Table", Children: []},
	//    SideBar {Type: "sub-menu", Href: "javascript:;", Icon: "Icon_documents_alt", Name: "Pages", Children: [
	//       SideBar {Type: "", Href: "profile.html", Icon: "", Name: "Profile", Children: []},
	//       SideBar {Type: "", Href: "login.html", Icon: "", Name: "Login Page", Children: []},
	//       SideBar {Type: "", Href: "contact.html", Icon: "", Name: "Contact Page", Children: []},
	//       SideBar {Type: "", Href: "blank.html", Icon: "", Name: "Blank Page", Children: []},
	//       SideBar {Type: "", Href: "404.html", Icon: "", Name: "404 Error", Children: []}
	//    ]},
	// arr := JsonType{}
	jsonresult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintln(w, string(jsonresult))
	//
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Test() {

}

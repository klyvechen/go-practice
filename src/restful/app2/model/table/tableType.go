package table

type Tbody struct {
	Trs []Tr `json:"tr"`
}

type Tr struct {
	Tds []interface{} `json:"td"`
}

type Thead struct {
	Ths []string `json:"th"`
}

type TableData struct {
	Thead Thead `json:"thead"`
	Tbody Tbody `json:"tbody"`
}

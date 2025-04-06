package domain

type Table struct {
	TableName string          `json:"table_name"`
	Columns   []Column        `json:"columns"`
	Rows      [][]interface{} `json:"rows"`
}

type Column struct {
	Name      string `json:"name"`
	DataType  string `json:"data_type"`
	IsPrimary bool   `json:"is_primary"`
}

package adminservice

import (
	"encoding/json"
	"fmt"
	"os"
	// domain paketini import ediyoruz
)

type Animal struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	OwnerID  int     `json:"ownerid"`
	Nickname string  `json:"nickname,omitempty"` // Eğer varsa takma ad

}

func ListAllAnimals(filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	type Column struct {
		Name      string `json:"name"`
		DataType  string `json:"data_type"`
		IsPrimary bool   `json:"is_primary"`
	}
	type Table struct {
		TableName string          `json:"table_name"`
		Columns   []Column        `json:"columns"`
		Rows      [][]interface{} `json:"rows"`
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	for _, row := range table.Rows {
		animal := Animal{
			ID:       int(row[0].(float64)),
			Name:     row[1].(string),
			Type:     row[2].(string),
			Price:    row[3].(float64),
			OwnerID:  int(row[4].(float64)),
			Nickname: row[5].(string),
		}

		// Animal bilgilerini yazdırma
		fmt.Printf("ID: %d\n", animal.ID)
		fmt.Printf("Name: %s\n", animal.Name)
		fmt.Printf("Type: %s\n", animal.Type)
		fmt.Printf("Price: %.2f\n", animal.Price)
		fmt.Printf("OwnerID: %d\n", animal.OwnerID)
		fmt.Printf("Nickname: %s\n", animal.Nickname)
		fmt.Println("-----------") // Satır ayırıcı
	}

	return nil
}

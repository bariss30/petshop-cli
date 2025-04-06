package userservice

import (
	"encoding/json"
	"fmt"
	"os"

	
)

type Animal struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	OwnerID  int     `json:"ownerid"`
	Nickname string  `json:"nickname,omitempty"` // Eğer varsa takma ad
}

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

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

func ListAllAnimals(filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
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

		fmt.Printf("ID: %d\n", animal.ID)
		fmt.Printf("Name: %s\n", animal.Name)
		fmt.Printf("Type: %s\n", animal.Type)
		fmt.Printf("Price: %.2f\n", animal.Price)
		fmt.Printf("OwnerID: %d\n", animal.OwnerID)
		fmt.Printf("Nickname: %s\n", animal.Nickname)
		fmt.Println("--------")
	}

	return nil
}

func ListOwnedAnimals(userID int, filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	// Kullanıcının sahip olduğu hayvanları listeleme
	for _, row := range table.Rows {
		ownerID := int(row[4].(float64))
		if ownerID == userID {
			animal := Animal{
				ID:       int(row[0].(float64)), // ID'yi float64'dan int'e dönüştürme
				Name:     row[1].(string),       // Name string olarak
				Type:     row[2].(string),       // Type string olarak
				Price:    row[3].(float64),      // Price'ı float64 olarak
				OwnerID:  ownerID,               // OwnerID'yi doğrudan alıyoruz
				Nickname: row[5].(string),       // Nickname string olarak
			}
			// Sahip olunan hayvan bilgilerini yazdırma
			fmt.Printf("ID: %d\n", animal.ID)
			fmt.Printf("Name: %s\n", animal.Name)
			fmt.Printf("Type: %s\n", animal.Type)
			fmt.Printf("Price: %.2f\n", animal.Price)
			fmt.Printf("OwnerID: %d\n", animal.OwnerID)
			fmt.Printf("Nickname: %s\n", animal.Nickname)
			fmt.Println("-----------") // Satır ayırıcı
		}
	}

	return nil
}

func AdoptAnimal(userID int, animalID int, filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	// Hayvanı bulup sahibini güncelleme
	for i, row := range table.Rows {
		if int(row[0].(float64)) == animalID {
			table.Rows[i][4] = userID // OwnerID'yi güncelle
			break
		}
	}

	updatedJSON, err := json.Marshal(table)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, updatedJSON, 0644)
}

func NameAnimal(animalID int, name string, filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	// Hayvanın ismini güncelleme
	for i, row := range table.Rows {
		if int(row[0].(float64)) == animalID {
			table.Rows[i][5] = name // Nickname'yi güncelle
			break
		}
	}

	updatedJSON, err := json.Marshal(table)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, updatedJSON, 0644)
}

func ListAllItems(filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	// Bütün market ürünlerini listeleme
	for _, row := range table.Rows {
		item := Item{
			ID:    int(row[0].(float64)), // ID'yi float64'dan int'e dönüştürme
			Name:  row[1].(string),       // Name string olarak
			Price: row[2].(float64),      // Price'ı float64 olarak
		}

		// Market ürün bilgilerini yazdırma
		fmt.Printf("ID: %d\n", item.ID)
		fmt.Printf("Name: %s\n", item.Name)
		fmt.Printf("Price: %.2f\n", item.Price)
		fmt.Println("-----------") // Satır ayırıcı
	}

	return nil
}

func PurchaseItem(userID int, itemID int, filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	// Ürün satın alma işlemi
	for _, row := range table.Rows {
		if int(row[0].(float64)) == itemID {
			// Ürün fiyatını ve kullanıcının bakiyesini kontrol et
			price := row[2].(float64)
			// Bakiye kontrolü ve para yükleme mantığı burada yapılabilir.
			// Burada sadece basit bir çıktı veriyoruz.
			fmt.Printf("Ürün satın alındı: %s, Fiyat: %.2f\n", row[1].(string), price)
			break
		}
	}

	return nil
}

func AddBalance(userID int, amount float64, filePath string) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var table Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		return err
	}

	// Kullanıcı bakiyesini güncelleme
	for i, row := range table.Rows {
		if int(row[0].(float64)) == userID {
			table.Rows[i][4] = row[4].(float64) + amount // Balance'yi güncelle
			break
		}
	}

	updatedJSON, err := json.Marshal(table)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, updatedJSON, 0644)
}

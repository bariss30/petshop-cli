package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	userservice "petshop-cli/service/userService"

	"strings"

	dbmodule "github.com/bariss30/go-module-json-based-database/dbmodule"
	"golang.org/x/crypto/argon2"
)

type User struct {
	Id       int
	Username string
	Password string
	UserType string
	Balance  float64
}

var users []User      // Kullanıcı listesi
var currentUser *User // Geçerli kullanıcı

/*func init() {
	err := RegisterUser("Database/users.json", "admin", "admin", "admin", 100000000000.0)

	if err != nil {
		fmt.Println("Kayıt başarısız:", err)
	} else {
		fmt.Println("Kayıt başarılı!")
	}
}*/

func AdminPage() {
	fmt.Println("\nAdmin Sayfasına Hoş Geldiniz!")

	for {
		fmt.Println("\nAdmin Menüsü:")
		fmt.Println("1. Müşteri Ekle")
		fmt.Println("2.Müşteri Sil,Hayvan Ekle,Hayvan Sil,Hayvan Güncelle,Hayvan Listele,İtem Ekle,İtem Sil")
		fmt.Println("3. Para Yükle")
		fmt.Println("4. Update")
		fmt.Println("5. Çıkış")

		var choice string
		fmt.Print("Seçiminiz: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("Kayıt Yapılıyor...")
			var username, password string

			fmt.Print("Kullanıcı Adı: ")
			fmt.Scan(&username)

			fmt.Print("Şifre: ")
			fmt.Scan(&password)

			err := RegisterUser("Database/users.json", username, password, "customer", 0.0)
			if err != nil {
				log.Println("Kayıt yapılamadı:", err)
			} else {
				log.Println("Kayıt başarılı!")
			}

		case "2":
			dbmodule.CreateTable()
		case "3":
			//addbalance
		case "4":
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func CustomerPage() {
	fmt.Println("\nMüşteri Sayfasına Hoş Geldiniz!")

	for {
		fmt.Println("\nMüşteri Menüsü:")

		fmt.Println("1. Hayvan Sahiplen")
		fmt.Println("2. Hayvana İsim Tak")
		fmt.Println("3. Kendi Hayvanlarını Listele")
		fmt.Println("4. Tüm Hayvanları Gör")
		fmt.Println("5. Marketten Satın Al")
		fmt.Println("6. Satın Alınan Ürünleri Gör")
		fmt.Println("7. Tüm Market Ürünlerini Gör")
		fmt.Println("8. Para Yükle")
		fmt.Println("9. Çıkış")

		var choice string
		fmt.Print("Seçiminiz: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			userservice.AdoptAnimal(2, 1, "Database/animals.json")
		case "2":
			userservice.NameAnimal(1, "canavar", "Database/animals.json")
		case "3":
			userservice.ListOwnedAnimals(2, "Database/animals.json")

		case "4":
			filePath := "Database/animals.json"

			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				fmt.Println("Hata: Dosya bulunamadı:", filePath)
				return
			}

			// Dosyayı okuma işlemi ve hata kontrolü
			err := userservice.ListAllAnimals(filePath)
			if err != nil {
				fmt.Println("Hayvanları listeleme sırasında hata oluştu:", err)
				return
			}

			fmt.Println("Tüm hayvanlar başarıyla listelendi.")
		case "5":
			// purchaseItem()
		case "6":
			// listPurchasedItems()
		case "7":
			// listAllItems()
		case "8":
			userservice.AddBalance(2, 500.0, "Database/users.json")
		case "9":
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func LoginAsAdmin() {
	var username, password string

	fmt.Print("Admin Kullanıcı adı: ")
	fmt.Scan(&username)
	fmt.Print("Şifre: ")
	fmt.Scan(&password)

	if username == "selam" && password == "selam" {

		AdminPage()

	}
}

func RegisterUser(filePath string, username, password, userType string, balance float64) error {
	tableFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return err
	}

	var table dbmodule.Table
	if err := json.Unmarshal(tableFile, &table); err != nil {
		fmt.Println("JSON parse hatası:", err)
		return err
	}

	newId := len(table.Rows) + 1

	hashedPassword := GenerateArgon2Hash(password)

	primaryKeys := []interface{}{newId}
	otherValues := []interface{}{username, hashedPassword, userType, balance}

	newRow := append(primaryKeys, otherValues...)
	table.Rows = append(table.Rows, newRow)

	updatedJson, err := json.Marshal(table)
	if err != nil {
		fmt.Println("JSON oluşturulamadı:", err)
		return err
	}

	return os.WriteFile(filePath, updatedJson, 0644)
}

func GenerateArgon2Hash(password string) string {

	salt := []byte("somesalt")

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return fmt.Sprintf("%x", hash)
}

func VerifyPassword(password, hashedPassword string) bool {
	salt := []byte("somesalt")
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return fmt.Sprintf("%x", hash) == hashedPassword
}

func LoginUser(filePath, username, password string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return err
	}

	var table dbmodule.Table
	if err := json.Unmarshal(file, &table); err != nil {
		fmt.Println("JSON parse hatası:", err)
		return err
	}

	for _, row := range table.Rows {
		dbUsername := fmt.Sprintf("%v", row[1])
		dbPassword := fmt.Sprintf("%v", row[2])
		dbUserType := fmt.Sprintf("%v", row[3])
		dbBalance := fmt.Sprintf("%v", row[4])

		if strings.ToLower(dbUsername) == strings.ToLower(username) {
			if VerifyPassword(password, dbPassword) {
				fmt.Printf("Giriş başarılı! Hoş geldiniz, %s (%s). Bakiyeniz: %s TL\n", username, dbUserType, dbBalance)
				currentUser = &User{
					Id:       int(row[0].(float64)),
					Username: dbUsername,
					Password: dbPassword,
					UserType: dbUserType,
					Balance:  row[4].(float64),
				}
				CustomerPage()
				return nil
			} else {
				fmt.Println("Hatalı şifre!")
				return fmt.Errorf("şifre yanlış")
			}
		}
	}

	fmt.Println("Kullanıcı bulunamadı!")
	return fmt.Errorf("kullanıcı bulunamadı")
}

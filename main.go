package main

import (
	"fmt"
	"log"
	"os"
	"petshop-cli/service/auth"
	/*userservice "petshop-cli/service/userService"*/)

func main() {
	for {
		fmt.Println("1. Üye Ol")
		fmt.Println("2. Giriş Yap")
		fmt.Println("3. Admin Girişi Yap")
		fmt.Println("4. Çıkış")
		fmt.Print("Seçiminiz: ")

		var secim int
		fmt.Scan(&secim)

		switch secim {
		case 0:
			/*cat := &domain.Cat{
				ID:         1,
				Name:       "Whiskers",
				AnimalType: "Cat",
				Price:      100.0,
				OwnerID:    123,
			}
			fmt.Println("Cat's Name:", cat.GetName())*/

			/* userservice.ListOwnedAnimals(2, "Database/animals.json") */

			/*userservice.AdoptAnimal(2, 1, "Database/animals.json")*/

			/*userservice.ListAllAnimals("Database/animals.json")*/

			/*userservice.NameAnimal(1, "canavar", "Database/animals.json")*/

			/*userservice.AddBalance(2, 500.0, "Database/users.json")*/

		case 1:
			fmt.Println("Kayıt Yapılıyor...")
			var username, password string

			fmt.Print("Kullanıcı Adı: ")
			fmt.Scan(&username)

			fmt.Print("Şifre: ")
			fmt.Scan(&password)

			err := auth.RegisterUser("Database/users.json", username, password, "customer", 0.0)
			if err != nil {
				log.Println("Kayıt yapılamadı:", err)
			} else {
				log.Println("Kayıt başarılı!")
			}

		case 2:
			fmt.Println("Giriş Yapılıyor...")
			var username, password string

			fmt.Print("Kullanıcı Adı: ")
			fmt.Scan(&username)

			fmt.Print("Şifre: ")
			fmt.Scan(&password)

			err := auth.LoginUser("Database/users.json", username, password)
			if err != nil {
				log.Println("Giriş yapılamadı:", err)
			} else {
				log.Println("Giriş başarılı!")
			}

		case 3:
			fmt.Println("Admin Giriş İşlemi...")

			auth.LoginAsAdmin()

		case 4:
			os.Exit(0)

		default:
			fmt.Println("Geçersiz Seçim")
		}
	}
}

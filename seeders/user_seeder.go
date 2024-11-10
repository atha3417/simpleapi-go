package seeders

import (
	"fmt"
	"log"
	"simpleapi/helpers"
	"simpleapi/models"
	repo "simpleapi/repositories"
	"strconv"

	"github.com/go-faker/faker/v4"
)

// Fungsi untuk menambahkan data palsu ke tabel Users
func UserSeeder(count int) {
	for i := 0; i < count; i++ {
		// Buat data Name terlebih dahulu
		name := models.Name{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
		}

		// Simpan Name dan ambil ID-nya
		if err := repo.CreateName(&name); err != nil {
			log.Printf("Gagal menambahkan name: %v", err)
			continue
		}

		// Buat data Geolocation
		geolocation := models.Geolocation{
			Latitude:  faker.GetRealAddress().Coordinates.Latitude,
			Longitude: faker.GetRealAddress().Coordinates.Longitude,
		}

		// Simpan Geolocation
		if err := repo.CreateGeolocation(&geolocation); err != nil {
			log.Printf("Gagal menambahkan geolocation: %v", err)
			continue
		}

		// Buat data Address yang terkait dengan Geolocation
		zipCode, _ := strconv.Atoi(faker.GetRealAddress().PostalCode)
		address := models.Address{
			City:       faker.GetRealAddress().City,
			Street:     faker.GetRealAddress().Address,
			Number:     helpers.RandomNumberBetween(1, 1000),
			ZipCode:    zipCode,
			Geolocation: geolocation, // Menghubungkan Geolocation ke Address
		}

		// Simpan Address
		if err := repo.CreateAddress(&address); err != nil {
			log.Printf("Gagal menambahkan address: %v", err)
			continue
		}

		// Buat data User dengan relasi ke Name dan Address
		user := models.User{
			Email:    faker.Email(),
			Username: faker.Username(),
			Password: faker.Password(),
			Phone:    faker.Phonenumber(),
			Name:     name,
			Address:  address,
		}

		// Simpan User
		if err := repo.CreateUser(&user); err != nil {
			log.Printf("Gagal menambahkan user: %v", err)
		} else {
			fmt.Printf("User %d berhasil ditambahkan: %v\n", i+1, user)
		}
	}
}
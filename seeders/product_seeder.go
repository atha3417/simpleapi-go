package seeders

import (
	"fmt"
	"log"
	"simpleapi/helpers"
	"simpleapi/models"
	repo "simpleapi/repositories"
)

func ProductSeeder(count int) {
	for i := 0; i < count; i++ {
		rating := models.Rating{
			Rate: helpers.RandomFloatNumberBetween(1, 10),
			Count: helpers.RandomNumberBetween(100, 1000),
		}

		if err := repo.CreateRating(&rating); err != nil {
			log.Printf("Gagal menambahkan rating: %v", err)
			continue
		}

		product := models.Product{
			Title: helpers.RandomSentenceBetween(5, 15),
			Price: helpers.RandomNumberBetween(10000, 1000000),
			Description: helpers.RandomSentenceBetween(15, 30),
			Category: helpers.RandomCategory(),
			Image: "https://api.api-ninjas.com/v1/randomimage",
			Rating: rating,
		}

		if err := repo.CreateProduct(&product); err != nil {
			log.Printf("Gagal menambahkan product: %v", err)
		} else {
			fmt.Printf("Product %d berhasil ditambahkan: %v\n", i+1, product)
		}
	}
}
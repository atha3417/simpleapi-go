package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
)

// Inisialisasi sumber acak
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomStreetNumber menghasilkan nomor jalan acak dalam bentuk string
func RandomNumberBetween(start, end int) int {
	return rng.Intn(end-start)
}

func RandomFloatNumberBetween(start, end float64) float64 {
	r, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", start + rng.Float64()*(end-start)), 64)
    return r
}

func RandomSentenceBetween(minWords, maxWords int) string {
    var sentence string
    for {
        sentence = faker.Sentence()
        wordCount := len(strings.Fields(sentence))

        if wordCount >= minWords && wordCount <= maxWords {
            break
        }
    }
    return sentence
}

func RandomCategory() string {
    var categories = []string{"Food", "Drink", "Clothes", "Electronics", "Books", "Toys", "Furniture", "Tools", "Health", "Beauty", "Sports", "Automotive", "Music", "Movies", "Games", "Garden", "Pet", "Baby", "Jewelry", "Shoes", "Bags", "Accessories", "Software", "Hardware", "Phones", "Computers", "Tablets", "Cameras", "TVs", "Audio", "Smartwatches", "Wearable", "Home", "Kitchen", "Office", "Outdoor", "Travel", "Luggage", "Vehicles", "Bikes", "Motorcycles", "Boats", "Planes", "Houses", "Apartments", "Land", "Commercial", "Industrial", "Services", "Jobs", "Events", "Courses", "Tutoring", "Lessons", "Classes", "Workshops", "Conferences", "Meetups", "Webinars", "Seminars", "Festivals", "Expos", "Fairs", "Shows", "Concerts", "Plays", "Movies", "Games", "Sports", "Competitions", "Tournaments"}
    return categories[rng.Intn(len(categories))]
}
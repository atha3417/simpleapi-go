package seeders

import (
	"fmt"
	"log"
	"simpleapi/config"
)

func Seed() {
	// Mengosongkan tabel sebelum seeding
	// truncateAllTables()

	// Menjalankan seeders jika diperlukan
	// UserSeeder(100)
	ProductSeeder(1000)
}

// Fungsi untuk menghapus semua data table
func truncateAllTables() {
	// Menghapus dan mereset kolom ID semua table
	dbs := []string{"users", "names", "addresses", "geolocations", "rating", "product"}

	for _, db := range dbs {
		if err := config.DB.Exec("DELETE FROM " + db).Error; err != nil {
			log.Fatalf("Gagal mengosongkan tabel %s: %v", db, err)
		} else {
			fmt.Printf("Tabel %s berhasil dikosongkan\n", db)
		}

		// Mereset sequence untuk kolom ID agar mulai dari 1 lagi
		if err := config.DB.Exec("ALTER SEQUENCE " + db + "_id_seq RESTART WITH 1").Error; err != nil {
			log.Fatalf("Gagal mereset sequence: %v", err)
		} else {
			fmt.Printf("Sequence untuk %s telah di-reset ke 1\n", db)
		}
	}
}
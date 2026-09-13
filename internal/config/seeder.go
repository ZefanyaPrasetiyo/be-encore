package config 

import (
	"log"
	"gorm.io/gorm"
	"encore-be/internal/models"
)

func strPtr(s string) *string {
	return &s
}

func SeedUsers(db *gorm.DB)  {
	var count int64

	db.Model(&models.User{}).Count(&count)

	if count > 0 {
		log.Println("Data user sudah ada, skip...")
		return
	}

	dummyPasswordHash := "$2a$10$wN1Q/X3hX.q9z.P2/x/q.e/q/q.e.q.e.q.e.q.e.q.e.q.e.q.e"

	users := []models.User {
		{
			FullName:     "Super Admin",
			Email:        "admin@tiket.com",
			PasswordHash: strPtr(dummyPasswordHash),
			Role:         models.RoleAdmin,
		},
		{
			FullName:     "Staff Penjaga",
			Email:        "staff@tiket.com",
			PasswordHash: strPtr(dummyPasswordHash),
			Role:         models.RoleStaff,
		},
		{
			FullName:     "Budi Pembeli",
			Email:        "budi@gmail.com",
			PasswordHash: strPtr(dummyPasswordHash),
			Role:         models.RoleBuyer,
		},
	}
	// Insert banyak data sekaligus (Bulk Insert)
	if err := db.Create(&users).Error; err != nil {
		log.Println("Gagal menjalankan seeder user: ", err)
		return
	}
	log.Println("Seeder user berhasil! 3 data awal telah ditambahkan.")
}
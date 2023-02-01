package model

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Task{})
	if err != nil {
		return
	}
}

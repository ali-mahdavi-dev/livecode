package scripts

import (
	"encoding/json"
	"fmt"
	"livecode/internal/adapter/repository"
	"livecode/internal/domain/entity"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func getUsersFromJSON(filename string) (*[]entity.User, error) {
	var users []entity.User
	userByte, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	err = json.Unmarshal(userByte, &users)

	return &users, err
}

func saveToDB(data []entity.User, userRepo repository.UserRepository) error {
	if err := userRepo.BulkInsert(data); err != nil {
		return fmt.Errorf("save to db: %w", err)
	}

	return nil
}

func dbConnect(host, username, password, dbname string, ssl bool) (*gorm.DB, error) {
	sslStr := "disable"
	if ssl {
		sslStr = "enable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		username,
		password,
		dbname,
		sslStr)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func RunScriptReadUserFromJSON(fileName string) error {
	users, err := getUsersFromJSON(fileName)
	if err != nil {
		panic(err)
	}

	db, err := dbConnect("localhost", "ali", "", "livecode", false)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)

	err = saveToDB(*users, userRepo)
	if err != nil {
		panic(err)
	}

	return nil
}

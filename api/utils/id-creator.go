package utils

import (
	"errors"
	"url-shortner/api/config"
	"url-shortner/api/models"

	gonanoid "github.com/matoous/go-nanoid"
	"gorm.io/gorm"
)

// IdCreator generates a unique short ID and assigns it to new_id
func IdCreator(new_id *string) error {
	var url models.Url

	// Generate a new ID
	id, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	if err != nil {
		return err // Return the error if ID generation fails
	}

	// Check if the generated ID already exists in the database
	result := config.DB.Where("short_id = ?", id).First(&url)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// If there's a database error, return it
		return result.Error
	}

	// If the generated ID doesn't exist, assign it to new_id
	if result.RowsAffected == 0 {
		*new_id = id // Dereference the pointer and assign the value
		return nil   // Return nil to indicate success
	}

	// If the ID exists, optionally retry or handle it (e.g., recursion)
	return IdCreator(new_id)
}

package db

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ResetPassword struct {
	ID              int    `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type CreateReset struct {
	Email string `json:"email"`
}

// TODO: Build upon this user struct
type User struct {
	ID        string    `json:"id,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Assets    []Asset   `json:"assets,omitempty"`
}

type Asset struct {
	AssetID           int          `json:"asset_id,omitempty"`
	UserID            string       `json:"user_id,omitempty"`
	AssetName         string       `json:"asset_name,omitempty"`
	PurchasePrice     float64      `json:"purchase_price,omitempty"`
	PurchaseDate      time.Time    `json:"purchase_date,omitempty"`
	CurrentPrice      float64      `json:"current_price,omitempty"`
	CurrentDate       time.Time    `json:"current_date,omitempty"`
	AssetPriceHistory []AssetPrice `json:"asset_price_history,omitempty"`
}

type AssetPrice struct {
	Price   float64   `json:"price,omitempty"`
	Date    time.Time `json:"date,omitempty"`
	AssetID int       `json:"asset_id,omitempty"` // Store the asset's ID
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) UserExists(dbConn *sql.DB) bool {
	rows, err := dbConn.Query(GetUserByEmailQuery, user.Email)
	if err != nil || !rows.Next() {
		return false
	}

	return true
}

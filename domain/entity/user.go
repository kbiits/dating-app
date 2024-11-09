package entity

import "time"

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserPurchase struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	PackageID    string    `db:"package_id"`
	PurchaseDate time.Time `db:"purchase_date"`
	IsActive     bool      `db:"is_active"`
}

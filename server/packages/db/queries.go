package db

const (
	CheckUserExists         = `SELECT true from users WHERE email = $1`
	LoginQuery              = `SELECT * from users WHERE email = $1`
	UpdateUserPasswordQuery = `UPDATE users SET password = $2 WHERE id = $1`
	DeleteUser              = `DELETE FROM users WHERE email = $1`
	CreateUserQuery         = `INSERT INTO users(id, name, password, email) VALUES (DEFAULT, $1 , $2, $3);`
	GetUserByIDQuery        = `SELECT * FROM users WHERE id = $1`
	GetUserByEmailQuery     = `SELECT * FROM users WHERE email = $1`
	CreateUserAssetQuery    = `INSERT INTO assets(asset_id, user_id, asset_name, purchase_price, purchase_date, current_price, "current_date") VALUES (DEFAULT, $1, $2, $3, $4, $5, $6)`
	GetAssetByID            = `SELECT * FROM assets WHERE asset_id = $1`
)

//type Asset struct {
//	AssetID            int          `json:"asset_id,omitempty"`
//	AssetName          string       `json:"asset_name,omitempty"`
//	PurchasePrice      float64      `json:"purchase_price,omitempty"`
//	PurchaseDate       time.Time    `json:"purchase_date,omitempty"`
//	CurrentPrice       float64      `json:"current_price,omitempty"`
//	CurrentDate        time.Time    `json:"current_date,omitempty"`
//	PriceChangeHistory []AssetPrice `json:"price_change_history,omitempty"`
//}

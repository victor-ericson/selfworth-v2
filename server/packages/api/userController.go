package api

import (
	"database/sql"
	"goapp/packages/config"
	"goapp/packages/db"
	"goapp/packages/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	db.User
	jwt.StandardClaims
}

func Pong(c *fiber.Ctx) error {
	return c.SendString("pong")
}

// User creation
func CreateUser(c *fiber.Ctx, dbConn *sql.DB) error {
	//creates new user in DB
	user := &db.User{}

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if errs := utils.ValidateUser(*user); len(errs) > 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"success": false, "errors": errs})
	}
	//if user already exists
	if user.UserExists(dbConn) {
		return c.Status(400).JSON(&fiber.Map{"success": false, "errors": []string{"email already exists"}})
	}

	//password hashing using bcrypt
	user.HashPassword()
	_, err := dbConn.Query(db.CreateUserQuery, user.Name, user.Password, user.Email)
	if err != nil {
		return err
	}
	return c.JSON(&fiber.Map{"success": true})
}

func GetUser(c *fiber.Ctx, dbConn *sql.DB) error {
	tokenUser := c.Locals("user").(*jwt.Token)
	claims := tokenUser.Claims.(jwt.MapClaims)
	userID, ok := claims["id"].(string)

	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	// Fetch user information (following GetUser semantics)
	user := &db.User{}
	err := dbConn.QueryRow(db.GetUserByIDQuery, userID).
		Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"success": false, "errors": []string{"Incorrect credentials"}})
		}
		return err // You might want to handle other database errors differently
	}
	user.Password = ""

	// Fetch assets data
	//TODO: TEST IF WORKS
	assets, err := fetchAssetsByUserID(dbConn, userID)
	if err != nil {
		return err
	}

	// Create a combined response structure
	response := fiber.Map{"success": true, "user": user, "assets": assets}
	return c.JSON(response)
}

// TODO: TEST IF WORKS
func fetchAssetsByUserID(dbConn *sql.DB, userID string) ([]db.Asset, error) {
	//creates an assets array
	assets := []db.Asset{}

	//iterates through every asset in database with the userID
	rows, err := dbConn.Query(db.GetAllAssetsQuery, userID)
	if err != nil {
		return assets, err
	}
	defer rows.Close()

	//for each row, scan the asset and append it to the assets array
	for rows.Next() {
		asset := db.Asset{}
		if err := rows.Scan(&asset.AssetID, &asset.AssetName, &asset.CurrentPrice, &asset.CurrentDate, &asset.PurchasePrice, &asset.PurchaseDate, &asset.AssetPriceHistory); err != nil {
			return assets, err
		}
		assets = append(assets, asset)
	}

	if err := rows.Err(); err != nil {
		return assets, err
	}

	return assets, nil
}

// creating a new session
func Session(c *fiber.Ctx, dbConn *sql.DB) error {
	//retrieves the user
	tokenUser := c.Locals("user").(*jwt.Token)
	claims := tokenUser.Claims.(jwt.MapClaims)
	userID, ok := claims["id"].(string)

	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}
	//retrieves the user
	user := &db.User{}
	if err := dbConn.QueryRow(db.GetUserByIDQuery, userID).
		Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"success": false, "errors": []string{"Incorrect credentials"}})
		}
	}
	user.Password = ""

	return c.JSON(&fiber.Map{"success": true, "user": user})
}

func Login(c *fiber.Ctx, dbConn *sql.DB) error {
	loginUser := &db.User{}

	if err := c.BodyParser(loginUser); err != nil {
		return err
	}

	user := &db.User{}
	if err := dbConn.QueryRow(db.GetUserByEmailQuery, loginUser.Email).
		Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"success": false, "errors": []string{"Incorrect credentials"}})
		}
	}

	match := utils.ComparePassword(user.Password, loginUser.Password)
	if !match {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"success": false, "errors": []string{"Incorrect credentials"}})
	}

	//expiration time of the token ->30 mins
	//TODO: change expiration time?
	expirationTime := time.Now().Add(30 * time.Minute)

	user.Password = ""
	claims := &Claims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(config.Config[config.JWT_KEY])
	tokenValue, err := token.SignedString(jwtKey)

	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenValue,
		Expires:  expirationTime,
		Domain:   config.Config[config.CLIENT_URL],
		HTTPOnly: true,
	})

	return c.JSON(&fiber.Map{"success": true, "user": claims.User, "token": tokenValue})
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.SendStatus(http.StatusOK)
}

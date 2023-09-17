package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/KleinSpeedy/language-helper-backend/datatypes"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database controller struct holding
// database name
// database username and passwor
// port for database
// gets loaded on program startup from .env file
type Controller struct {
	name     string
	username string
	password string
	port     int
}

var db *gorm.DB

// Create new database controller
func NewController(name, username, pw string, port int) *Controller {
	return &Controller{
		name:     name,
		username: username,
		password: pw,
		port:     port,
	}
}

// Open gorm database connection for accessing information
// Uses environment variables from controller creation
// returns nil, true on success, err, false otherwise
func (c *Controller) OpenConnection() (error, bool) {
	// go sql database
	var sqldb *sql.DB
	// error variable
	var err error

	// see https://github.com/go-sql-driver/mysql#dsn-data-source-name for more information
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s", c.username, c.password, c.port, c.name)

	// open generic go sql connection with formatted dsn
	sqldb, err = sql.Open("mysql", dsn)
	if err != nil {
		return err, false
	}

	// open gorm connection to sql database
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqldb,
	}), &gorm.Config{})

	if err != nil {
		return err, false
	} else {
		return nil, true
	}
}

// Saves a user with password
// converts password to hash with custom salt
func (c *Controller) SaveNewUser(name, pw string) (error, bool) {
	// hash password for safety
	hashedPw, err := hashPassword(pw)
	if err != nil {
		return err, false
	}
	// store user with hashed password
	user := datatypes.User{
		Username: name,
		Password: string(hashedPw),
		Created:  time.Now().Unix(),
	}

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error, false
	}

	return nil, true
}

// check if user exists
func (c *Controller) UserExists(name string) (error, bool) {
	var user datatypes.User

	result := db.Where("Username = ?", name).First(&user)
	if result.Error != nil {
		// nil if user not yet exists
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false
		}
		// some other error
		return result.Error, false
	}
	// user exists
	return nil, true
}

// Retrieve user id from username
// returns id and nil on success, 0, false otherwise
func (c *Controller) getIdFromUsername(name string) (uint, error) {
	var user datatypes.User
	// Query database for username
	result := db.Where("Username = ?", name).First(&user)
	if result.Error != nil {
		fmt.Println("Error")
		return 0, result.Error
	}

	return user.UserId, nil
}

// get hashed password from username
// returns password and true on success, nil, false otherwise
func (c *Controller) getHashedPassword(name string) ([]byte, error) {
	// get id of user
	id, err := c.getIdFromUsername(name)
	if err != nil || id == 0 {
		return nil, err
	}
	var user datatypes.User
	// retrieve password from user by primary key
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return []byte(user.Password), nil
}

// Check if password for user exists and matches the saved hash
func (c *Controller) PasswordMatches(user, pw string) (error, bool) {
	hashedPw, err := c.getHashedPassword(user)
	if err != nil {
		return err, false
	}

	// err is nil if password matches
	err = bcrypt.CompareHashAndPassword(hashedPw, []byte(pw))
	return nil, (err == nil)
}

// hash the password string and return the hash or error on failure
func hashPassword(pw string) ([]byte, error) {
	// bcrypt function already adds a custom salt to the hash
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

// fetch user by username from database
// returns pointer to user or error on failure
func (c *Controller) GetUserAfterLogin(username string) (*datatypes.User, error) {
	var user datatypes.User
	// Query database for username
	result := db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		fmt.Println("Error")
		return nil, result.Error
	}

	return &user, nil
}

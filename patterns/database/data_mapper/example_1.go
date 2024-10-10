package data_mapper

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Сам шаблон

// User - Domain model
type User struct {
	ID    int
	Name  string
	Email string
}

// UserMapper - Data Mapper for User
type UserMapper struct {
	db *sql.DB
}

// NewUserMapper - Constructor for UserMapper
func NewUserMapper(db *sql.DB) *UserMapper {
	return &UserMapper{db: db}
}

// FindByID - Retrieves a user by ID
func (mapper *UserMapper) FindByID(id int) (*User, error) {
	row := mapper.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Save - Persists a user entity to the database
func (mapper *UserMapper) Save(user *User) error {
	_, err := mapper.db.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
	return err
}

// Клиентский код
func main() {
	connStr := "user=youruser dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userMapper := NewUserMapper(db)

	// Save a new user
	user := &User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	err = userMapper.Save(user)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the user
	retrievedUser, err := userMapper.FindByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %+v\n", retrievedUser)
}

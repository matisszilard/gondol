package store

import "github.com/matisszilard/gondol/model"

// UserStore users related operations
type UserStore interface {
	Get(string) (*model.User, error)
	GetAll() (*[]model.User, error)
	Create(*model.User) (string, error)
	Save(*model.User) error
}

// GetUser gets a user by the given id from the database
func GetUser(id string) (*model.User, error) {
	return instance.Users.Get(id)
}

// GetUsers returns all users from the database
func GetUsers() (*[]model.User, error) {
	return instance.Users.GetAll()
}

// CreateUser creates and saves a new user to the database
// returns the newly created users db id
func CreateUser(user *model.User) (string, error) {
	return instance.Users.Create(user)
}

// SaveUser saves the user to the database
func SaveUser(user *model.User) error {
	return instance.Users.Save(user)
}

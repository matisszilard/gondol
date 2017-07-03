package rethinkstore

import (
	"github.com/matisszilard/gondol/model"
	r "gopkg.in/gorethink/gorethink.v3"
)

// Users struct containing a rethink db session
type Users struct {
	session *r.Session
}

// Get Return a User structure
func (u *Users) Get(name string) (*model.User, error) {
	res, err := r.Table(UsersTableName).Filter(r.Row.Field("name").Eq(name)).Run(u.session)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var user model.User

	err = res.One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAll Return a User structure
func (u *Users) GetAll() (*[]model.User, error) {
	var users []model.User
	err := getAll(u.session, UsersTableName, &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// Create a user in the database
func (u *Users) Create(user *model.User) (string, error) {
	res, err := r.Table(UsersTableName).Insert(user).RunWrite(u.session)
	if err != nil {
		return "", err
	}
	return res.GeneratedKeys[0], nil
}

// Save a user in the database
func (u *Users) Save(user *model.User) error {
	_, err := r.Table(UsersTableName).Get(user.ID).Update(user).RunWrite(u.session)
	return err
}

package model

// User a model for storing a user's information
type User struct {
	ID    string `gorethink:"id,omitempty" json:"id"`
	Name  string `gorethink:"name"         json:"name"  form:"name" binding:"required"`
	Email string `gorethink:"email"        json:"email"  form:"email" binding:"required"`
}

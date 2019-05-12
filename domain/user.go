package domain

//User strutct
type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var UserModel []User

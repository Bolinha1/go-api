package domain

//User strutct
type User struct {
	id    int    `json:"id,omitempty"`
	name  string `json:"firstname,omitempty"`
	email string `json:"lastname,omitempty"`
	//Address *Address `json:"address,omitempty"`
}

//Andress struct
/*type Address struct {
	city  string `json:"city,omitempty"`
	state string `json:"state,omitempty"`
}*/

var users []User

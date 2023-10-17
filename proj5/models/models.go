package models

type User struct {
	FName    string `json:"name"` // giving the name of the field in the json ouptut
	Password string `json:"-"`    //ignore the field while creating the json
	LName    string `json:"l_name"`
	Email    string `json:"email"`
}

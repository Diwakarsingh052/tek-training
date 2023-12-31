// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewVideo struct {
	VideoID  string `json:"videoId"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

type UpdateVideoInput struct {
	Title *string `json:"title,omitempty"`
	URL   *string `json:"url,omitempty"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Video struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author *User  `json:"author"`
}

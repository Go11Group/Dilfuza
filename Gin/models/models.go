package models

type User struct {
    Id        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

type Problem struct {
    Id          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Difficulty  string `json:"difficulty"`
}

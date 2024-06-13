package models

type  Users struct {
    UserId   string `json:"id" db:"id"`
    Name      string `json:"name" db:"name"`
    Email     string `json:"email" db:"email"`
    Password  string `json:"password" db:"password"`
    CreatedAt  string `json:"created_at" db:"created_at"`
    UpdatedAt  string `json:"udated_at" db:"updated_at"`
    DeletedAt  string   `json:"deleted_at" db:"deleted_at"`
    Birthday   *string  `json:"birthday" db:"birthday"`
}


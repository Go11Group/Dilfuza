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



type SolvedProblem struct {
    UserId    string `json:"user_id"`
    ProblemId string `json:"problem_id"`
    SolvedAt  string `json:"solved_at"`
}


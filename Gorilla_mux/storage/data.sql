DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

CREATE TABLE Problem (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    titil  TEXT NOT NULL,
    Description TEXT NOT NULL,
    Difficulty   TEXT NOT NULL

);

CREATE TABLE SolvedProblem(
    UserId    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ProblemId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    SolvedAt UUID PRIMARY KEY DEFAULT gen_random_uuid()
);

CREATE TABLE  gtype as enum('m', 'f');

CREATE TABLE users(
    user_id   uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name   VARCHAR(100) NOT NULL,
    email  VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW (),
    deleted_at   BIGINT DEFAULT 0,
    birthday    TIMESTAMP 
);

CREATE TABLE courses (
    course_id  uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title  VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE lessons (
    lesson_id  uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id  VARCHAR(100) NOT NULL,
    title      VARCHAR(100) NOT NULL,
    content    TEXT  NOT NULL,
    created_at   TIMESTAMP DEFAULT NOW (),
    updated_at   TIMESTAMP DEFAULT NOW (),
    deleted_at   BIGINT  DEFAULT 0
);


CREATE TABLE enrollments (
    enrollment_id uuid  PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      VARCHAR(100) NOT NULL,
    course_id     VARCHAR(100) NOT NULL,
    enrollment_date  TIMESTAMP DEFAULT NOW () ,
    created_at    TIMESTAMP DEFAULT NOW (),
    updated_at    TIMESTAMP DEFAULT NOW (),
    deleted_at    BIGINT DEFAULT  0
);


--update table_name set
 --deleted_at = date_part('epoch', current_timestamp)::INT
--where id = $1 and deleted_at = 0


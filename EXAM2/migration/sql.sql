CREATE TABLE users (
    user_id uuid  PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    email VARCHAR(100),
    birthday TIMESTAMP,
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  BIGINT DEFAULT 0
);

CREATE TABLE courses (
    course_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title  VARCHAR (100),
    description VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  BIGINT DEFAULT 0
);

CREATE TABLE lessons (
    lesson_id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULl,
    course_id UUID REFERENCES courses(course_id),
    title VARCHAR(100),
    content TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  BIGINT DEFAULT 0

);

CREATE TABLE enrollments (
    enrollment_id UUID PRIMARY KEY  DEFAULT gen_random_uuid() NOT NULl,
    user_id UUID REFERENCES users(user_id),
    course_id UUID REFERENCES courses(course_id),
    enrollment_date TIMESTAMP ,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    deleted_at  BIGINT DEFAULT 0
);
create  table  if not exists  cource(
   id uuid primary key ,cource_name varchar,
   rector varchar,cource_email varchar,
   course_phone varchar not null  check ( '+9989' ),
   created_at timestamp,
   updated_at timestamp ,
   deleted_at timestamp
);                                       

INSERT INTO cource (id, cource_name, rector, cource_email, course_phone, created_at, updated_at, deleted_at) VALUES
    ('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Computer Science', 'John Doe', 'cs@example.com', '+998901234567', '2024-06-01 10:00:00', '2024-06-01 10:00:00', NULL),
    ('f47ac10b-58cc-4372-a567-0e02b2c3d480', 'Mathematics', 'Jane Smith', 'math@example.com', '+998912345678', '2024-06-02 11:00:00', '2024-06-02 11:00:00', NULL),
    ('f47ac10b-58cc-4372-a567-0e02b2c3d481', 'Physics', 'Albert Einstein', 'physics@example.com', '+998923456789', '2024-06-03 12:00:00', '2024-06-03 12:00:00', NULL),
    ('f47ac10b-58cc-4372-a567-0e02b2c3d482', 'Chemistry', 'Marie Curie', 'chemistry@example.com', '+998934567890', '2024-06-04 13:00:00', '2024-06-04 13:00:00', NULL),
    ('f47ac10b-58cc-4372-a567-0e02b2c3d483', 'Biology', 'Charles Darwin', 'biology@example.com', '+998945678901', '2024-06-05 14:00:00', '2024-06-05 14:00:00', NULL);

create  table if not exists students (
    id uuid   primary key  default gen_random_uuid(),
    first_name varchar,
    last_name varchar,
    age int ,
    email  varchar not null,
    phone_number varchar not null,
    created_at timestamp,
    updated_at timestamp ,
    deleted_at timestamp

);

INSERT INTO students (first_name, last_name, age, email, phone_number, created_at, updated_at, deleted_at) VALUES
    ('Alice', 'Smith', 20, 'alice.smith@example.com', '+998901234567', '2024-06-01 10:00:00', '2024-06-01 10:00:00', NULL),
    ('Bob', 'Johnson', 22, 'bob.johnson@example.com', '+998912345678', '2024-06-02 11:00:00', '2024-06-02 11:00:00', NULL),
    ('Charlie', 'Brown', 21, 'charlie.brown@example.com', '+998923456789', '2024-06-03 12:00:00', '2024-06-03 12:00:00', NULL),
    ('David', 'Williams', 23, 'david.williams@example.com', '+998934567890', '2024-06-04 13:00:00', '2024-06-04 13:00:00', NULL),
    ('Eve', 'Davis', 20, 'eve.davis@example.com', '+998945678901', '2024-06-05 14:00:00', '2024-06-05 14:00:00', NULL);


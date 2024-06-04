-- Users jadvali
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

INSERT INTO users (username, email, password) VALUES
('Dilfuza', 'dilfuza@example.com', '1234'),
('Akmal', 'akmar@example.com', '8765'),
('Imron', 'imron@example.com', '4326'),
('Lazin', 'lazina@example.com', '9084'),
('Samina', 'samina@example.com', '1007'),
('Merojidin', 'merajidin@example.com', '3510'),
('Humayro', 'humayro@example.com', '4612'),
('Ibroxim', 'ibroxim@example.com', '8610'),
('Mustafo', 'mustafo@example.com', '7608'),
('Odil', 'odil@example.com', '1493');



select* from users;


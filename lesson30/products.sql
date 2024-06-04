-- Products jadvali
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

INSERT INTO products (name, description, price, stock_quantity) VALUES
('Non',       'Description ', 2800.99, 10),
('Sut',       'Description ', 12900.90, 10),
('Qaymoq',    'Description ', 15000.10, 20),
('Shokolad',  'Description ', 65000.15, 25),
('Guruch',    'Description ', 17000.99, 12),
('Yog',       'Description ', 14000.90, 12),
('Kofie',     'Description ', 26000.45, 15),
('Pomidor',   'Description ', 5000.00, 15),
('Sharbat',   'Description ', 13000.00, 25),
('Meneral suv','Description ',3000.00, 15);


select  *from products;

 
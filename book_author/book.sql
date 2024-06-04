


CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birthdate DATE
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES author(id)
);




INSERT INTO author(name,birthdate) VALUES
('Otkir Hoshimov','1950-01-05'),
('Togay Murod','1965-04-02'),
('Abdulfattoh Abu Gudda','1850-09-03'),
('Nathaniel hawthorne','1975-09-09'),
('Ahmad Hodiy Maqsudiy','1980-08-07'),
('Doktor Oiz Qarniy','1800-09-07'),
('Javlon Abdullo','1968-01-01');



-- tablega ma'lumot yozish
INSERT INTO book (title, genre, author_id) VALUES
('Ikki eshik orasi','Romon', 1),
('Otkan kunlar','Roman', 1),
('Oydinda yurgan odamlar','Qissa' , 2),
('Ot kishnagan oqshom','Asar', 2),
('Ulomolar Nazdida Vaqtning qadri','Asar',3),
('The Scarlet Letter','Qissa',4),
('Ibodati Islomiya','Asar', 4),
('Ukinma','Asar', 6),
('Muallimi Soniy','Tuplam',4),
('Mukammal Dasturlash','Darslik',5),
('Dunyoning Ishlari','Roman',1),
('Uliss Roman','Roman', 6),
('Alvido bolalik','Roman', 6),
('Alvido Vatan','Roman', 7),
('Yashamoq','Asar', 7);

*****************************
SELECT a.name, COUNT(b.title) AS "Yozgan kitoblar soni"
FROM
    author AS a
INNER JOIN book AS b
    ON a.id = b.author_id
GROUP BY
    a.name


SELECT a.name, ARRAY_AGG(DISTINCT b.genre) AS "IJON Qilgan Janrlar"
FROM 
    author AS a
INNER JOIN book AS b
    ON a.id = b.author_id
GROUP BY a.name

SELECT a.name, b.genre, COUNT(b.genre) AS "Har bir janrda yozgan kitoblar soni"
FROM 
    author AS a
INNER JOIN book AS b
    ON a.id = b.author_id
GROUP BY a.name, b.genre;

SELECT a.name, b.title AS "Kitob nomi"
FROM author AS a
LEFT JOIN book AS b
    ON a.id = b.author_id;

SELECT a.name, b.title AS "Eng yoshi katta yozuvchining asarlari"
FROM 
    book AS b
INNER JOIN author AS a
    ON b.author_id = a.id
WHERE 
    a.birthdate = (SELECT min(birthdate) FROM author);
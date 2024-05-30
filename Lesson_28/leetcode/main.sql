---Leet sql 1045


create table Customer(customer_id int, product_key int);

create table Product(product_key int);

insert into Customer(customer_id ,product_key) values(1 , 5),(2, 6),(3,5),(3,6),(1,6);

insert into Product(product_key)values(5),(6);

SELECT customer_id
FROM Customer
GROUP BY customer_id
HAVING COUNT(DISTINCT product_key) = (SELECT COUNT(*) FROM Product);
 ***************************************************


-- leet sql 1729


create table Followers(user_id int, follower_id int);
insert into Followers(user_id,follower_id) values(0, 1),(1,0),(2,0),(2,1);

 SELECT user_id, COUNT(follower_id) AS followers_count
FROM Followers
GROUP BY user_id
ORDER BY user_id ASC;
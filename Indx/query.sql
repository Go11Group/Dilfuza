create table Users
(
    id              uuid DEFAULT gen_random_uuid(),
    first_name      varchar,
    last_name       varchar,
    email           varchar,
    password        varchar,
    about           text,
    followers       int,
    address         varchar,
    is_employed     boolean,
    phone_number    varchar
);

explain (analyse )
select * from Users where first_name  = 'Luis' and last_name='Boyle' and followers = 222780 and is_employed = true;

create index product_id_idx on Users (id);
create index product_name_indx on Users (first_name);
create index product_is_employed_indx on Users (is_employed);

create index product_name_indx on Users using hash(first_name);

drop index product_id_idx;
drop index product_name_indx;
drop index product_is_employed_indx;

explain (analyse)
select * 
from Users 
where id = 'b2457480-75c7-481b-8b76-c40795ec8ff0';

and 
first_name  = 'Luis' and last_name='Boyle' and
followers = 222780 and is_employed = true;

explain (analyse)
select * 
from Users 
where first_name  = 'Luis' and last_name='Boyle' and
followers = 222780 and is_employed = true;

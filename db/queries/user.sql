-- name: GetUser :one
select * from users where id = $1 limit 1;

-- name: GetUserByLogin :one
select * from users where login = $1 limit 1;

-- name: GetUserByEmail :one
select * from users where email = $1 limit 1;

-- name: CreateUser :one
insert into users (id, email, login, password_hash, avatar) 
values ($1, $2,  $3, $4, $5) returning *;

-- update users set deleted_at = CURRENT_TIMESTAMP where id = (select id from users limit 1); 
-- update users set updated_at = CURRENT_TIMESTAMP where id = (select id from users limit 1);
-- update users set password_hash = 'new_password' where id = (select id from users limit 1);
-- update users set email = 'email@test.com' where id = (select id from users limit 1);
-- update users set first_name = 'first_name' where id = (select id from users limit 1);
-- update users set last_name = 'last_name' where id = (select id from users limit 1);
-- update users set middle_name = 'middle_name' where id = (select id from users limit 1);
-- update users set avatar = 'new_avatar' where id = (select id from users limit 1);

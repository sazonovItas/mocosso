-- name: GetUserAccount :one
select * from user_account
where id = $1 limit 1;

-- name: GetUserAccountByUsername :one
select * from user_account
where username = $1 limit 1;

-- name: ListUserAccount :many
select * from user_account;

-- name: ListUserRoles :many
select r.* from role as r
join (
  select role_id from user_role
  where user_id = $1
) as ur on r.id = ur.id;

-- name: CreateUserAccount :one
insert into user_account (
  email, username, password_hash
) values (
  $1, $2, $3
) returning *;

-- name: UpdateUserAccount :exec
update user_account
  set email = $1,
  username = $2
where id = $1;

-- name: UpdateUserAccountVerifiedStatus :exec
update user_account
  set is_verified = $1
where email = $2;

-- name: UpdateUserAccountAvatar :exec
update user_account
  set avatar = $1
where id = $2;

-- name: DeleteUserAccount :exec
update user_account
  set deleted_at = $1
where id = $2;

-- name: GetUserAccount :one
select * from user_account
where id = $1 limit 1;

-- name: GetUserAccountByUsername :one
select * from user_account
where username = $1 limit 1;

-- name: ListUserAccount :many
select * from user_account;

-- name: ListUserAccountRoles :many
select r.* from role as r
join (
  select role_id from user_role
  where user_id = $1
) as ur on r.id = ur.role_id;

-- name: ListUserAccountScopes :many
select distinct sc.* 
from scope as sc
join (
  select distinct rsc.scope_id 
  from role_scope as rsc
  join (
    select id as role_id from role as r
    join (
      select role_id from user_role
      where user_id = $1
    ) as ur on r.id = ur.role_id
  ) as r on rsc.role_id = r.role_id
) as usc on sc.id = rsc.scope_id;

-- name: CreateUserAccount :one
insert into user_account (
  email, username, password_hash
) values (
  $1, $2, $3
) returning *;

-- name: AddRoleToUserAccount :exec
insert into user_role (
  user_id, role_id
) values (
  $1, $2
);

-- name: UpdateUserAccount :exec
update user_account
  set email = $1,
  username = $2
where id = $3;

-- name: UpdateUserAccountPassword :exec
update user_account
  set password_hash = $1
where id = $2;

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

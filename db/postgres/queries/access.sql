-- name: GetUserAccess :one
select * from access
where id = $1 limit 1;

-- name: GetUserDeviceAccess :one
select * from access
where device_id = $1 limit 1;

-- name: ListUserAccess :many
select * from access
where user_id = $1;

-- name: CreateUserAccess :one
insert into access (
  user_id, device_id, refresh_token, expires_at
) values (
  $1, $2, $3, $4
) returning *;

-- name: DeleteUserAccess :exec
delete from access where id = $1;

-- name: DeleteUserAccessByUserIDAndDeviceID :exec
delete from access
where user_id = $1 and device_id = $2;

-- name: GetUserDevice :one
select * from device
where id = $1 limit 1;

-- name: GetUserDeviceByHashID :one
select * from device
where hash_id = $1 limit 1;

-- name: ListUserDevices :many
select * from device
where user_id = $1;

-- name: CreateUserDevice :one
insert into device (
  name, user_id, hash_id 
) values (
  $1, $2, $3
) returning *;

-- name: UpdateUserDevice :exec
update device
  set name = $1
where id = $2;

-- name: DeleteUserDevice :exec
delete from device 
where id = $1;

-- name: DeleteUserDeviceByHashID :exec
delete from device
where hash_id = $1;

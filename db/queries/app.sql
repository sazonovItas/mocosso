-- name: GetApp :one
select * from apps where id = $1 limit 1;

-- name: CreateApp :one
insert into apps (name, secret, scope) 
values ($1, $2, $3) returning *;

-- name: DeleteApp :exec
delete from apps where id = $1;

-- name: UpdateAppName :exec
update apps set name = $2 where id = $1;

-- name: UpdateAppScope :exec
update apps set scope = $2 where id = $1;

-- name: UpdateAppSecret :exec
update apps set secret = $2 where id = $1;

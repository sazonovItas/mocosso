-- name: GetRoleByName :one
select * from role
where name = $1 limit 1;

-- name: CreateRole :one
insert into role (
  name
) values (
  $1
) returning *;

-- name: UpdateRoleByName :exec
update role
  set name = $1
where name = $2;

-- name: DeleteRoleByName :exec
delete from role
where name = $1;

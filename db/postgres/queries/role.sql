-- name: GetRoleByName :one
select * from role
where name = sqlc.arg(name) limit 1;

-- name: ListRoleScopes :many
select sc.* from scope as sc
join (
  select scope_id from role_scope
  where role_id = sqlc.arg(role_id)
) as rsc on sc.id = rsc.scope_id;

-- name: CreateRole :one
insert into role (
  name
) values (
  sqlc.arg(name)
) returning *;

-- name: AddScopeToRole :exec
insert into role_scope (
  role_id, scope_id 
) values (
  sqlc.arg(role_id), sqlc.arg(scope_id)
);

-- name: UpdateRoleByName :exec
update role
  set name = sqlc.arg(name)
where name = sqlc.arg(name);

-- name: DeleteRole :exec
delete from role
where id = sqlc.arg(role_id);

-- name: DeleteRoleByName :exec
delete from role
where name = sqlc.arg(name);

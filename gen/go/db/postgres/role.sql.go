// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: role.sql

package postgresdb

import (
	"context"
)

const addScopeToRole = `-- name: AddScopeToRole :exec
insert into role_scope (
  role_id, scope_id 
) values (
  $1, $2
)
`

type AddScopeToRoleParams struct {
	RoleID  int32
	ScopeID int32
}

func (q *Queries) AddScopeToRole(ctx context.Context, arg AddScopeToRoleParams) error {
	_, err := q.db.Exec(ctx, addScopeToRole, arg.RoleID, arg.ScopeID)
	return err
}

const createRole = `-- name: CreateRole :one
insert into role (
  name
) values (
  $1
) returning id, name, description
`

func (q *Queries) CreateRole(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deleteRole = `-- name: DeleteRole :exec
delete from role
where id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, roleID int32) error {
	_, err := q.db.Exec(ctx, deleteRole, roleID)
	return err
}

const deleteRoleByName = `-- name: DeleteRoleByName :exec
delete from role
where name = $1
`

func (q *Queries) DeleteRoleByName(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, deleteRoleByName, name)
	return err
}

const getRoleByName = `-- name: GetRoleByName :one
select id, name, description from role
where name = $1 limit 1
`

func (q *Queries) GetRoleByName(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listRoleScopes = `-- name: ListRoleScopes :many
select sc.id, sc.name, sc.description from scope as sc
join (
  select scope_id from role_scope
  where role_id = $1
) as rsc on sc.id = rsc.scope_id
`

func (q *Queries) ListRoleScopes(ctx context.Context, roleID int32) ([]Scope, error) {
	rows, err := q.db.Query(ctx, listRoleScopes, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scope
	for rows.Next() {
		var i Scope
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRoleByName = `-- name: UpdateRoleByName :exec
update role
  set name = $1
where name = $1
`

func (q *Queries) UpdateRoleByName(ctx context.Context, name string) error {
	_, err := q.db.Exec(ctx, updateRoleByName, name)
	return err
}

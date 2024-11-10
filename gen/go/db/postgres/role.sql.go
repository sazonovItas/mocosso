// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: role.sql

package postgresdb

import (
	"context"
)

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

const updateRoleByName = `-- name: UpdateRoleByName :exec
update role
  set name = $1
where name = $2
`

type UpdateRoleByNameParams struct {
	Name   string
	Name_2 string
}

func (q *Queries) UpdateRoleByName(ctx context.Context, arg UpdateRoleByNameParams) error {
	_, err := q.db.Exec(ctx, updateRoleByName, arg.Name, arg.Name_2)
	return err
}

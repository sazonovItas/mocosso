-- name: GetVerificationByEmail :one
select * from verification
where email = $1 limit 1;

-- name: GetVerificationByToken :one
select * from verification
where token = $1 limit 1;

-- name: CreateVerification :one
insert into verification (
  email, type, code, token, expires_at
) values (
  $1, $2, $3, $4, $5
) returning *;

-- name: UpdateVerification :exec
update verification 
  set type = $1,
  code = $2,
  token = $3,
  expires_at = $4
where email = $5;

-- name: DeleteVerification :exec
delete from verification
where email = $1;

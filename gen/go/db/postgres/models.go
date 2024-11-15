// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgresdb

import (
	"time"
)

type Access struct {
	ID           int64
	UserID       int64
	DeviceID     int64
	RefreshToken string
	CreatedAt    time.Time
	LastUsedAt   time.Time
	ExpiresAt    time.Time
}

type Device struct {
	ID         int64
	Name       string
	UserID     int64
	HashID     string
	LastUsedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Role struct {
	ID          int32
	Name        string
	Description string
}

type RoleScope struct {
	RoleID  int32
	ScopeID int32
}

type Scope struct {
	ID          int32
	Name        string
	Description string
}

type UserAccount struct {
	ID           int64
	Email        string
	Username     string
	PasswordHash string
	Avatar       string
	IsVerified   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type UserRole struct {
	UserID int64
	RoleID int32
}

type Verification struct {
	Email     string
	Type      string
	Code      string
	Token     string
	ExpiresAt time.Time
	UpdatedAt time.Time
}

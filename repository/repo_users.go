package repository

import (
	"Latihan-JWT-User/model/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Update(ctx context.Context, tx *sql.Tx, user entity.Users) entity.Users
	Delete(ctx context.Context, tx *sql.Tx, user entity.Users)
	FindByID(ctx context.Context, tx *sql.Tx, userId string) (entity.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Users
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Users, error)
}

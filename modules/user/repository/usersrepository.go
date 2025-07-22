package repository

import (
	"aplikasi-adakost-be/modules/user/model"
	"database/sql"
)

type UsersRepository interface {
	Register(user model.Users) (model.Users, error)
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (u *usersRepository) Register(user model.Users) (model.Users, error) {

	sql := "INSERT INTO adk_users(username, password, role_id, created_at, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	errs := u.db.QueryRow(sql, &user.Username, &user.Password, &user.RoleId, &user.CreatedAt, &user.CreatedBy).Scan(&user.Id)
	if errs != nil {
		panic(errs)
	}

	return user, nil
}

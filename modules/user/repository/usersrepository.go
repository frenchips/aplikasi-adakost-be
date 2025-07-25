package repository

import (
	"aplikasi-adakost-be/modules/user/model"
	"database/sql"
)

type UsersRepository interface {
	Register(user model.Users) (model.Users, error)
	Login(user model.UserLogin) (model.UserLogin, error)
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

func (u *usersRepository) Login(user model.UserLogin) (model.UserLogin, error) {

	var rolesName string
	sql := `SELECT au.id, au.username, au.password, ar.name FROM adk_users au JOIN adk_roles ar ON au.role_id = ar.id WHERE au.username = $1 and au.password = $2`
	err := u.db.QueryRow(sql, user.Username, user.Password).Scan(&user.Id, &user.Username, &user.Password, &rolesName)
	if err != nil {
		return model.UserLogin{}, err
	}
	user.Roles = []model.Roles{{RoleName: rolesName}}

	return user, nil

}

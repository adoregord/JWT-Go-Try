package repository

import (
	"database/sql"
	"jwt-try/internal/domain"
	utils "jwt-try/internal/utils/hash"
)

type UserRepoInterface interface {
	CheckCredential(user *domain.User) bool
	RegisterUser(user *domain.User) error
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoInterface {
	return UserRepo{
		db: db,
	}
}

func (repo UserRepo) CheckCredential(user *domain.User) bool {
	query := `
	select password
	from account a
	where username = $1`

	var pass string

	err := repo.db.QueryRow(query, user.Username).Scan(&pass)
	if err != nil {
		return false
	}

	ok := utils.CheckPasswordHash(user.Password, pass)

	return ok
}

func (repo UserRepo) RegisterUser(User *domain.User) error {
	query := `
  insert into account (username, password)
	values ($1, $2)
	`
	// encrypt the password
	repo.db.QueryRow(query, User.Username, User.Password)

	return nil
}

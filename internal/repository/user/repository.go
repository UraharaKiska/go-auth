package auth

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	// "github.com/UraharaKiska/go-auth/internal/client/db"
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-auth/internal/model"
	"github.com/UraharaKiska/go-auth/internal/repository"
	"github.com/UraharaKiska/go-auth/internal/repository/user/converter"
	modelRepo "github.com/UraharaKiska/go-auth/internal/repository/user/model"
)

const (
	tableName             = "auth"
	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	roleColumn            = "role"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	log.Printf("REPOSITORY - CREATE")
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn).
		Values(info.Name, info.Email, info.Password, info.PasswordConfirm, info.Role).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	log.Printf("REPOSITORY - GET")
	builderSelect := sq.Select(idColumn, nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	log.Printf("User: %v", user)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, info *modelRepo.UpdateUserInfo, id int64) error {
	log.Printf("REPOSITORY - UPDATE")
	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar)
	if info.Name.Valid {
		builderUpdate = builderUpdate.Set(nameColumn, info.Name.String)
	}
	if info.Email.Valid {
		builderUpdate = builderUpdate.Set(emailColumn, info.Email.String)
	}
	builderUpdate = builderUpdate.Where(sq.Eq{idColumn: id})
	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	log.Printf("REPOSITORY - DELETE")
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})
	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

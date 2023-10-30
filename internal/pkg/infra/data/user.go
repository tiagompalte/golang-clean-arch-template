package data

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type UserRepository struct {
	conn         pkgRepo.Connector
	mainTable    string
	selectFields string
}

func NewUserRepository(conn pkgRepo.Connector) repository.UserRepository {
	return UserRepository{
		conn:      conn,
		mainTable: "tb_user",
		selectFields: `
			u.id
			, u.created_at
			, u.updated_at
			, u.uuid
			, u.name
			, u.email

			FROM tb_user u
		`,
	}
}

func (r UserRepository) parseEntity(s pkgRepo.Scanner) (entity.User, error) {
	var user entity.User
	err := s.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.UUID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return entity.User{}, errors.Repo(err, r.mainTable)
	}

	return user, nil
}

func (r UserRepository) Insert(ctx context.Context, user entity.User, passEncrypted string) (uint32, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	res, err := r.conn.ExecContext(ctx,
		"INSERT INTO tb_user (uuid, name, email, pass_encrypted) VALUES (?,?,?,?)",
		uuid,
		user.Name,
		user.Email,
		passEncrypted,
	)
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	return uint32(id), nil
}

func (r UserRepository) FindByID(ctx context.Context, id uint32) (entity.User, error) {
	query := `
		SELECT %s
			WHERE NOT u.deleted_at AND u.id = ?`

	q := fmt.Sprintf(query, r.selectFields)

	user, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			id,
		))
	if err != nil {
		return entity.User{}, errors.Repo(err, r.mainTable)
	}

	return user, nil
}

func (r UserRepository) FindByUUID(ctx context.Context, uuid string) (entity.User, error) {
	query := `
		SELECT %s
			WHERE NOT u.deleted_at AND u.uuid = ?`

	q := fmt.Sprintf(query, r.selectFields)

	user, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			uuid,
		))
	if err != nil {
		return entity.User{}, errors.Repo(err, r.mainTable)
	}

	return user, nil
}

func (r UserRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `
		SELECT %s
			WHERE NOT u.deleted_at AND u.email = ?`

	q := fmt.Sprintf(query, r.selectFields)

	user, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			email,
		))
	if err != nil {
		return entity.User{}, errors.Repo(err, r.mainTable)
	}

	return user, nil
}

func (r UserRepository) GetPassEncryptedByEmail(ctx context.Context, email string) (string, error) {
	query := `
		SELECT pass_encrypted
			FROM tb_user			
			WHERE NOT deleted_at AND email = ?`

	var passEncrypted string
	err := r.conn.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&passEncrypted,
	)
	if err != nil {
		return "", errors.Repo(err, r.mainTable)
	}

	return passEncrypted, nil
}

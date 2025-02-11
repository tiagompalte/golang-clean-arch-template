package nativemigrate

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type VersionRepositoryImpl struct {
	conn repository.ConnectorSql
}

func NewVersionRepositoryImpl(conn repository.ConnectorSql) VersionRepository {
	return VersionRepositoryImpl{conn: conn}
}

func (r VersionRepositoryImpl) CreateTable(ctx context.Context) error {
	_, err := r.conn.Exec(
		ctx,
		`CREATE TABLE IF NOT EXISTS tb_db_version (
			name 				VARCHAR(250) NOT NULL PRIMARY KEY
			, created_at        DATETIME NOT NULL DEFAULT NOW()
			, applied_at		DATETIME NULL
		)`,
	)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (r VersionRepositoryImpl) InsertBatch(ctx context.Context, names []string) error {
	args := []any{}
	for _, n := range names {
		args = append(args, n)
	}

	params := strings.Repeat(",(?)", len(args))[1:]
	query := fmt.Sprintf(`INSERT IGNORE INTO tb_db_version (name) VALUES %s`, params)

	_, err := r.conn.Exec(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (r VersionRepositoryImpl) IsAlreadyApply(ctx context.Context, name string) (bool, error) {
	const query = `
		SELECT EXISTS(
			SELECT 1
			FROM tb_db_version
			WHERE name = ? AND applied_at IS NOT NULL
		);
	`

	var exists bool
	err := r.conn.QueryRow(
		ctx,
		query,
		name,
	).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err)
	}

	return exists, nil
}

func (r VersionRepositoryImpl) ExecScript(ctx context.Context, script string) error {
	_, err := r.conn.Exec(ctx, script)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (r VersionRepositoryImpl) FindLastInserted(ctx context.Context) (string, error) {
	const query = `
		SELECT name
		FROM tb_db_version
		WHERE applied_at IS NOT NULL
		ORDER BY applied_at DESC
		LIMIT 1
	;`

	var name string
	err := r.conn.QueryRow(
		ctx,
		query,
	).Scan(&name)
	if err != nil {
		return "", errors.Wrap(err)
	}

	return name, nil
}

func (r VersionRepositoryImpl) UpdateAppliedAt(ctx context.Context, name string, appliedAt *time.Time) error {
	_, err := r.conn.Exec(
		ctx,
		`UPDATE tb_db_version SET applied_at = ? WHERE name = ?`,
		appliedAt,
		name,
	)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

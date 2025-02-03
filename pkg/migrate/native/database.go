package nativemigrate

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func (m NativeMigrate) createTable(ctx context.Context) error {
	_, err := m.data.ExecContext(
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

func (m NativeMigrate) insertBatch(ctx context.Context, names []string) error {
	args := []any{}
	for _, n := range names {
		args = append(args, n)
	}

	params := strings.Repeat(",(?)", len(args))[1:]
	query := fmt.Sprintf(`INSERT IGNORE INTO tb_db_version (name) VALUES %s`, params)

	_, err := m.data.ExecContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) insertScripts(ctx context.Context, scripts map[string]Script) error {
	err := m.createTable(ctx)
	if err != nil {
		return errors.Wrap(err)
	}

	names := make([]string, 0, len(scripts))
	for name, script := range scripts {
		if !script.IsValid() {
			return errors.Wrap(ErrScriptInvalid)
		}

		names = append(names, name)
	}

	err = m.insertBatch(ctx, names)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) alreadyApply(ctx context.Context, name string) (bool, error) {
	const query = `
		SELECT EXISTS(
			SELECT 1
			FROM tb_db_version
			WHERE name = ? AND applied_at IS NOT NULL
		);
	`

	var exists bool
	err := m.data.QueryRowContext(
		ctx,
		query,
		name,
	).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err)
	}

	return exists, nil
}

func (m NativeMigrate) execUpScript(ctx context.Context, name string, script string) error {
	tx, err := m.data.Begin()
	if err != nil {
		return errors.Wrap(err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, script)
	if err != nil {
		return errors.Wrap(err)
	}

	_, err = tx.ExecContext(
		ctx,
		`UPDATE tb_db_version SET applied_at = ? WHERE name = ?`,
		time.Now(),
		name,
	)
	if err != nil {
		return errors.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) findLastInserted(ctx context.Context) (string, error) {
	const query = `
		SELECT name
		FROM tb_db_version
		WHERE applied_at IS NOT NULL
		ORDER BY applied_at DESC
		LIMIT 1
	;`

	var name string
	err := m.data.QueryRowContext(
		ctx,
		query,
	).Scan(&name)
	if err != nil {
		return "", errors.Wrap(err)
	}

	return name, nil
}

func (m NativeMigrate) execDownScript(ctx context.Context, name string, script string) error {
	tx, err := m.data.Begin()
	if err != nil {
		return errors.Wrap(err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, script)
	if err != nil {
		return errors.Wrap(err)
	}

	_, err = tx.ExecContext(
		ctx,
		`UPDATE tb_db_version SET applied_at = null WHERE name = ?`,
		name,
	)
	if err != nil {
		return errors.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

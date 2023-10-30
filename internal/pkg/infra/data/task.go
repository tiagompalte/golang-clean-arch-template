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

type TaskRepository struct {
	conn         pkgRepo.Connector
	mainTable    string
	selectFields string
}

func NewTaskRepository(conn pkgRepo.Connector) repository.TaskRepository {
	return TaskRepository{
		conn:      conn,
		mainTable: "tb_task",
		selectFields: `
			t.id
			, t.created_at
			, t.updated_at
			, t.uuid
			, t.name
			, t.description
			, t.done

			, c.id
			, c.created_at
			, c.updated_at
			, c.slug
			, c.name

			, t.user_id
			
			FROM tb_task t

			JOIN tb_category c ON NOT c.deleted_at AND t.category_id = c.id
		`,
	}
}

func (r TaskRepository) parseEntity(s pkgRepo.Scanner) (entity.Task, error) {
	var task entity.Task
	var categorySlug string
	err := s.Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.UUID,
		&task.Name,
		&task.Description,
		&task.Done,
		&task.Category.ID,
		&task.Category.CreatedAt,
		&task.Category.UpdatedAt,
		&categorySlug,
		&task.Category.Name,
		&task.UserID,
	)
	if err != nil {
		return entity.Task{}, errors.Repo(err, r.mainTable)
	}
	task.Category.SetSlug(categorySlug)

	return task, nil
}

func (r TaskRepository) Insert(ctx context.Context, task entity.Task) (uint32, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	res, err := r.conn.ExecContext(ctx,
		"INSERT INTO tb_task (uuid, name, description, category_id, user_id) VALUES (?,?,?,?,?)",
		uuid,
		task.Name,
		task.Description,
		task.Category.ID,
		task.UserID,
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

func (r TaskRepository) FindByID(ctx context.Context, id uint32) (entity.Task, error) {
	query := `
		SELECT %s
			WHERE NOT t.deleted_at AND t.id = ?`

	q := fmt.Sprintf(query, r.selectFields)

	task, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			id,
		))
	if err != nil {
		return entity.Task{}, errors.Repo(err, r.mainTable)
	}

	return task, nil
}

func (r TaskRepository) FindByUUID(ctx context.Context, uuid string) (entity.Task, error) {
	query := `
		SELECT %s			
			WHERE NOT t.deleted_at AND t.uuid = ?`

	q := fmt.Sprintf(query, r.selectFields)

	task, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			uuid,
		))
	if err != nil {
		return entity.Task{}, errors.Repo(err, r.mainTable)
	}

	return task, nil
}

func (r TaskRepository) FindByUserID(ctx context.Context, userID uint32) ([]entity.Task, error) {
	query := `
		SELECT %s
			WHERE NOT t.deleted_at AND t.user_id = ?`

	q := fmt.Sprintf(query, r.selectFields)

	result, err := r.conn.QueryContext(
		ctx,
		q,
		userID,
	)
	list, err := pkgRepo.ParseEntities[entity.Task](
		r.parseEntity,
		result,
		err,
	)
	if err != nil {
		return []entity.Task{}, errors.Repo(err, r.mainTable)
	}

	return list, nil
}

func (r TaskRepository) UpdateDone(ctx context.Context, task entity.Task) error {
	_, err := r.conn.ExecContext(ctx,
		`UPDATE tb_task
			SET done = ?
		WHERE NOT deleted_at AND id = ?
		`,
		task.Done,
		task.ID,
	)
	if err != nil {
		return errors.Repo(err, r.mainTable)
	}

	return nil
}

func (r TaskRepository) DeleteByID(ctx context.Context, taskID uint32) error {
	_, err := r.conn.ExecContext(ctx,
		`UPDATE tb_task	SET deleted_at = NOW() WHERE NOT deleted_at AND id = ?`,
		taskID,
	)
	if err != nil {
		return errors.Repo(err, r.mainTable)
	}

	return nil
}

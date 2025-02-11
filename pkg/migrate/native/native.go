package nativemigrate

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type NativeMigrate struct {
	file File
	repo RepositoryManager
}

func NewNativeMigrate(
	file File,
	repo RepositoryManager,
) NativeMigrate {
	return NativeMigrate{
		file: file,
		repo: repo,
	}
}

func (m NativeMigrate) Create(ctx context.Context, name string) error {
	version := time.Now().Unix()
	return m.file.CreateUpAndDownFile(version, name)
}

func (m NativeMigrate) insertNames(ctx context.Context, names []string) error {
	tx, err := m.repo.Data().Begin()
	if err != nil {
		return errors.Wrap(err)
	}
	defer tx.Rollback()

	versionRepo := m.repo.Version(tx.Command())

	err = versionRepo.CreateTable(ctx)
	if err != nil {
		return errors.Wrap(err)
	}

	err = versionRepo.InsertBatch(ctx, names)
	if err != nil {
		return errors.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) execUpScript(ctx context.Context, name string) error {
	tx, err := m.repo.Data().Begin()
	if err != nil {
		return errors.Wrap(err)
	}
	defer tx.Rollback()

	versionRepo := m.repo.Version(tx.Command())

	isAlreadyApply, err := versionRepo.IsAlreadyApply(ctx, name)
	if err != nil {
		return errors.Wrap(err)
	}

	if !isAlreadyApply {
		filenameUp := m.file.PathFileUp(name)
		script, err := m.file.ReadScript(filenameUp)
		if err != nil {
			return errors.Wrap(err)
		}

		err = versionRepo.ExecScript(ctx, script)
		if err != nil {
			return errors.Wrap(err)
		}

		now := time.Now()
		err = versionRepo.UpdateAppliedAt(ctx, name, &now)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) Up(ctx context.Context) error {
	scripts, err := m.file.ListFileScripts()
	if err != nil {
		return errors.Wrap(err)
	}

	names := make([]string, 0, len(scripts))
	for name, script := range scripts {
		if !script.IsValid() {
			return errors.Wrap(fmt.Errorf("migrate(%s): script is invalid", name))
		}

		names = append(names, name)
	}
	sort.Strings(names)

	err = m.insertNames(ctx, names)
	if err != nil {
		return errors.Wrap(err)
	}

	for _, name := range names {
		err = m.execUpScript(ctx, name)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

func (m NativeMigrate) Down(ctx context.Context) error {
	tx, err := m.repo.Data().Begin()
	if err != nil {
		return errors.Wrap(err)
	}
	defer tx.Rollback()

	versionRepo := m.repo.Version(tx.Command())

	name, err := versionRepo.FindLastInserted(ctx)
	if err != nil {
		return errors.Wrap(err)
	}

	filenameDown := m.file.PathFileDown(name)
	script, err := m.file.ReadScript(filenameDown)
	if err != nil {
		return errors.Wrap(err)
	}

	err = versionRepo.ExecScript(ctx, script)
	if err != nil {
		return errors.Wrap(err)
	}

	err = versionRepo.UpdateAppliedAt(ctx, name, nil)
	if err != nil {
		return errors.Wrap(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) Fix(_ context.Context, version int) error {
	return nil
}

package nativemigrate

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type NativeMigrate struct {
	data          repository.DataManager
	configMigrate configs.ConfigMigrate
}

func NewNativeMigrate(
	data repository.DataManager,
	configMigrate configs.ConfigMigrate,
) NativeMigrate {
	return NativeMigrate{
		data:          data,
		configMigrate: configMigrate,
	}
}

func (m NativeMigrate) Create(ctx context.Context, name string) error {
	version := time.Now().Unix()

	filenameUp := fmt.Sprintf("%s/%d_%s.up.sql", m.pathMigrations(), version, name)
	filenameDown := fmt.Sprintf("%s/%d_%s.down.sql", m.pathMigrations(), version, name)

	var err error
	defer func(err error) {
		if err == nil {
			return
		}
		os.Remove(filenameUp)
		os.Remove(filenameDown)
	}(err)

	err = m.createFile(filenameUp)
	if err != nil {
		return errors.Wrap(err)
	}

	err = m.createFile(filenameDown)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) Up(ctx context.Context) error {
	scripts, err := m.listFileScripts()
	if err != nil {
		return errors.Wrap(err)
	}

	err = m.insertScripts(ctx, scripts)
	if err != nil {
		return errors.Wrap(err)
	}

	names := make([]string, 0, len(scripts))
	for name := range scripts {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		isApplied, err := m.alreadyApply(ctx, name)
		if err != nil {
			return errors.Wrap(err)
		}

		if isApplied {
			continue
		}

		filenameUp := fmt.Sprintf("%s/%s.up.sql", m.pathMigrations(), name)

		script, err := m.readScript(filenameUp)
		if err != nil {
			return errors.Wrap(err)
		}

		err = m.execUpScript(ctx, name, script)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

func (m NativeMigrate) Down(ctx context.Context) error {
	scriptName, err := m.findLastInserted(ctx)
	if err != nil {
		return errors.Wrap(err)
	}

	filenameDown := fmt.Sprintf("%s/%s.down.sql", m.pathMigrations(), scriptName)

	script, err := m.readScript(filenameDown)
	if err != nil {
		return errors.Wrap(err)
	}

	err = m.execDownScript(ctx, scriptName, script)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m NativeMigrate) Fix(_ context.Context, version int) error {
	return nil
}

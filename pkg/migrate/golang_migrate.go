package migrate

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/tiagompalte/golang-clean-arch-template/configs"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type GolangMigrate struct {
	configDatabase configs.ConfigDatabase
	configMigrate  configs.ConfigMigrate
}

func NewGolangMigrate(
	config configs.Config,
) Migrate {
	return GolangMigrate{
		configDatabase: config.Database,
		configMigrate:  config.Migrate,
	}
}

func (m GolangMigrate) sourceURL() string {
	return fmt.Sprintf("file://%s", m.configMigrate.PathMigrations)
}

func (m GolangMigrate) databaseURL() string {
	return fmt.Sprintf("%s://%s", m.configDatabase.DriverName, m.configDatabase.ConnectionSource)
}

func (m GolangMigrate) Create(ctx context.Context, name string) error {
	cmd := exec.CommandContext(ctx, "migrate", "create", "-ext", "sql", "-dir", m.configMigrate.PathMigrations, "-seq", name)

	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (m GolangMigrate) Up(_ context.Context) error {
	log.Printf("sourceURL: %s, databaseURL: %s", m.sourceURL(), m.databaseURL())

	mig, err := migrate.New(m.sourceURL(), m.databaseURL())
	if err != nil {
		return errors.Wrap(err)
	}
	defer mig.Close()

	err = mig.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err)
	} else if err == migrate.ErrNoChange {
		log.Printf("no change")
	}

	version, dirty, err := mig.Version()
	if err != nil {
		return errors.Wrap(err)
	}

	log.Printf("version: %d, dirty: %t", version, dirty)

	return nil
}

func (m GolangMigrate) Down(_ context.Context) error {
	log.Printf("sourceURL: %s, databaseURL: %s", m.sourceURL(), m.databaseURL())

	mig, err := migrate.New(m.sourceURL(), m.databaseURL())
	if err != nil {
		return errors.Wrap(err)
	}
	defer mig.Close()

	err = mig.Down()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err)
	} else if err == migrate.ErrNoChange {
		log.Printf("no change")
	}

	version, dirty, err := mig.Version()
	if err != nil {
		return errors.Wrap(err)
	}

	log.Printf("version: %d, dirty: %t", version, dirty)

	return nil
}

func (m GolangMigrate) Fix(_ context.Context, version int) error {
	log.Printf("sourceURL: %s, databaseURL: %s", m.sourceURL(), m.databaseURL())

	mig, err := migrate.New(m.sourceURL(), m.databaseURL())
	if err != nil {
		return errors.Wrap(err)
	}
	defer mig.Close()

	err = mig.Force(version)
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err)
	} else if err == migrate.ErrNoChange {
		log.Printf("no change")
	}

	ver, dirty, err := mig.Version()
	if err != nil {
		return errors.Wrap(err)
	}

	log.Printf("version: %d, dirty: %t", ver, dirty)

	return nil
}

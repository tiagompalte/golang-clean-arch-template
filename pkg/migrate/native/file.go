package nativemigrate

import (
	"os"
	"regexp"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func (m NativeMigrate) createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err)
	}
	defer file.Close()

	return nil
}

func (m NativeMigrate) listFileScripts() (map[string]Script, error) {
	files, err := os.ReadDir(m.pathMigrations())
	if err != nil {
		return map[string]Script{}, errors.Wrap(err)
	}

	mapFiles := make(map[string]Script)
	reg := regexp.MustCompile(`(?<name>.+).(?<type>(?:up|down)).sql`)

	for _, file := range files {
		filename := file.Name()
		if file.IsDir() || !reg.MatchString(filename) {
			continue
		}

		substring := reg.FindStringSubmatch(filename)
		if len(substring) < 3 {
			continue
		}

		name := substring[1]
		script, ok := mapFiles[name]
		if !ok {
			script.Name = name
		}

		switch substring[2] {
		case "up":
			script.HasUp = true
		case "down":
			script.HasDown = true
		}

		mapFiles[name] = script
	}

	return mapFiles, nil
}

func (m NativeMigrate) readScript(filename string) (string, error) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err)
	}

	return string(fileByte), nil
}

package nativemigrate

import (
	"fmt"
	"os"
	"regexp"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FileImpl struct {
	path string
}

func NewFileImpl(path string) File {
	return FileImpl{path: path}
}

func (f FileImpl) createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err)
	}
	defer file.Close()

	return nil
}

func (f FileImpl) CreateUpAndDownFile(version int64, name string) error {
	filenameUp := f.PathFileUp(fmt.Sprintf("%d_%s", version, name))
	filenameDown := f.PathFileDown(fmt.Sprintf("%d_%s", version, name))

	var err error
	defer func(err error) {
		if err == nil {
			return
		}
		os.Remove(filenameUp)
		os.Remove(filenameDown)
	}(err)

	err = f.createFile(filenameUp)
	if err != nil {
		return errors.Wrap(err)
	}

	err = f.createFile(filenameDown)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func (f FileImpl) ListFileScripts() (map[string]Migrate, error) {
	files, err := os.ReadDir(f.path)
	if err != nil {
		return map[string]Migrate{}, errors.Wrap(err)
	}

	mapFiles := make(map[string]Migrate)
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

func (f FileImpl) ReadScript(filename string) (string, error) {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err)
	}

	return string(fileByte), nil
}

func (f FileImpl) PathFileUp(name string) string {
	return fmt.Sprintf("%s/%s.up.sql", f.path, name)
}

func (f FileImpl) PathFileDown(name string) string {
	return fmt.Sprintf("%s/%s.down.sql", f.path, name)
}

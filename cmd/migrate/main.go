package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/config"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/migrate"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("arg command is empty")
	}

	command := os.Args[1]
	if strings.TrimSpace(command) == "" {
		log.Fatalf("arg command is empty")
	}

	conf := config.ProviderSet()
	dataManager := repository.ProviderDataSqlManagerSet(conf)

	mig := migrate.ProviderSet(conf, dataManager)

	var err error
	ctx := context.Background()
	switch command {
	case "up":
		err = mig.Up(ctx)
	case "down":
		err = mig.Down(ctx)
	case "fix":
		var version int
		version, err = strconv.Atoi(os.Args[2])
		if err == nil {
			err = mig.Fix(ctx, version)
		}
	case "create":
		err = mig.Create(ctx, os.Args[2])
	}

	if err != nil {
		log.Fatalf("error to execute comand %s (error: %v)", command, err)
	}

}

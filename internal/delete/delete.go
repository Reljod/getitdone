package delete

import (
	"log"
	"os"
	"strings"

	"github.com/Reljod/getitdone/internal/config"
	"github.com/Reljod/getitdone/internal/read"
)

type Delete struct {
	Read *read.Read
}

func (delete Delete) DeleteCmd(cmdName string) {
	cmdsRaw, err := delete.Read.ReadConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}

	commandsRaw := strings.Split(strings.TrimSpace(cmdsRaw), "\n")
	commandsRawFiltered := []string{}

	for _, c := range commandsRaw {
		splitCmd := strings.Split(c, " => ")
		name := splitCmd[0]
		if name != cmdName {
			commandsRawFiltered = append(commandsRawFiltered, c)
		}
	}

	newCmds := strings.Join(commandsRawFiltered, "\n")

	configPath := config.ConfigPath
	f, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err := f.Write([]byte(newCmds)); err != nil {
		log.Fatalln(err)
		return
	}
}

package read

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/Reljod/getitdone/internal"
	"github.com/Reljod/getitdone/internal/config"
)

type Read struct{}

func (read Read) ReadConfig() (string, error) {
	configPath := config.ConfigPath

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	out, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	return string(out), nil
}

func (read Read) ReadCommands() ([]internal.Command, error) {
	cmdsRaw, err := read.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return []internal.Command{}, err
	}

	var commands []internal.Command = make([]internal.Command, 0)

	commandsRaw := strings.Split(strings.TrimSpace(cmdsRaw), "\n")

	for _, commandRaw := range commandsRaw {
		splitCmd := strings.Split(commandRaw, " => ")
		command := internal.Command{
			Name:    splitCmd[0],
			Command: splitCmd[1],
		}

		commands = append(commands, command)
	}

	return commands, nil
}

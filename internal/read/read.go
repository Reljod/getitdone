package read

import (
	"io"
	"log"
	"os"

	"github.com/Reljod/getitdone/internal/config"
)

func ReadConfig() (string, error) {
	configPath := config.ConfigPath

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	out, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	return string(out), nil
}

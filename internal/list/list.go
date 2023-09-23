package list

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Reljod/getitdone/internal/config"
)

func ListCmds() {
	configPath := config.ConfigPath

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	out, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("\n[ getitdone ls ] Commands list:\n\n%v", string(out))
}

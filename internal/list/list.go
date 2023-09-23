package list

import (
	"fmt"
	"log"

	"github.com/Reljod/getitdone/internal/read"
)

func ListCmds() {
	cmds, err := read.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("\n[ getitdone ls ] Commands list:\n\n%v\n", cmds)
}

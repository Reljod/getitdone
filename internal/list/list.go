package list

import (
	"fmt"
	"log"

	"github.com/Reljod/getitdone/internal/read"
)

type List struct {
	Read *read.Read
}

func (list List) ListCmds() {
	cmds, err := list.Read.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("\n[ getitdone ls ] Commands list:\n\n%v\n", cmds)
}

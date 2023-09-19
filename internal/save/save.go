package save

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

func Save(name string, command *string) error {
	cmd := *command

	if *command == "" {
		fmt.Println("Getting last command")
		var err error
		cmd, err = getLastCommand()
		if err != nil {
			error := fmt.Errorf("cannot get last command - %v", err)
			fmt.Printf("%v\n", error)
			return error
		}
	}

	fmt.Printf("%v => %v\n", name, cmd)

	return nil
}

func getLastCommand() (string, error) {
	c1 := exec.Command("bash", "-c", "cat $HOME/.zsh_history")
	c2 := exec.Command("tail", "-2")
	c3 := exec.Command("head", "-1")

	var err error
	c2.Stdin, _ = c1.StdoutPipe()
	c3.Stdin, _ = c2.StdoutPipe()

	var b3 bytes.Buffer
	c3.Stdout = &b3

	_ = c2.Start()

	err = c1.Run()
	if err != nil {
		error := fmt.Errorf("cannot execute history command - %v", err)
		return "", error
	}

	err = c3.Start()
	_ = c2.Wait()
	if err != nil {
		error := fmt.Errorf("cannot execute tail command - %v", err)
		return "", error
	}

	err = c3.Wait()
	if err != nil {
		error := fmt.Errorf("cannot execute echo command - %v", err)
		return "", error
	}

	reg := regexp.MustCompile(`.+;(.+)`)
	res := reg.ReplaceAllString(b3.String(), "${1}")
	return res, nil
}

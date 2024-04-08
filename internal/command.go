package internal

import (
    "errors"
    "fmt"
    "os/exec"
    "strings"
)


func ExecCommand(commandString string) error {
	command := strings.Split(commandString, " ")
	if len(command) < 2 {
        return errors.New("Для команды обязательно передать хотя бы 1 аргумент")
	}
    cmd := exec.Command(command[0], command[1:]...)
    stdout, err := cmd.Output()

    if err != nil {
        return err
    }

    fmt.Println(string(stdout))

    return nil
}
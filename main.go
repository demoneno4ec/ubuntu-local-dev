package main

import (
    "fmt"
    "os/exec"

    "github.com/fatih/color"
    "demoneno4ec/ubuntu-local-dev/config"
)


func main() {
	config := config.GetFromYaml()
    fmt.Println(config)
	updateDependencies()
	addingPpaRepositories()
	installAptPackages()
	installSnapPackages()
	installCustomPackages()
	configPackages()
// test line
	exampleExecCommand()
}


func updateDependencies() {
	color.Cyan("Prints text in cyan.")
}

func addingPpaRepositories() {
	color.Cyan("Prints text in cyan.")
}

func installAptPackages() {
	color.Cyan("Prints text in cyan.")
}

func installSnapPackages() {
	color.Cyan("Prints text in cyan.")
}

func installCustomPackages() {
	color.Cyan("Prints text in cyan.")
}

func configPackages() {
	color.Cyan("Prints text in cyan.")
}

func exampleExecCommand() {
	app := "echo"

    arg0 := "-e"
    arg1 := "Hello world"
    arg2 := "\n\tfrom"
    arg3 := "golang"

    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(stdout))

}

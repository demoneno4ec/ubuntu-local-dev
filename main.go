package main

import (
    _ "fmt"

    "github.com/fatih/color"
    "demoneno4ec/ubuntu-local-dev/config"
    "demoneno4ec/ubuntu-local-dev/internal"
)


func main() {
	config := config.GetFromYaml()
	err := addingRepositories(config.Repositories)

    if err != nil {
        color.Red(err.Error())
        return
    }
	updateDependencies()
	installAptPackages()
	installSnapPackages()
	installCustomPackages()
	configPackages()
    // test line
	err = internal.ExecCommand("echo -e Hello go!")

    if err != nil {
        color.Red(err.Error())
        return
    }
}

func addingRepositories(repositories config.Repositories) error {
    color.Cyan("Добавление репозиториев")
    for _, repository := range repositories.AddApt {
        color.Cyan(repository)
    }

    return nil
}

func updateDependencies() {
	color.Cyan("Обновление зависимостей")
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

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

func addingRepositories(repositories internal.Repository[]) error {
    /*
    * TODO
    * Надо сделать следующее поведение
    * 1. из yaml мы должны получать slice internal/repositories/repository сразу с установленными типами
    *     1.1 наиболее вероятно в unmarshal нет возможности динамически сетить преобразовывать в нужные типы, сквозь рефлексию и нужно будет это доописать, как итог конфиг на выходе и изначальная структура Config для unmarshal будет отличаться
    *     1.2 возможно нужно возвращать некоторую DTO
    * 2. поменять структуру функции на 
    *     func addingRepositories(repositories []internal.Repository) error {}
    * 3. примерное поведение for _, repository := range repositories
    *     err := repository.Add()
    *     if err != nil {
    *         return err
    *     }
    *
    *
    *
    *
    */
    color.Cyan("Добавление репозиториев")
    for _, repository := range repositories {
        color.Cyan(repository.GetName())
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

package config

import (
    "log"
    "os"

    "gopkg.in/yaml.v2"

    "github.com/fatih/color"
    "demoneno4ec/ubuntu-local-dev/internal"
)

type Config struct {
    Packages Packages `yaml:"packages"`
    Repositories Repositories `yaml:"repositories"`
    Settings Settings `yaml:"settings"`
}

type Packages struct {
    Apt []string `yaml:"apt"`
    Snap []string `yaml:"snap"`
    Ppa []string `yaml:"ppa"`
    Deb []string `yaml:"deb"`
}

type Repositories struct {
    AddApt []string `yaml:"add-apt"`
    Dpkg []string `yaml:"dpkg"`
    Custom []string `yaml:"custom"`
}

type Settings struct {
    Git Git `yaml: "git"`
}

type Git struct {
    Name string `yaml: "name"`
    Email string `yaml: "email"`
    SshDir string `yaml: "ssh_dir"`
}

func GetFromYaml() internal.Config {
    var yamlConfig Config
    var config internal.Config

    color.Cyan("Получаем конфигурацию скрипта из yaml")

    yamlFile, err := os.ReadFile("./config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }

    err = yaml.Unmarshal(yamlFile, &yamlConfig)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    var repositories internal.Repository[]
    for _, repository := range yamlConfig.Repositories.AddApt {
        repositories = append(repositories, repository)
    }

    color.Green("Конфиги успешно установлены")

    return config
}
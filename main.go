package main

import (
    "fmt"
    "os/exec"
	"github.com/fatih/color"
)

func main() {
	setConfigFromYaml()
	updateDependencies()
	addingPpaRepositories()
	installAptPackages()
	installSnapPackages()
	installCustomPackages()
	configPackages()
// test line
	exampleExecCommand()
}

func setConfigFromYaml() {
	color.Cyan("Prints text in cyan.")
type conf struct {
    Hits int64 `yaml:"hits"`
    Time int64 `yaml:"time"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := os.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c conf
    c.getConf()

    	fmt.Println(c)
	}
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

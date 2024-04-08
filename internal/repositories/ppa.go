package "internal/repositories"

type Ppa struct {
	Name string
}

func (p Ppa) Add() error {
	color.Green("Успешно добавленный репозиторий")
}
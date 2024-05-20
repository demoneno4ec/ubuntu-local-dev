package internal

type Ppa struct {
	name string
}

func (p Ppa) Add() error {
	color.Green("Успешно добавленный репозиторий")
}

func (p Ppa) GetName() string {
	return name
}
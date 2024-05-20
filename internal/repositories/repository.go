package internal

type Repository interface {
	Add() error
	GetName() string
}
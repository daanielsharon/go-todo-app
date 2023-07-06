package repository

type TodoRepository interface {
	Save()
	Delete()
	FindByName()
	FindAll()
}
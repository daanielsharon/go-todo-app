package domain

type Todo struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	ContainerId int `json:"container_id"`
}
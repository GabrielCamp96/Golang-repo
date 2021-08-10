package models

type usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func NewUsuario() (u usuario) {
	return
}

func NewSliceUsuarios() (u []usuario) {
	return
}

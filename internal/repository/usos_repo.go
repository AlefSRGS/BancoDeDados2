package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"

type UsosRepository interface {
	Create(usos model.Usos) error
    GetByID(id uint) (model.Usos, error)
    Update(usos model.Usos) error
    Delete(id uint) error
    GetAll() ([]model.Usos, error)
}
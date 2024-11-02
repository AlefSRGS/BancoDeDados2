package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"



type PratoRepository interface {
	Create(prato model.Prato) error
    GetByID(id uint) (model.Prato, error)
    Update(prato model.Prato) error
    Delete(id uint) error
    GetAll() ([]model.Prato, error)
}
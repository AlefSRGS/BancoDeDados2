package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"

type ClienteRepository interface {
	Create(cliente model.Cliente) error
    GetByID(id uint) (model.Cliente, error)
    Update(cliente model.Cliente) error
    Delete(id uint) error
    GetAll() ([]model.Cliente, error)
}
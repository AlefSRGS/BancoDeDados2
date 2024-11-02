package repository

import "github.com/vinicius457/BancoDeDados2/internal/model"

type FornecedorRepository interface {
	Create(fornecedor model.Fornecedor) error
    GetByID(id uint) (model.Fornecedor, error)
    Update(fornecedor model.Fornecedor) error
    Delete(id uint) error
    GetAll() ([]model.Fornecedor, error)
	
}
package repository

import (
	"time"
	"gorm.io/gorm"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

func InsertClientes(db *gorm.DB) error {
	clientes := []model.Cliente{
		{Nome: "João", Sexo: "m", Idade: 30, Nascimento: time.Date(1994, 2, 15, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Maria", Sexo:  "f", Idade: 25, Nascimento: time.Date(1999, 10, 22, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Carlos", Sexo: "m", Idade: 35, Nascimento: time.Date(1989, 6, 10, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Ana", Sexo: "f", Idade: 40, Nascimento: time.Date(1984, 3, 5, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Pedro", Sexo: "m", Idade: 28, Nascimento: time.Date(1996, 9, 12, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Beatriz", Sexo: "f", Idade: 22, Nascimento: time.Date(2002, 1, 18, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Ricardo", Sexo: "m", Idade: 45, Nascimento: time.Date(1978, 11, 30, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Fernanda", Sexo: "f", Idade: 32, Nascimento: time.Date(1992, 4, 21, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Lucas", Sexo: "m", Idade: 29, Nascimento: time.Date(1995, 12, 8, 0, 0, 0, 0, time.UTC), Pontos: 0},
		{Nome: "Sofia", Sexo: "f", Idade: 27, Nascimento: time.Date(1997, 7, 15, 0, 0, 0, 0, time.UTC), Pontos: 0},
	}

	return db.Create(&clientes).Error
}

func InsertPratos(db *gorm.DB) error {
	pratos := []model.Prato{
		{Nome: "Pizza Margherita", Descricao: "Queijo, tomate e manjericão", Valor: 35.00, Disponibilidade: true},
		{Nome: "Lasanha", Descricao: "Massa, queijo e molho de carne", Valor: 45.00, Disponibilidade: true},
		{Nome: "Salada Caesar", Descricao: "Alface, croutons, frango e molho Caesar", Valor: 25.00, Disponibilidade: true},
		{Nome: "Hambúrguer", Descricao: "Pão, carne, queijo e salada", Valor: 20.00, Disponibilidade: true},
		{Nome: "Sopa de Legumes", Descricao: "Sopa com diversos legumes frescos", Valor: 18.00, Disponibilidade: true},
		{Nome: "Frango Assado", Descricao: "Frango temperado e assado", Valor: 30.00, Disponibilidade: true},
		{Nome: "Espaguete", Descricao: "Espaguete com molho de tomate", Valor: 28.00, Disponibilidade: true},
		{Nome: "Bife à Parmegiana", Descricao: "Bife, queijo e molho de tomate", Valor: 40.00, Disponibilidade: true},
		{Nome: "Risoto de Camarão", Descricao: "Arroz com camarão e temperos", Valor: 55.00, Disponibilidade: true},
		{Nome: "Torta de Limão", Descricao: "Sobremesa doce de limão", Valor: 15.00, Disponibilidade: true},
	}

	return db.Create(&pratos).Error
}

func InsertFornecedores(db *gorm.DB) error {
	fornecedores := []model.Fornecedor{
		{Nome: "Fornecedor A", EstadoOrigem: "SP"},
		{Nome: "Fornecedor B", EstadoOrigem: "RJ"},
		{Nome: "Fornecedor C", EstadoOrigem: "MG"},
		{Nome: "Fornecedor D", EstadoOrigem: "BA"},
		{Nome: "Fornecedor E", EstadoOrigem: "PR"},
		{Nome: "Fornecedor F", EstadoOrigem: "SC"},
		{Nome: "Fornecedor G", EstadoOrigem: "PE"},
		{Nome: "Fornecedor H", EstadoOrigem: "RS"},
		{Nome: "Fornecedor I", EstadoOrigem: "GO"},
		{Nome: "Fornecedor J", EstadoOrigem: "ES"},
	}

	return db.Create(&fornecedores).Error
}


func InsertIngredientes(db *gorm.DB) error {
	ingredientes := []model.Ingrediente{
		{Nome: "Tomate", DataFabricacao: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), Quantidade: 100, Observacao: "Fresco"},
		{Nome: "Alface", DataFabricacao: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), Quantidade: 50, Observacao: "Orgânico"},
		{Nome: "Queijo", DataFabricacao: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), Quantidade: 200, Observacao: "Mussarela"},
		{Nome: "Frango", DataFabricacao: time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC), Quantidade: 150, Observacao: "Congelado"},
		{Nome: "Carne", DataFabricacao: time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC), Quantidade: 180, Observacao: "Fresca"},
		{Nome: "Manjericão", DataFabricacao: time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC), Quantidade: 70, Observacao: "Fresco"},
		{Nome: "Camarão", DataFabricacao: time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC), Quantidade: 90, Observacao: "Descongelado"},
		{Nome: "Arroz", DataFabricacao: time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 5, 5, 0, 0, 0, 0, time.UTC), Quantidade: 300, Observacao: "Integral"},
		{Nome: "Limão", DataFabricacao: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC), Quantidade: 120, Observacao: "Taiti"},
		{Nome: "Farinha", DataFabricacao: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC), DataValidade: time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC), Quantidade: 250, Observacao: "De trigo"},
	}

	return db.Create(&ingredientes).Error
}


func InsertVendas(db *gorm.DB) error {
	vendas := []model.Venda{
		{IDCliente: 1, IDPrato: 1, Quantidade: 2, Dia: time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 1, 12, 30, 0, 0, time.UTC), Valor: 70.00},
		{IDCliente: 2, IDPrato: 2, Quantidade: 1, Dia: time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 2, 13, 15, 0, 0, time.UTC), Valor: 45.00},
		{IDCliente: 3, IDPrato: 3, Quantidade: 3, Dia: time.Date(2024, 3, 3, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 3, 14, 0, 0, 0, time.UTC), Valor: 75.00},
		{IDCliente: 4, IDPrato: 4, Quantidade: 1, Dia: time.Date(2024, 3, 4, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 4, 15, 45, 0, 0, time.UTC), Valor: 20.00},
		{IDCliente: 5, IDPrato: 5, Quantidade: 2, Dia: time.Date(2024, 3, 5, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 5, 16, 20, 0, 0, time.UTC), Valor: 36.00},
		{IDCliente: 6, IDPrato: 6, Quantidade: 1, Dia: time.Date(2024, 3, 6, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 6, 17, 35, 0, 0, time.UTC), Valor: 30.00},
		{IDCliente: 7, IDPrato: 7, Quantidade: 2, Dia: time.Date(2024, 3, 7, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 7, 18, 50, 0, 0, time.UTC), Valor: 80.00},
		{IDCliente: 8, IDPrato: 8, Quantidade: 3, Dia: time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 8, 19, 10, 0, 0, time.UTC), Valor: 120.00},
		{IDCliente: 9, IDPrato: 9, Quantidade: 2, Dia: time.Date(2024, 3, 9, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 9, 20, 25, 0, 0, time.UTC), Valor: 56.00},
		{IDCliente: 10, IDPrato: 10, Quantidade: 1, Dia: time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC), Hora: time.Date(2024, 3, 10, 21, 40, 0, 0, time.UTC), Valor: 15.00},
	}

	return db.Create(&vendas).Error
}

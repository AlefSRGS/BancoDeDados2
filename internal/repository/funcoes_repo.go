package repository

import (
	"database/sql"
)

type FuncoesRepo struct {
	db *sql.DB
}

func NewFuncoesRepo(db *sql.DB) *FuncoesRepo {
	return &FuncoesRepo{db: db}
}

// Aplica reajuste percentual nos pratos
func (r *FuncoesRepo) Reajuste(percentual float64) error {
	_, err := r.db.Exec("CALL reajuste($1)", percentual)
	return err
}

// Sorteia um cliente e adiciona pontos
func (r *FuncoesRepo) SorteioCliente() error {
	_, err := r.db.Exec("CALL sorteio_cliente()")
	return err
}

// Retorna estatísticas de vendas
func (r *FuncoesRepo) EstatisticasVendas() (string, string, error) {
	var maisVendido, menosVendido string
	err := r.db.QueryRow("CALL estatisticas_vendas()").Scan(&maisVendido, &menosVendido)
	if err != nil {
		return "", "", err
	}
	return maisVendido, menosVendido, nil
}

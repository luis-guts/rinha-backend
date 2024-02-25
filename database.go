package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=rinha2024q1! dbname=rinhabank sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}

func salvarTransacao(idCliente int64, transacao *Transacao) (*Cliente, error) {
	con, err := OpenConn()
	if err != nil {
		return nil, err
	}
	defer con.Close()
	cliente := &Cliente{}
	sql := "SELECT * FROM cliente WHERE id = $1"
	row := con.QueryRow(sql, idCliente)
	err = row.Scan(&cliente.Id, &cliente.Limite, &cliente.SaldoInicial)
	if err != nil {
		return nil, err
	}
	return cliente, nil
}

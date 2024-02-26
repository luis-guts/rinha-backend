package main

type Transacao struct {
	Valor     float64 `json:"valor"`
	Tipo      string  `json:"tipo"`
	Descricao string  `json:"descricao"`
}

type ResultadoTransacao struct {
	Limite string `json:"limite"`
	Saldo  string `json:"saldo"`
}

type Cliente struct {
	Id           int     `json:"id"`
	Limite       float64 `json:"limite"`
	SaldoInicial float64 `json:"saldo_inicial"`
}

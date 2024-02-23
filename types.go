package main

type Transacao struct {
	valor     string
	tipo      string
	descricao string
}

type ResultadoTransacao struct {
	Limite string `json:"limite"`
	Saldo  string `json:"saldo"`
}

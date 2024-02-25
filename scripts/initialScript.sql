CREATE DATABASE rinhabank;


\c rinhabank;

CREATE TABLE cliente (
    "id" SERIAL PRIMARY KEY,
    "limite" NUMERIC(12,2),
    "saldo_inicial" NUMERIC(12,2)
);

CREATE TABLE transacao(
    "id" SERIAL PRIMARY KEY,
    "id_cliente" INTEGER REFERENCES cliente(id),
    "descricao" varchar(20),
    "realizada_em" DATE
);

INSERT INTO cliente("limite", "saldo_inicial") VALUES (100000, 0);

INSERT INTO cliente("limite", "saldo_inicial") VALUES (80000, 0);

INSERT INTO cliente("limite", "saldo_inicial") VALUES (1000000, 0);

INSERT INTO cliente("limite", "saldo_inicial") VALUES (10000000, 0);

INSERT INTO cliente("limite", "saldo_inicial") VALUES (500000, 0);

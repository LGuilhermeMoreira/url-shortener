# URL Shortener

[![Project Link](https://img.shields.io/badge/Project-Link-blue)](https://url-shortener-1-q6pg.onrender.com)

Este é um simples projeto de encurtador de URL com as seguintes rotas:

## Rotas

- **`/`**: Página inicial que renderiza um front-end simples usando templates.
- **`/redirecionar/{id}`**: Redireciona para a página associada ao ID fornecido.

## Demonstração

Você pode acessar a aplicação [aqui](https://url-shortener-1-q6pg.onrender.com).

## Como Usar

1. Vá para a página inicial.
2. Insira a URL que você deseja encurtar.
3. Receba uma URL curta que pode ser compartilhada e usada para redirecionar para a URL original.

## Tecnologias Utilizadas

- **Back-end**: Go
- **Front-end**: Templates HTML
- **Banco de Dados**: PostgreSQL
- **Outros**: GORM, Air para hot-reloading

## Estrutura do Projeto

```plaintext
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   └── config_test.go
├── internal
│   └── infra
│       └── database
│           ├── connection.go
│           └── connection_test.go
├── .env
├── README.md
└── ...

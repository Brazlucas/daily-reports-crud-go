# Report Manager

Este é um pequeno sistema de gerenciamento de reports/anotações diário(a)s, desenvolvido em Go, que utiliza SQLite para armazenar os registros e uma interface web estilizada com Tailwind CSS.

## Funcionalidades

- Adicionar reports
- Listar reports
- Editar reports
- Excluir reports
- Interface web responsiva

## Requisitos

- Go instalado ([Download](https://go.dev/dl/))
- SQLite3 instalado

## Instalação

Clone este repositório e acesse a pasta do projeto:

```sh
git clone https://github.com/Brazlucas/daily-reports-crud-go.git
cd daily-reports-crud-go
```

Instale as dependências:

```sh
go get github.com/mattn/go-sqlite3
go mod tidy
```

## Executando o Projeto

Basta rodar o seguinte comando:

```sh
go run main.go
```

Isso abrirá automaticamente o navegador com a interface web.

## Estrutura do Projeto

```
daily-reports-crud-go/

│── main.go         # Código principal do servidor
│── reports.db      # Banco de dados SQLite
│── README.md       # Este arquivo de documentação
|── database        # database
  |── database.go
  |── models.go
|── handlers        # handlers
  |── home.go
  |── report.go
|── templates       # views
  |── edit.html
  |── index.html
|── utils           # utils
  |── browser.go
```

## Tecnologias Utilizadas

- Go (Golang)
- Html
- SQLite3
- Tailwind CSS

## Licença

Este projeto está sob a licença MIT. Sinta-se à vontade para contribuir e melhorar!

---
Desenvolvido por [Lucas](https://github.com/Brazlucas) 🚀


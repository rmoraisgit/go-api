# go-api

Este projeto é uma API simples desenvolvida em Go utilizando o framework [Gin](https://github.com/gin-gonic/gin).

## Descrição

O código implementa um servidor HTTP que expõe endpoints para manipulação de produtos:

- `/ping`: Retorna um JSON com a mensagem `"pong"`.
- `/produtos` (GET): Retorna uma lista de produtos em formato JSON, buscando os dados no banco PostgreSQL.
- `/produtos` (POST): Cria um novo produto no banco de dados a partir de um JSON enviado no corpo da requisição.
- `/produtos/:id` (GET): Retorna um produto específico pelo ID.

## Como executar

1. Certifique-se de ter o Go instalado (versão 1.24.5 ou superior).
2. Instale as dependências com:

   ```sh
   go mod tidy
   ```

3. Execute o servidor:

   ```sh
   go run ./cmd/main.go
   ```

4. O servidor estará disponível em [http://localhost:8080](http://localhost:8080).

## Utilizando o PostgreSQL com Docker Compose

O projeto inclui um arquivo `docker-compose.yml` para facilitar a execução de um banco de dados PostgreSQL localmente.

### Subindo o banco de dados

1. Certifique-se de ter o [Docker](https://www.docker.com/) e o [Docker Compose](https://docs.docker.com/compose/) instalados.
2. Execute o comando abaixo na raiz do projeto para subir o PostgreSQL:

   ```sh
   docker-compose up -d
   ```

3. O banco estará disponível na porta 5432. As credenciais padrão podem ser ajustadas no arquivo `docker-compose.yml`.

### Criando a tabela e inserindo dados

Após o banco estar rodando, conecte-se ao PostgreSQL (por exemplo, usando o DBeaver, TablePlus, ou o utilitário `psql`) e execute os comandos abaixo para criar a tabela e inserir um produto de exemplo:

```sql
create table produtos (
    id SERIAL primary key,
    nome varchar(50) not null,
    preco numeric(10, 2) not null
);

insert into produtos(nome, preco) values ('iphone', 5000);

select * from produtos;
```

## Endpoints

- `GET /ping`  
  **Resposta:**  
  ```json
  {
    "message": "pong"
  }
  ```

- `GET /produtos`  
  **Resposta:**  
  ```json
  [
    {
      "id": 1,
      "nome": "Produto A",
      "preco": 10.0
    },
    {
      "id": 2,
      "nome": "Produto B",
      "preco": 20.0
    }
  ]
  ```

- `POST /produtos`  
  **Request Body:**  
  ```json
  {
    "nome": "Produto Novo",
    "preco": 99.99
  }
  ```
  **Resposta:**  
  ```json
  {
    "id": 3,
    "nome": "Produto Novo",
    "preco": 99.99
  }
  ```

- `GET /produtos/:id`  
  Retorna um produto específico pelo ID.  
  **Exemplo de resposta:**  
  ```json
  {
    "id": 1,
    "nome": "Produto A",
    "preco": 10.0
  }
  ```

## Estrutura do Projeto

- [`cmd/main.go`](cmd/main.go): Arquivo principal que inicializa e executa o servidor.
- [`internal/controllers/produtos_controller.go`](internal/controllers/produtos_controller.go): Controller responsável pelos endpoints de produtos.
- [`internal/models/produto.go`](internal/models/produto.go): Model que representa o produto (struct `Produto`).
- [`internal/repositories/produtos_repository.go`](internal/repositories/produtos_repository.go): Repository responsável por acessar os dados dos produtos no banco.
- [`docker-compose.yml`](docker-compose.yml): Arquivo de configuração para subir o PostgreSQL via Docker.

## Dependências

- [Gin](https://github.com/gin-gonic/gin): Framework web para Go.
- [PostgreSQL](https://www.postgresql.org/): Banco de dados relacional utilizado
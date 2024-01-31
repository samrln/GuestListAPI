# Lista de Convidados - API Simples

Este é um projeto de uma API simples para gerenciar uma lista de convidados para um evento. A API permite adicionar, atualizar, listar e confirmar a presença dos convidados.

## Funcionalidades

- Adicionar um novo convidado
- Atualizar o status de confirmação de um convidado
- Listar todos os convidados
- Listar convidados confirmados
- Listar convidados não confirmados
- Obter detalhes de um convidado específico

## Tecnologias Utilizadas

- Linguagem: Go (Golang)
- Framework: Gorilla Mux
- Banco de Dados: Não utilizado (armazenamento em memória)

## Instalação e Uso

1. Clone o repositório:

    ```bash
    git clone https://github.com/samrln/guestListAPI.git
    ```

2. Navegue até o diretório do projeto:

    ```bash
    cd guestListAPI
    ```

3. Execute o projeto:

    ```bash
    go run ./cmd/main.go
    ```

4. Acesse a API em `http://localhost:8080`

## Rotas da API

- `GET /getAllList`: Obtém a lista completa de convidados.
- `GET /getGuest/{name}`: Obtém detalhes de um convidado pelo nome.
- `GET /getConfirmed`: Obtém a lista de convidados confirmados.
- `GET /getNotConfirmed`: Obtém a lista de convidados não confirmados.
- `POST /addGuests`: Adiciona um novo convidado.
- `PUT /updateConfirmed/{name}`: Atualiza o status de confirmação de um convidado pelo nome.

## Contribuindo

Se deseja contribuir com este projeto, siga estes passos:

1. Faça um fork do projeto
2. Crie uma branch com sua feature: `git checkout -b feature-nova-feature`
3. Faça commit das suas mudanças: `git commit -m 'Adiciona nova feature'`
4. Faça push para a branch: `git push origin feature-nova-feature`
5. Envie um Pull Request

## Próximos Passos

- Adicionar testes automatizados para garantir a estabilidade do código.
- Implementar persistência de dados utilizando um banco de dados.
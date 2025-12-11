# Jogo da Velha 2 (Ultimate Tic-Tac-Toe)

[🇺🇸 Read in English](./README.md)

**Uma implementação estratégica e multiplayer em tempo real do Jogo da Velha Ultimate.**

Desafie seus amigos em um jogo onde cada movimento conta não apenas para o tabuleiro atual, mas dita onde seu oponente deve jogar em seguida. Construído com foco em desempenho e código limpo usando Go e WebSockets.

---

## ✨ Funcionalidades

- **🎮 Mecânica Ultimate**: Aproveite a camada estratégica profunda da variante "Ultimate" do Jogo da Velha.
- **⚡ Multiplayer em Tempo Real**: Jogabilidade fluida impulsionada por **Gorilla WebSockets**.
- **🖥️ UI Limpa e Responsiva**: Construída com **HTML5** semântico e **CSS3**, renderizada via **Go Templates**.
- **🚀 Backend Rápido**: Servidor de alta performance construído com o **Gin Framework**.

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** [Go (Golang)](https://go.dev/)
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Comunicação em Tempo Real:** [Gorilla WebSockets](https://github.com/gorilla/websocket)
- **Frontend:** HTML / CSS / Go Templates

## 🚀 Começando

Siga estas etapas para executar o projeto em sua máquina local.

### Pré-requisitos

- [Go](https://go.dev/dl/) (versão 1.25 ou superior)
- [Git](https://git-scm.com/)

### Instalação

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/lucasramosdev/jogo-da-velha-dois.git
   cd jogo-da-velha-dois
   ```

2. **Instale as dependências:**
   ```bash
   go mod download
   ```

3. **Execute a aplicação:**
   ```bash
   go run cmd/app/main.go
   ```

4. **Jogue:**
   Abra seu navegador e acesse `http://localhost:8080`.

## 📂 Estrutura do Projeto

- **`cmd/`**: Pontos de entrada da aplicação.
- **`internal/`**: Código da aplicação e bibliotecas privadas.
  - **`game/`**: Lógica central do jogo e gerenciamento de estado.
  - **`web/`**: Manipuladores HTTP e roteamento.
- **`web/`**: Assets do frontend (Templates HTML, CSS, JS).

## 📄 Licença

Este projeto é open-source e está disponível sob a [Licença MIT](LICENSE).

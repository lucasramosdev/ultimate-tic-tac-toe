# Jogo da Velha 2 (Ultimate Tic-Tac-Toe)

[🇧🇷 Leia em Português](./README.pt-br.md)

**A strategic, real-time multiplayer implementation of Ultimate Tic-Tac-Toe.**

Challenge your friends to a game where every move counts not just for the current board, but dictates where your opponent must play next. Built with performance and clean code in mind using Go and WebSockets.

---

## ✨ Features

- **🎮 Ultimate Tic-Tac-Toe Mechanics**: Enjoy the deep strategic layer of the "Ultimate" variant.
- **⚡ Real-time Multiplayer**: Seamless gameplay powered by **Gorilla WebSockets**.
- **🖥️ Clean & Responsive UI**: Built with semantic **HTML5** and **CSS3**, rendered via **Go Templates**.
- **🚀 Fast Backend**: High-performance server built with the **Gin Framework**.

## 🛠️ Tech Stack

- **Language:** [Go (Golang)](https://go.dev/)
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Real-time Communication:** [Gorilla WebSockets](https://github.com/gorilla/websocket)
- **Frontend:** HTML / CSS / Go Templates

## 🚀 Getting Started

Follow these steps to get the project running on your local machine.

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.25 or higher)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/lucasramosdev/jogo-da-velha-dois.git
   cd jogo-da-velha-dois
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the application:**
   ```bash
   go run cmd/app/main.go
   ```

4. **Play:**
   Open your browser and navigate to `http://localhost:8080`.

## 📂 Project Structure

- **`cmd/`**: Application entry points.
- **`internal/`**: Private application and library code.
  - **`game/`**: Core game logic and state management.
  - **`web/`**: HTTP handlers and routing.
- **`web/`**: Frontend assets (HTML templates, CSS, JS).

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).

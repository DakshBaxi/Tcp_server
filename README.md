# Go TCP Chat Server 💬

A simple, concurrent TCP chat server and client application written in Go. This project demonstrates the use of goroutines for handling multiple clients, `sync.Map` for concurrent data access, and basic TCP networking in Go.

![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## 🌟 Features

- **Real-time Group Chat**: Messages are broadcast to all connected clients instantly
- **User Identification**: Each user is prompted for a username upon connecting, which is prepended to their messages
- **Presence Notifications**: The chat room automatically announces when a user joins or leaves
- **Concurrent Connections**: The server uses goroutines to handle multiple clients simultaneously without blocking
- **Safe Concurrency**: A `sync.Map` is used to safely manage active client connections across multiple goroutines

## 📋 Table of Contents

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation & Running](#installation--running)
- [Example Session](#example-session)
- [How It Works](#how-it-works)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## 🚀 Getting Started

Follow these instructions to get the chat server and client running on your local machine.

### Prerequisites

You must have Go installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

Verify your installation:
```bash
go version
```

### Installation & Running

1. **Clone the repository**:
```bash
git clone https://github.com/yourusername/go-tcp-chat-server.git
cd go-tcp-chat-server
```

2. **Run the Server**:
```bash
go run server.go
```

You should see the following output, indicating the server is ready:
```
Listening on port 8080
```

3. **Run the Client(s)**:

Open one or more new terminal windows and run:
```bash
go run client/client.go
```

- The client will first ask for your name
- After entering your name, you'll be connected to the chat
- Repeat this step in multiple terminals to simulate a multi-user chat room

## 📸 Example Session

**Terminal 1 (Server):**
```
Listening on port 8080
Client IP address: 127.0.0.1
Client IP address: 127.0.0.1
```

**Terminal 2 (Client 1):**
```
Please enter your name: Alice
You are now connected. Type a message and press Enter.
--- Bob has joined the chat ---
Hello Bob!
```

**Terminal 3 (Client 2):**
```
Please enter your name: Bob
You are now connected. Type a message and press Enter.
--- Alice has joined the chat ---
Alice: Hello Bob!
Hey Alice!
```

## 🛠 How It Works

### Server (`server.go`)

1. The server starts by listening for TCP connections on port **8080**
2. Uses a `sync.Map` to store active connections (safe for concurrent reads and writes)
3. In an infinite loop, the server calls `ln.Accept()` to wait for new client connections
4. For each new connection, it spawns a dedicated goroutine running the `handleConn` function
5. The `handleConn` function:
   - Reads the client's chosen username
   - Broadcasts a "joined" message to all other clients
   - Uses a `defer` block to broadcast a "left" message when the client disconnects
   - Enters a loop reading messages from the client and broadcasting them with the username prepended

### Client (`client/client.go`)

1. Connects to the server at `localhost:8080`
2. Prompts the user to enter a username and sends it to the server
3. Launches a separate goroutine (`receiveMessages`) to read data from the server and print it to the console
4. The main goroutine reads user input from the keyboard and sends it to the server
5. This two-goroutine approach enables simultaneous sending and receiving of messages without blocking

## 📁 Project Structure

```
chat-project/
├── server.go           # TCP chat server
├── client/
│   └── client.go      # TCP chat client
└── README.md          # This file
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with Go's standard library
- Inspired by the need for simple, educational chat server examples
- Thanks to the Go community for excellent documentation

---

**Made with ❤️ and Go**

If you found this project helpful, please consider giving it a ⭐!
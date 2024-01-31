package main

import (
    "guestListAPI/server"
)

func main() {
    port := ":8080"

    server.StartServer(port)
}

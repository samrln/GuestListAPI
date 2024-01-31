package server

import (
    "log"
    "net/http"
    "guestListAPI/internal/routes"
	"guestListAPI/server/signals"
	"guestListAPI/server/logger"
)

func StartServer(port string) {
    router := routes.InitRoutes()

    logger.Info("Server running on port: http://localhost%s", port)
    logger.Info("Get all list: http://localhost%s/GetAllList", port)
    logger.Info("Press Ctrl+C to stop server")

    // capture signal interrupt
    stop := signals.SetupSignalHandler()

    // init server HTTP in new goroutine
    go func() {
        if err := http.ListenAndServe(port, router); err != nil {
            log.Fatalf("Error to init server: %v", err)
        }
    }()

    // wait signal interrupt
    <-stop
    logger.Info("Server stopped")
}

package signals

import (
    "os"
    "os/signal"
    "syscall"
)

func SetupSignalHandler() chan os.Signal {
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    return stop
}

package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/urfave/cli/v2"
)

const (
	flagDirName  = "dir"
	flagPortName = "port"
)

var (
	flagDir = &cli.StringFlag{
		Name:        flagDirName,
		Aliases:     []string{"d"},
		Usage:       "The root directory of server",
		DefaultText: ".",
	}
	flagPort = &cli.UintFlag{
		Name:        flagPortName,
		Aliases:     []string{"p"},
		Usage:       "The http port",
		DefaultText: "8000",
	}
)

const (
	defaultPort = 8000
	defaultDir  = "."
)

func actionHandler(c *cli.Context) error {
	dir := defaultDir
	var port uint
	if c.IsSet(flagDirName) {
		dir = c.String(flagDirName)
	}

	if c.IsSet(flagPortName) {
		port = c.Uint(flagPortName)
		if port == 0 || port > 65535 {
			port = defaultPort
		}
	} else {
		port = defaultPort
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	go func() {
		select {
		case <-sc:
			srv.Shutdown(context.TODO())
		}
	}()

	return srv.ListenAndServe()
}

func Command() *cli.Command {
	cmd := &cli.Command{
		Name:  "server",
		Usage: "A simple HTTP static file server",
		Flags: []cli.Flag{
			flagDir,
			flagPort,
		},
		Action: actionHandler,
	}
	return cmd
}

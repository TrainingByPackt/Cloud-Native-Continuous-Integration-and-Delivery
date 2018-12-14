package main

import (
	"flag"
	"os"

	"gitlab.com/onuryilmaz/book-server/pkg/books"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/denisenkom/go-mssqldb"

	"os/signal"
	"sync"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
	"gitlab.com/onuryilmaz/book-server/pkg/commons"
	"gitlab.com/onuryilmaz/book-server/pkg/server"
)

var options commons.Options

var version = "latest"

func init() {
	pflag.StringVar(&options.ServerPort, "port", "8080", "Server port for listening REST calls")
	pflag.StringVar(&options.DatabaseAddress, "db", "", "Database instance")
	pflag.StringVar(&options.LogLevel, "log-level", "info", "Log level, options are panic, fatal, error, warning, info and debug")
}

func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	level, err := logrus.ParseLevel(options.LogLevel)
	if err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Fatal("Error during log level parse:", err)
	}

	sigs := make(chan os.Signal, 1)
	stop := make(chan struct{})
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	wg := &sync.WaitGroup{}

	if options.DatabaseAddress == "" {
		databaseEnvironment := os.Getenv("DATABASE")
		if databaseEnvironment == "" {
			logrus.Warn("Database address is empty, some functionality may not work..")
		} else {
			options.DatabaseAddress = databaseEnvironment
		}

	}

	if options.DatabaseAddress != "" {
		options.Books, err = books.NewSQLBookDatabase(options.DatabaseAddress)
		if err != nil {
			logrus.Fatal("Error during database creation:", err)
		}
	}

	options.Version = version

	webserver := server.NewREST(options)
	webserver.Start()

	<-sigs
	logrus.Warn("Shutting down...")
	webserver.Stop()

	close(stop)
	wg.Wait()
}

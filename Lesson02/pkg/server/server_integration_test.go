// +build integration

package server

import (
	"flag"
	"testing"

	"net/http"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/denisenkom/go-mssqldb"

	"fmt"

	"gitlab.com/onuryilmaz/book-server/pkg/books"

	"github.com/phayes/freeport"
	"gitlab.com/onuryilmaz/book-server/pkg/commons"
)

var databaseAddress string

func init() {
	flag.StringVar(&databaseAddress, "db", "", "Database instance")
}

func TestBookServerIntegration(t *testing.T) {

	if databaseAddress == "" {
		t.Skip()
	}

	options := commons.Options{}
	options.ServerPort = fmt.Sprintf("%v", freeport.GetPort())

	db, err := books.NewSQLBookDatabase(databaseAddress)
	options.Books = db

	RESTServer := NewREST(options)
	RESTServer.Start()

	Convey("Start and check RESTServer", t, func() {

		So(err, ShouldBeNil)
		So(RESTServer, ShouldNotBeNil)

		// Wait for server is up
		time.Sleep(time.Second)

		Convey("Ping server", func() {

			response, err := http.Get("http://localhost:" + options.ServerPort + "/ping")
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)

		})

		Convey("Init books database", func() {

			response, err := http.Get("http://localhost:" + options.ServerPort + "/v1/init")
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)

		})

		Convey("Get books", func() {

			response, err := http.Get("http://localhost:" + options.ServerPort + "/v1/books")
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)

		})

	})

}

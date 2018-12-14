// +build smoke

package server

import (
	"flag"
	"fmt"
	"testing"

	"gitlab.com/onuryilmaz/book-server/pkg/books"

	"net/http"
	"time"

	"github.com/phayes/freeport"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/onuryilmaz/book-server/pkg/commons"

	_ "github.com/lib/pq"
)

var databaseAddress string

func init() {
	flag.StringVar(&databaseAddress, "db", "", "Database instance")
}

func TestBookServerSmoke(t *testing.T) {

	if databaseAddress == "" {
		t.Skip()
	}

	options := commons.Options{}
	options.ServerPort = fmt.Sprintf("%v", freeport.GetPort())

	db, err := books.NewSQLBookDatabase(databaseAddress)
	options.Books = db

	RESTServer := NewREST(options)
	RESTServer.Start()

	Convey("Start book server with database", t, func() {
		So(err, ShouldBeNil)
		So(RESTServer, ShouldNotBeNil)

		// Wait for server is up
		time.Sleep(time.Second)

		Convey("Ping server", func() {

			response, err := http.Get("http://localhost:" + options.ServerPort + "/ping")
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)

		})

	})
}

// +build unit

package server

import (
	"fmt"
	"testing"

	"gitlab.com/onuryilmaz/book-server/pkg/books"

	"net/http"
	"time"

	"github.com/phayes/freeport"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/onuryilmaz/book-server/pkg/commons"
)

func TestBookServer(t *testing.T) {

	options := commons.Options{}
	options.ServerPort = fmt.Sprintf("%v", freeport.GetPort())
	options.Books = MockBookDatabase{}
	RESTServer := NewREST(options)
	RESTServer.Start()

	Convey("Start book server with mocked database", t, func() {
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

type MockBookDatabase struct{}

func (mbd MockBookDatabase) GetBooks() ([]books.Book, error) {

	b := make([]books.Book, 0)

	b = append(b, books.Book{ISBN: "ISBN", Title: "Title", Author: "Author"})

	return b, nil
}

func (mbd MockBookDatabase) Initialize() error {

	return nil

}

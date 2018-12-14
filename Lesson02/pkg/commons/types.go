// Package commons provides data structures used between packages
package commons

import "gitlab.com/onuryilmaz/book-server/pkg/books"

// Options provides overall configuration data
type Options struct {
	ServerPort string
	LogLevel   string

	Books           books.BookDatabase
	DatabaseAddress string

	Version string
}

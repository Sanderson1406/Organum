package server

import (
	"net/http"

	handler "organum/Handler"
)

type Route struct {
	Method       string
	Path         string
	Handler      http.HandlerFunc
	RequiresAuth bool
}

var Routes = map[string]Route{
	"Login": {
		Method:       "POST",
		Path:         "/login",
		Handler:      handler.Login,
		RequiresAuth: false,
	},
	"CreateUser": {
		Method:       "POST",
		Path:         "/createuser",
		Handler:      handler.CreateUser,
		RequiresAuth: false,
	},
	"Logout": {
		Method:       "POST",
		Path:         "/logout",
		Handler:      handler.Logout,
		RequiresAuth: true,
	},
	"CreateAndAddNewBook": {
		Method:       "POST",
		Path:         "/lectio/book",
		Handler:      handler.CreateAndAddNewBook,
		RequiresAuth: true,
	},
	"ListBooksByStatus": {
		Method:       "GET",
		Path:         "/lectio/book/l",
		Handler:      handler.ListBooksByStatus,
		RequiresAuth: true,
	},
	"UpdateBookTag": {
		Method:       "POST",
		Path:         "/lectio/book/updattag",
		Handler:      handler.UpdateBookTag,
		RequiresAuth: true,
	},
	"CreateReview": {
		Method:       "POST",
		Path:         "/lectio/review",
		Handler:      handler.CreateReview,
		RequiresAuth: true,
	},
	"GetReview": {
		Method:       "GET",
		Path:         "/lectio/review",
		Handler:      handler.GetReview,
		RequiresAuth: true,
	},
	"UpdateBookProgress": {
		Method:       "PUT",
		Path:         "/lectio/book/progress",
		Handler:      handler.UpdateBookProgress,
		RequiresAuth: true,
	},
	"GetBookProgress": {
		Method:       "GET",
		Path:         "/lectio/book/progress",
		Handler:      handler.GetBookProgress,
		RequiresAuth: true,
	},
	"CreateQuote": {
		Method:       "POST",
		Path:         "/lectio/quote",
		Handler:      handler.CreateQuote,
		RequiresAuth: true,
	},
	"GetQuotes": {
		Method:       "GET",
		Path:         "/lectio/quote",
		Handler:      handler.GetQuotes,
		RequiresAuth: true,
	},
	"DeleteBook": {
		Method:       "DELETE",
		Path:         "/lectio/book",
		Handler:      handler.DeleteBook,
		RequiresAuth: true,
	},
}

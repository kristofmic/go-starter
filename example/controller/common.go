package controller

import (
	"context"
	"net/http"
)

// Handler implements our context aware HTTP handlers.
type Handler interface {
	ServeHTTP(context.Context, http.ResponseWriter, *http.Request) error
}

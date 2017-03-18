package controller

import (
	"context"
	"io"
	"net/http"
)

// RootHandler adapts the standard lib's http.Handler to our context aware handler interface and chains in a standard set of middleware.
type RootHandler struct {
	Next Handler
}

func (rh *RootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	if rh.Next != nil {
		if err := rh.Next.ServeHTTP(ctx, w, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	io.WriteString(w, "You're drunk go home")

}

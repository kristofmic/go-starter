package controller

import (
	"context"
	"io"
	"net/http"
)

// GetData adapts the standard lib's http.Handler to our context aware handler interface and chains in a standard set of middleware.
type GetData struct {
	Next Handler
}

func (gd *GetData) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	if gd.Next != nil {
		if err := gd.Next.ServeHTTP(ctx, w, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	io.WriteString(w, `{"foo":"bar"}`)

}

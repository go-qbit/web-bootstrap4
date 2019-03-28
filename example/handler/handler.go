package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/go-qbit/web-bootstrap4/template"
	"github.com/go-qbit/web/form"
	"github.com/go-qbit/web/handler"
)

const ctxKeyHttpR  = "HttpR"

type RootHandler struct {
	*handler.Handler
}

func New() *RootHandler {
	if err := form.InitAntiCSRF(&form.AntiCSRFConfig{
		Salt:      "VerySecretString",
		Length:    10,
		ErrorText: "AntiCSRF key is invalid, try to refresh page",
		CtxParser: func(ctx context.Context) (userID string, formPath string) {
			return "", ctx.Value(ctxKeyHttpR).(*http.Request).URL.Path
		},
	}); err != nil {
		log.Fatal(err)
	}

	h := &RootHandler{
		&handler.Handler{},
	}

	h.HandleForm("test_form", &testForm{}, "Form", nil)

	return h
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := template.OptionsToContext(r.Context(), &template.PageOptions{
		RootHandler: h,
		Title:       "Test Bootstrap 4",
		NavBar: template.PONavBar{
			Brand: "Test Bootstrap 4",
		},
	})

	ctx = context.WithValue(ctx, ctxKeyHttpR, r)

	h.ServeSubHandlers(w, r.WithContext(ctx))
}

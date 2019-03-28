package handler

import (
	"context"
	"io"
	"net/http"

	"github.com/go-qbit/qerror"
	"github.com/go-qbit/web-bootstrap4/bsform"
	"github.com/go-qbit/web-bootstrap4/bsform/bsfield"
	"github.com/go-qbit/web/form"
	"github.com/go-qbit/web/form/field"
)

type testForm struct {
}

func (fi *testForm) GetFields(ctx context.Context) []field.IField {
	inputWithError := &bsfield.Input{
		Name:  "test-input-error",
		Type:  "text",
		Label: "Text input with error",
	}
	inputWithError.SetError(qerror.PublicErrorf("Test error"))

	return []field.IField{
		&bsfield.Input{
			Name:        "test-input",
			Type:        "text",
			Label:       "Text input",
			Placeholder: "placeholder",
		},

		&bsfield.Input{
			Name:     "required-test-input",
			Type:     "text",
			Label:    "Required text input",
			Required: true,
		},

		&bsfield.Input{
			Name:    "test-input-default",
			Type:    "text",
			Label:   "Text input with default value",
			Default: "Default value",
		},

		&bsfield.Input{
			Name:  "test-input-password",
			Type:  "password",
			Label: "Password",
		},

		inputWithError,

		&bsfield.Textarea{
			Name:  "textarea",
			Label: "Textarea",
			Rows:  10,
			Cols:  50,
		},
	}
}

func (*testForm) GetCaption(ctx context.Context) string {
	return "Test form"
}
func (*testForm) GetSubmitCaption(ctx context.Context) string {
	return "Submit"
}

func (*testForm) RenderHTML(ctx context.Context, w io.Writer, f *form.Form) {
	bsform.ProcessPageForm(ctx, w, f)
}

func (*testForm) OnSave(ctx context.Context, w http.ResponseWriter, f *form.Form) qerror.PublicError {
	return nil
}

func (*testForm) OnComplete(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Submitted"))
}

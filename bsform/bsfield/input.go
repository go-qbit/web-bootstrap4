//go:generate ttgen -package bsfield input.gtt
package bsfield

import (
	"context"
	"io"
	"net/url"

	"github.com/go-qbit/qerror"
	"github.com/go-qbit/web/form/field"
)

var _ field.IField = &Input{}

type Input struct {
	Name        string
	Label       string
	Type        string
	Required    bool
	Default     string
	Placeholder string
	Class       []string
	CheckFunc   func(ctx context.Context, value string) qerror.PublicError
	value       *string
	err         qerror.PublicError
}

func (f *Input) Init(ctx context.Context, form url.Values) {
	if v := form[f.Name]; v != nil {
		f.value = &v[0]
	}
}

func (f *Input) GetName() string {
	return f.Name
}

func (f *Input) GetValue() interface{} {
	return f.GetStringValue()
}

func (f *Input) GetStringValue() string {
	if f.value == nil {
		return f.Default
	}

	return *f.value
}

func (f *Input) Check(ctx context.Context) qerror.PublicError {
	if f.GetStringValue() == "" {
		return field.ErrMissedReqField(ctx)
	}

	if f.CheckFunc != nil {
		if err := f.CheckFunc(ctx, f.GetStringValue()); err != nil {
			return err
		}
	}

	return nil
}

func (f *Input) ProcessHtml(ctx context.Context, w io.Writer) {
	ProcessInput(ctx, w, f)
}

func (f *Input) SetError(err qerror.PublicError) qerror.PublicError {
	f.err = err

	return err
}

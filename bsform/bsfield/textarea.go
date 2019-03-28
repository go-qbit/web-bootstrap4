//go:generate ttgen -package bsfield textarea.gtt
package bsfield

import (
	"context"
	"io"
	"net/url"

	"github.com/go-qbit/qerror"
	"github.com/go-qbit/web/form/field"
)

var _ field.IField = &Textarea{}

type Textarea struct {
	Name      string
	Label     string
	Required  bool
	Default   string
	Rows      uint16
	Cols      uint16
	Class     []string
	CheckFunc func(ctx context.Context, value string) qerror.PublicError
	value     *string
	err       qerror.PublicError
}

func (f *Textarea) Init(ctx context.Context, form url.Values) {
	if v := form[f.Name]; v != nil {
		f.value = &v[0]
	}
}

func (f *Textarea) GetName() string {
	return f.Name
}

func (f *Textarea) GetValue() interface{} {
	return f.GetStringValue()
}

func (f *Textarea) GetStringValue() string {
	if f.value == nil {
		return f.Default
	}

	return *f.value
}

func (f *Textarea) Check(ctx context.Context) qerror.PublicError {
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

func (f *Textarea) ProcessHtml(ctx context.Context, w io.Writer) {
	ProcessTextarea(ctx, w, f)
}

func (f *Textarea) SetError(err qerror.PublicError) qerror.PublicError {
	f.err = err

	return err
}

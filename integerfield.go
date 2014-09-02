package gforms

import (
	"errors"
	"reflect"
	"strconv"
)

type IntegerField struct {
	BaseField
	ErrorMessage string
}

func (f *IntegerField) New() FieldInterface {
	fi := new(IntegerFieldInstance)
	if f.ErrorMessage == "" {
		fi.ErrorMessage = "This field should be specified as int."
	} else {
		fi.ErrorMessage = f.ErrorMessage
	}
	fi.Model = f
	fi.V = nilV("")
	return fi
}

type IntegerFieldInstance struct {
	FieldInstance
	ErrorMessage string
}

func NewIntegerField(name string, vs Validators, ws ...Widget) *IntegerField {
	f := new(IntegerField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	}
	return f
}

func (f *IntegerFieldInstance) Clean(data Data) error {
	m, hasField := data[f.Model.GetName()]
	if hasField {
		f.V = m
		v := m.rawValueAsString()
		m.Kind = reflect.Int
		if v != nil && (*v) != "" {
			iv, err := strconv.Atoi(*v)
			if err == nil {
				m.Value = iv
				m.IsNil = false
				return nil
			}
			return errors.New(f.ErrorMessage)
		}
	}
	return nil
}

func (f *IntegerFieldInstance) html() string {
	return renderTemplate("TextTypeField", newTemplateContext(f))
}

func (f *IntegerFieldInstance) Html() string {
	return fieldToHtml(f)
}

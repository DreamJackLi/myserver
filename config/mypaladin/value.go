package mypaladin

import (
	"encoding"
	"errors"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	ErrNotExist       = errors.New("paladin: value key not exist")
	ErrTypeAssertion  = errors.New("paladin: value type assertion no match")
	ErrDifferentTypes = errors.New("paladin: value different types")
)

type Value struct {
	val   interface{}
	slice interface{}
	raw   string
}

func NewValue(val interface{}, raw string) *Value {
	return &Value{
		val: val,
		raw: raw,
	}
}

func (v *Value) Raw() (string, error) {
	if v.val == nil {
		return "", ErrNotExist
	}

	return v.raw, nil
}

func (v *Value) Bool() (bool, error) {
	if v.val == nil {
		return false, ErrNotExist
	}

	b, ok := v.val.(bool)
	if !ok {
		return false, ErrTypeAssertion
	}
	return b, nil
}

func (v *Value) Int() (int, error) {
	i, err := v.Int64()
	return int(i), err
}

func (v *Value) Int32() (int32, error) {
	i, err := v.Int64()

	return int32(i), err
}

func (v *Value) Int64() (int64, error) {
	if v.val == nil {
		return 0, ErrNotExist
	}

	i, ok := v.val.(int64)

	if !ok {
		return 0, ErrTypeAssertion
	}

	return i, nil
}

func (v *Value) Float32() (float32, error) {
	f, err := v.Float64()

	return float32(f), err
}

func (v *Value) Float64() (float64, error) {
	if v.val == nil {
		return 0.0, ErrNotExist
	}

	f, ok := v.val.(float64)

	if !ok {
		return 0.0, ErrTypeAssertion
	}

	return f, nil
}

func (v *Value) String() (string, error) {
	if v.val == nil {
		return "", ErrNotExist
	}

	s, ok := v.val.(string)

	if !ok {
		return "", ErrTypeAssertion
	}

	return s, nil
}

func (v *Value) Duration() (time.Duration, error) {
	s, err := v.String()

	if err != nil {
		return time.Duration(0), err
	}
	return time.ParseDuration(s)
}

func (v *Value) Unmarshal(un encoding.TextUnmarshaler) error {
	text, err := v.Raw()

	if err != nil {
		return err
	}

	return un.UnmarshalText([]byte(text))
}

func (v *Value) UnmarshalTOML(dst interface{}) error {
	text, err := v.Raw()

	if err != nil {
		return err
	}

	return toml.Unmarshal([]byte(text), dst)
}

package validation

import (
	"errors"
	"reflect"
	"strconv"
)

func IsString(v interface{}) (string, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.String {
		return "", errors.New(notStringValueError)
	}

	return reflect.ValueOf(v).String(), nil
}

func IsInt(v interface{}) (int64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Int64 && t.Kind() != reflect.Int {
		return 0, errors.New(notIntegerValueError)
	}

	return reflect.ValueOf(v).Int(), nil
}

func IsStringInt(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New(notIntegerValueError)
	}

	return i, nil
}

func IsFloat(v interface{}) (float64, error) {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Float64 {
		return 0.0, errors.New(notFloatValueError)
	}

	return reflect.ValueOf(v).Float(), nil
}

func IsStringFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, errors.New(notFloatValueError)
	}

	return f, nil
}

func IsNumber(v interface{}) error {
	_, isF := IsFloat(v)
	if isF == nil {
		return nil
	}

	_, isI := IsInt(v)
	if isI == nil {
		return nil
	}

	return errors.New(notNumberValueError)
}

func IsStringNumber(v interface{}) error {
	s, err := IsString(v)
	if err != nil {
		return errors.New(notStringNumberValueError)
	}

	_, isSF := IsStringFloat(s)
	if isSF == nil {
		return nil
	}

	_, isSI := IsStringInt(s)
	if isSI != nil {
		return nil
	}

	return errors.New(notStringNumberValueError)
}

func IsNil(v interface{}) error {
	if v != nil {
		return errors.New(notNilValueError)
	}

	return nil
}

func IsEmpty(v string) bool {
	return v == Empty
}

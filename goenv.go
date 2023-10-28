package goenv

import (
	"os"
	"reflect"
	"strconv"
)

func GetEnv[T any](name string, d T) T {
	v := os.Getenv(name)
	if v == "" {
		return d
	}

	var refVal reflect.Value

	t := reflect.TypeOf(d)
	switch t.Kind() {
	case reflect.Bool:
		refVal = reflect.ValueOf(v == "true")
	case reflect.Int:
		i, err := strconv.Atoi(v)
		if err != nil {
			return d
		}

		refVal = reflect.ValueOf(i)
	case reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return d
		}

		refVal = reflect.ValueOf(f)
	case reflect.String:
		refVal = reflect.ValueOf(v)
	default:
		return d
	}

	tValue := refVal.Convert(t)
	return tValue.Interface().(T)
}

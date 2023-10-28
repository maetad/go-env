package goenv_test

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"

	goenv "github.com/maetad/go-env"
)

func TestGetEnvBool(t *testing.T) {
	os.Setenv("UNIT_TEST_BOOL_TRUE", "true")
	os.Setenv("UNIT_TEST_BOOL_FALSE", "false")

	type args struct {
		name string
		d    bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				name: "UNIT_TEST_BOOL_TRUE",
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				name: "UNIT_TEST_BOOL_FALSE",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goenv.GetEnv[bool](tt.args.name, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvInt(t *testing.T) {
	os.Setenv("UNIT_TEST_PORT", "8080")
	os.Setenv("UNIT_TEST_DB", "db1")

	type args struct {
		name string
		d    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "8080",
			args: args{
				name: "UNIT_TEST_PORT",
			},
			want: 8080,
		},
		{
			name: "fallback",
			args: args{
				name: "UNIT_TEST_DB",
				d:    1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goenv.GetEnv[int](tt.args.name, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvFloat64(t *testing.T) {
	os.Setenv("UNIT_TEST_THRESHOLD", "0.2")
	os.Setenv("UNIT_TEST_SPEED", fmt.Sprintf("1%f", math.MaxFloat64))

	type args struct {
		name string
		d    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0.2",
			args: args{
				name: "UNIT_TEST_THRESHOLD",
			},
			want: 0.2,
		},
		{
			name: "fallback",
			args: args{
				name: "UNIT_TEST_SPEED",
				d:    0.1,
			},
			want: 0.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goenv.GetEnv[float64](tt.args.name, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvString(t *testing.T) {
	os.Setenv("UNIT_TEST_HOST", "localhost")

	type args struct {
		name string
		d    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "localhost",
			args: args{
				name: "UNIT_TEST_HOST",
			},
			want: "localhost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goenv.GetEnv[string](tt.args.name, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvUnsupportyed(t *testing.T) {
	os.Setenv("UNIT_TEST_BYTE", "74 65 73 74")

	type args struct {
		name string
		d    []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test",
			args: args{
				name: "UNIT_TEST_BYTE",
				d:    []byte("test"),
			},
			want: []byte("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goenv.GetEnv[[]byte](tt.args.name, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvDefault(t *testing.T) {
	type args struct {
		name string
		d    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				name: "UNIT_TEST_ANY",
				d:    "any",
			},
			want: "any",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goenv.GetEnv[string](tt.args.name, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

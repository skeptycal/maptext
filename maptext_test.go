package maptext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRoot(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		data     []byte
		// parent   *node
		// left     *node
		// right    *node
		want    *node
		wantErr bool
	}{
		{"root", "bad_name", []byte("failData"), &node{"badName", []byte("failData"), nil, nil, nil}, true},
		{"root", "root", []byte("rootData"), &node{"root", []byte("rootData"), nil, nil, nil}, true},
		{"bad data - faker", "fail", []by te("failData"), &node{"fail", []byte("failData"), nil, nil, nil}, true},
		{"bad parent", "fail", []byte("failData"), &node{"fail", []byte("failData"), nil, nil, nil}, true},
		{"bad left", "fail", []byte("failData"), &node{"fail", []byte("failData"), nil, nil, nil}, true},
		{"bad right", "fail", []byte("failData"), &node{"fail", []byte("failData"), nil, nil, nil}, true},

		// False Negative for testing
		// {"false negative", "correct", []byte("failData"), &node{"correct", []byte("failData"), nil, nil, nil}, false},
		// False Positive for testing
		// {"false positive", "incorrect", []byte("rootData"), &node{"root", []byte("rootData"), nil, nil, nil}, true},
	}
	for _, tt := range tests {
		TRun(t, tt.testName, tt.name, NewRoot(tt.name, tt.data), tt.want, tt.wantErr)
	}
}

func TRun(t *testing.T, testName, argName string, got, want Any, wantErr bool) {
	t.Run(fmt.Sprintf("%s(%s)", testName, argName), func(t *testing.T) {
		if !reflect.DeepEqual(got, want) == !wantErr {
			t.Errorf("NewRoot() = %v, want %v", got, want)
		}
	})
}

package maptext

import (
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
		want *node
	}{
		// TODO: Add test cases.
		{"root", "root", []byte("rootData"), &node{"root", []byte("rootData"), nil, nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := NewRoot(tt.name, tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

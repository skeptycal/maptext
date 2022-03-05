package maptext

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	fakeSingle    = NewRoot("fakeSingle", []byte("fakeSingle"))
	fakeRoot      = NewRoot("fake", []byte("fake"))
	rightNode     = NewNode("right", []byte("right"), fakeRoot, fakeRoot)
	leftNode      = NewNode("left", []byte("left"), fakeRoot, fakeRoot)
	leftLeftNode  = NewNode("leftLeft", []byte("leftLeft"), fakeRoot, leftNode)
	rightLeftNode = NewNode("rightLeft", []byte("rightLeft"), fakeRoot, leftNode)

	NodeTests = []struct {
		testName string
		name     string
		data     []byte
		// parent   *node
		// left     *node
		// right    *node
		want    *node
		wantErr bool
	}{
		{"fakeSingle", "fakeSingle", []byte("fakeSingle"), &node{"fakeSingle", []byte("fakeSingle"), [16]byte{}, nil, nil, nil, nil}, false},
		{"fakeRoot", "fakeRoot", []byte("fakeRoot"), &node{"fakeSingle", []byte("fakeSingle"), [16]byte{}, nil, nil, nil, nil}, false},

		{"root", "root", []byte("rootData"), &node{"root", []byte("rootData"), [16]byte{}, nil, nil, nil, nil}, false},
		{"root", "badName", []byte("failData"), &node{"bad_Name", []byte("failData"), [16]byte{}, nil, nil, nil, nil}, true},
		{"bad data", "fail", []byte("same"), &node{"fail", []byte("different"), [16]byte{}, nil, nil, nil, nil}, true},
		{"bad root", "badRoot", []byte("same"), &node{"badRoot", []byte("same"), [16]byte{}, fakeRoot, nil, nil, nil}, true},
		{"bad parent", "badParent", []byte("same"), &node{"badParent", []byte("same"), [16]byte{}, nil, fakeRoot, nil, nil}, true},
		{"bad left", "badLeft", []byte("same"), &node{"badLeft", []byte("same"), [16]byte{}, nil, nil, fakeRoot, nil}, true},
		{"bad right", "badRight", []byte("same"), &node{"badRight", []byte("same"), [16]byte{}, nil, nil, nil, fakeRoot}, true},

		// False Negative and False Positive for testing
		{"false negative", "False Negative", []byte("false negative"), &node{"False Negative", []byte("false negative"), [16]byte{}, nil, &node{}, nil, nil}, false},
		{"false positive", "False Positive", []byte("false positive"), &node{"False Positive", []byte("false positive"), [16]byte{}, nil, nil, nil, nil}, true},
	}
)

func init() {
	fakeRoot.AddRight(rightNode)
	fakeRoot.AddLeft(leftNode)
	leftNode.AddRight(rightLeftNode)
	leftNode.AddLeft(leftLeftNode)
}

func TestNewRoot(t *testing.T) {
	testName := "NewRoot"
	for _, tt := range NodeTests {
		got := NewRoot(tt.name, tt.data)
		TRun(t, testName, tt.testName, got, tt.want, tt.wantErr)
		got.AddLeft(fakeRoot)
		TRun(t, "Len", tt.testName, got.Len(), 2, tt.wantErr)
		got.AddRight(fakeRoot)
		TRun(t, "Len", tt.testName, got.Len(), 3, tt.wantErr)
	}
}

func TRun(t *testing.T, testName, argName string, got, want Any, wantErr bool) {
	t.Run(fmt.Sprintf("%s(%s)", testName, argName), func(t *testing.T) {
		if !reflect.DeepEqual(got, want) && !wantErr {
			t.Errorf("%s(%s) = %v, want %+v", testName, argName, got, want)
		}
	})
}

func Test_id_String(t *testing.T) {
	tests := []struct {
		name string
		i    *id
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("id.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_edge_CanLeft(t *testing.T) {
	tests := []struct {
		name string
		e    *edge
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.CanLeft(); got != tt.want {
				t.Errorf("edge.CanLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_edge_CanRight(t *testing.T) {
	tests := []struct {
		name string
		e    *edge
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.CanRight(); got != tt.want {
				t.Errorf("edge.CanRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNodeMap(t *testing.T) {
	type args struct {
		name   string
		parent *node
	}
	tests := []struct {
		name string
		args args
		want *nodeMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNodeMap(tt.args.name, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		name   string
		data   []byte
		root   *node
		parent *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.name, tt.args.data, tt.args.root, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_String(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Len(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Len(); got != tt.want {
				t.Errorf("node.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		n    *node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_node_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		n    *node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("node.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Equal(t *testing.T) {
	type args struct {
		other *node
	}
	tests := []struct {
		name string
		n    *node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Equal(tt.args.other); got != tt.want {
				t.Errorf("node.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_CheckSum(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want [16]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.CheckSum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.CheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_AddLeft(t *testing.T) {
	type args struct {
		other *node
	}
	tests := []struct {
		name string
		n    *node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.AddLeft(tt.args.other)
		})
	}
}

func Test_node_AddRight(t *testing.T) {
	type args struct {
		other *node
	}
	tests := []struct {
		name string
		n    *node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.AddRight(tt.args.other)
		})
	}
}

func Test_node_Delete(t *testing.T) {
	tests := []struct {
		name string
		n    *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Delete()
		})
	}
}

func TestCheckSum(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want [16]byte
	}{
		// TODO: Add test cases.
		{"checksum test", args{data: []byte("data")}, [16]byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSum(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

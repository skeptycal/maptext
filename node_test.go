package maptext

import (
	"fmt"
	"testing"

	"github.com/skeptycal/types"
)

var (
	fakeData      = []byte("fakeData")
	emptyChecksum = [16]byte{}

	fakeSingle    = NewRoot("fakeSingle", fakeData)
	fakeRoot      = NewRoot("fake", fakeData)
	rightNode     = NewNode("right", fakeData, fakeRoot, fakeRoot)
	leftNode      = NewNode("left", fakeData, fakeRoot, fakeRoot)
	leftLeftNode  = NewNode("leftLeft", fakeData, fakeRoot, leftNode)
	rightLeftNode = NewNode("rightLeft", fakeData, fakeRoot, leftNode)

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
		{"fakeSingle", "fakeSingle", fakeData, &node{"fakeSingle", fakeData, emptyChecksum, nil, nil, nil, nil}, false},
		{"fakeRoot", "fakeRoot", fakeData, &node{"fakeSingle", fakeData, emptyChecksum, nil, nil, nil, nil}, false},

		{"root", "root", fakeData, &node{"root", fakeData, emptyChecksum, nil, nil, nil, nil}, false},

		{"root", "badName", fakeData, &node{"bad_Name", fakeData, emptyChecksum, nil, nil, nil, nil}, true},
		{"bad data", "fail", fakeData, &node{"fail", fakeData, emptyChecksum, nil, nil, nil, nil}, true},
		{"bad root", "badRoot", fakeData, &node{"badRoot", fakeData, emptyChecksum, fakeRoot, nil, nil, nil}, true},
		{"bad parent", "badParent", fakeData, &node{"badParent", fakeData, emptyChecksum, nil, fakeRoot, nil, nil}, true},
		{"bad left", "badLeft", fakeData, &node{"badLeft", fakeData, emptyChecksum, nil, nil, fakeRoot, nil}, true},
		{"bad right", "badRight", fakeData, &node{"badRight", fakeData, emptyChecksum, nil, nil, nil, fakeRoot}, true},

		// False Negative and False Positive for testing
		{"false negative", "False Negative", fakeData, &node{"False Negative", fakeData, emptyChecksum, nil, nil, nil, nil}, true},
		{"false positive", "False Positive", fakeData, &node{"False Positive", fakeData, emptyChecksum, nil, nil, nil, nil}, true},
	}
)

func init() {
	fakeRoot.AddRight(rightNode)
	fakeRoot.AddLeft(leftNode)
	leftNode.AddRight(rightLeftNode)
	leftNode.AddLeft(leftLeftNode)
}

func ExampleNewTree() {
	tree := makeTree("test")
	fmt.Println(tree)
	// output:
	// Tree
}

func TestNewRoot(t *testing.T) {
	testName := "NewRoot"
	for _, tt := range NodeTests {
		got := NewRoot(tt.name, tt.data)
		TRun(t, testName, tt.testName, got, tt.want, tt.wantErr)
		got.AddLeft(fakeRoot)
		TRun(t, "Len", tt.testName, got.Len(), len(got.data), tt.wantErr)
		got.AddRight(fakeRoot)
		TRun(t, "Len", tt.testName, got.Len(), len(got.data), tt.wantErr)
	}
}

func TRun(t *testing.T, testName, argName string, got, want Any, wantErr bool) {
	t.Run(fmt.Sprintf("%s(%s)", testName, argName), func(t *testing.T) {

		g := types.NewAnyValue(got)
		w := types.NewAnyValue(want)
		if g != w { // !reflect.DeepEqual(got, want)
			if !wantErr {
				t.Errorf("%s(%s)(wantErr: %v) = %v(%T), want %+v(%T)", testName, argName, wantErr, got, got, want, want)
			}
		}
	})
}

package ref

import (
	"reflect"
	"testing"

	"github.com/exfly/cslab/Code/Go/lngo/model"
)

type Assert struct {
	*testing.T
}

func (a Assert) Assert(want, get interface{}, errMsg string) {
	if !reflect.DeepEqual(want, get) {
		a.Errorf("\nerrMsg: %v get %v, want %v", errMsg, get, want)
	}
}
func Test_testReflect(t *testing.T) {
	t.Error(reflectValString(model.People{}))
	t.Error(reflectValString(&model.People{}))
}

func Test_reflectValName(t *testing.T) {
	t.Error(reflectValName(model.People{}))
}

func Test_reflectValFullName(t *testing.T) {
	t.Error(reflectValFullName(model.People{}))
	t.Error(reflectValFullName(model.People{}.Account))
	t.Error(reflectValFullName(&model.People{}))
}

func Test_FullPath(t *testing.T) {
	t.Error(FullPath(model.People{}))
	t.Error(FullPath(model.People{}.Account))
	t.Error(FullPath(&model.People{}))
}

func Test_FullPathV2(t *testing.T) {
	assert := Assert{T: t}
	assert.Assert("github.com/exfly/cslab/Code/Go/lngo/model.People", FullPath(model.People{}), "people fullpath")
}

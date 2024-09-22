package syntaxgo_reflect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

type Example struct{}

func (x *Example) methodFunction() {}

func TestGetPkgPath(t *testing.T) {
	t.Log(GetPkgPath(Example{}))
	t.Log(GetPkgName(Example{}))

	t.Log(GetPkgPath(&Example{}))
	t.Log(GetPkgName(&Example{}))

	a := &Example{}
	t.Log(GetPkgPath(a.methodFunction))
	t.Log(GetPkgName(a.methodFunction))

	b := Example{}
	t.Log(GetPkgPath(b.methodFunction))
	t.Log(GetPkgName(b.methodFunction))
}

func commonFunction() {}

func TestGetPkgPath1(t *testing.T) {
	t.Log(GetPkgPath(commonFunction))
	t.Log(GetPkgName(commonFunction))
}

type ExampleInterface interface {
	methodFunction()
}

type ExampleOneOne struct{}

func (e *ExampleOneOne) methodFunction() {}

func TestGetPkgPath2(t *testing.T) {
	var a ExampleInterface = &ExampleOneOne{}
	t.Log(GetPkgPath(a))
	t.Log(GetPkgName(a))
}

type ExampleTwoTwo struct{}

func (e ExampleTwoTwo) methodFunction() {}

func TestGetPkgPath3(t *testing.T) {
	var a ExampleInterface = ExampleTwoTwo{}
	t.Log(GetPkgPath(a))
	t.Log(GetPkgName(a))
}

func TestGetTypeName(t *testing.T) {
	typeName := GetTypeName(Example{})
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

func TestGetTypeNameV2(t *testing.T) {
	typeName := GetTypeNameV2[Example]()
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

func TestGetTypeNameV3(t *testing.T) {
	typeName := GetTypeNameV3(&Example{})
	t.Log(typeName)
	require.Equal(t, "Example", typeName)
}

func TestGetPkgPathV2(t *testing.T) {
	t.Log(GetPkgPathV2[Example]())
	t.Log(GetPkgNameV2[Example]())

	t.Log(GetPkgPathV2[*Example]()) //这个是不行的，目前不支持往里面传指针类型，但并不会panic而是返回空白
	t.Log(GetPkgNameV2[*Example]()) //这个是不行的，目前不支持往里面传指针类型，但并不会panic而是返回空白
}

func TestGetPkgNameV3(t *testing.T) {
	t.Log(GetPkgPathV3(Example{}))
	t.Log(GetPkgNameV3(Example{}))

	t.Log(GetPkgPathV3(&Example{}))
	t.Log(GetPkgNameV3(&Example{}))
}

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

func TestGetTypeUsageCode(t *testing.T) {
	usageCode := GetTypeUsageCode(Example{})
	t.Log(usageCode)
	require.Equal(t, "syntaxgo_reflect.Example", usageCode)
}

func TestGetTypeUsageCodeV2(t *testing.T) {
	usageCode := GetTypeUsageCodeV2[Example]()
	t.Log(usageCode)
	require.Equal(t, "syntaxgo_reflect.Example", usageCode)
}

func TestGetPkgPathV2(t *testing.T) {
	t.Log(GetPkgPathV2[Example]())
	t.Log(GetPkgNameV2[Example]())
}

func TestGetPkgPaths(t *testing.T) {
	objectsTypes := GetObjectsTypes([]any{Example{}, ExampleOneOne{}, ExampleTwoTwo{}})
	pkgPaths := GetPkgPaths(objectsTypes)
	t.Log(pkgPaths)
}

func TestGetPkgPaths4Imports(t *testing.T) {
	objectsTypes := GetObjectsTypes([]any{Example{}, ExampleOneOne{}, ExampleTwoTwo{}})
	pkgPaths := GetPkgPaths4Imports(objectsTypes)
	t.Log(pkgPaths)
}

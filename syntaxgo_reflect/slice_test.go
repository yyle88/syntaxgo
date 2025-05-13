package syntaxgo_reflect

import "testing"

func TestGetPkgPaths(t *testing.T) {
	objectsTypes := GetTypes([]any{Example{}, ExampleOneOne{}, ExampleTwoTwo{}})
	pkgPaths := GetPkgPaths(objectsTypes)
	t.Log(pkgPaths)
}

func TestGetPkgPathsFromPointers(t *testing.T) {
	objectsTypes := GetTypes([]any{&Example{}, &ExampleOneOne{}, &ExampleTwoTwo{}})
	pkgPaths := GetPkgPaths(objectsTypes)
	t.Log(pkgPaths)
}

func TestGetPkgPathsToImportWithQuotes(t *testing.T) {
	objectsTypes := GetTypes([]any{Example{}, ExampleOneOne{}, ExampleTwoTwo{}})
	pkgPaths := GetQuotedPackageImportPaths(objectsTypes)
	t.Log(pkgPaths)
}

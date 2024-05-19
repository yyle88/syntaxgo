package syntaxgo_test

import (
	"testing"

	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
)

func TestCurrentPackageName(t *testing.T) {
	t.Log(syntaxgo.CurrentPackageName())
}

func TestGetPkgName(t *testing.T) {
	t.Log(syntaxgo.GetPkgName(runtestpath.SrcName(t)))
}

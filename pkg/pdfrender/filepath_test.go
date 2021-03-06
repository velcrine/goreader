package pdfrender

import (
	"encoding/json"
	"testing"
)

func TestMod_GetFilePath(t *testing.T) {
	objectOptionsBytes := json.RawMessage(`{ "userStylesheetLocation": "css/envoy.css"}`)
	m := Mod{
		// here it is compulsary for scheme(https/http) to be present
		// otherwise net package wouldn't work.
		// hence, no need for exclusive normalization for scheme
		BaseDir:       "envoy-blogs",
		HistPointer:   "i",
		ObjectOptions: &objectOptionsBytes,
		dirVisited:    map[string]*DirVisited{},
	}
	filePath, err := m.GetRawFilePath("https://medium.com/.*/galgodas/@abhbose6/bazel-101-2b0272b15da8/#/whereisid/sdfsd/")
	if err != nil {
		t.Error(err)
	}
	if filePath != "galgodas/@abhbose6" {
		t.Errorf("path %v not expected", filePath)
	}
	//if fileName != "bazel-101-2b0272b15da8" {
	//	t.Errorf("name %v not expected", fileName)
	//}
	dir, index, err := m.GetIndexedDir(filePath)
	if err != nil {
		t.Error(err)
	}
	if dir != "envoy-blogs/i01galgodas/i01-01@abhbose6" {
		t.Error("got path: ", dir)
	}
	if index != "i01-01-01" {
		t.Error("got index: ", index)
	}
}

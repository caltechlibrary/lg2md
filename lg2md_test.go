package lg2md

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	libguide1 := &LibGuides{}

	src, err := ioutil.ReadFile("testdata/example_export.xml")
	if err != nil {
		t.Errorf("%s\n", err)
		t.FailNow()
	}

	err = xml.Unmarshal(src, &libguide1)
	if err != nil {
		t.Errorf("%s\n", err)
		t.FailNow()
	}

	libguide2, err := Decode(src)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("%+v\n%+v\n", libguide2, libguide2)
}

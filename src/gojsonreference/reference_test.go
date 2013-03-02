// author  			sigu-399
// author-github 	https://github.com/sigu-399
// author-mail		sigu.399@gmail.com
// 
// repository-name	gojsonreference
// repository-desc	An implementation of JSON Reference - Go language
// 
// description		Automated tests on package.
// 
// created      	03-03-2013

package gojsonreference

import (
	"testing"
)

func TestFull(t *testing.T) {

	in := "http://host/path/a/b/c#/f/a/b"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("Get(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != false {
		t.Errorf("Get(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, false)
	}

	if r1.HasFullUrl != true {
		t.Errorf("Get(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, true)
	}

	if r1.HasUrlPathOnly != false {
		t.Errorf("Get(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, false)
	}

	if r1.HasFileScheme != false {
		t.Errorf("Get(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestFullUrl(t *testing.T) {

	in := "http://host/path/a/b/c"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("Get(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != false {
		t.Errorf("Get(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, false)
	}

	if r1.HasFullUrl != true {
		t.Errorf("Get(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, true)
	}

	if r1.HasUrlPathOnly != false {
		t.Errorf("Get(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, false)
	}

	if r1.HasFileScheme != false {
		t.Errorf("Get(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestFragmentOnly(t *testing.T) {

	in := "#/fragment/only"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("Get(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != true {
		t.Errorf("Get(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, true)
	}

	if r1.HasFullUrl != false {
		t.Errorf("Get(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, false)
	}

	if r1.HasUrlPathOnly != false {
		t.Errorf("Get(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, false)
	}

	if r1.HasFileScheme != false {
		t.Errorf("Get(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestUrlPathOnly(t *testing.T) {

	in := "/documents/document.json"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("Get(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != false {
		t.Errorf("Get(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, false)
	}

	if r1.HasFullUrl != false {
		t.Errorf("Get(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, false)
	}

	if r1.HasUrlPathOnly != true {
		t.Errorf("Get(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, true)
	}

	if r1.HasFileScheme != false {
		t.Errorf("Get(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

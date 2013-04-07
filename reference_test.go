// Copyright 2013 sigu-399 ( https://github.com/sigu-399 )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		t.Errorf("NewJsonReference(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != false {
		t.Errorf("NewJsonReference(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, false)
	}

	if r1.HasFullUrl != true {
		t.Errorf("NewJsonReference(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, true)
	}

	if r1.HasUrlPathOnly != false {
		t.Errorf("NewJsonReference(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, false)
	}

	if r1.HasFileScheme != false {
		t.Errorf("NewJsonReference(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestFullUrl(t *testing.T) {

	in := "http://host/path/a/b/c"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("NewJsonReference(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != false {
		t.Errorf("NewJsonReference(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, false)
	}

	if r1.HasFullUrl != true {
		t.Errorf("NewJsonReference(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, true)
	}

	if r1.HasUrlPathOnly != false {
		t.Errorf("NewJsonReference(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, false)
	}

	if r1.HasFileScheme != false {
		t.Errorf("NewJsonReference(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestFragmentOnly(t *testing.T) {

	in := "#/fragment/only"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("NewJsonReference(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != true {
		t.Errorf("NewJsonReference(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, true)
	}

	if r1.HasFullUrl != false {
		t.Errorf("NewJsonReference(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, false)
	}

	if r1.HasUrlPathOnly != false {
		t.Errorf("NewJsonReference(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, false)
	}

	if r1.HasFileScheme != false {
		t.Errorf("NewJsonReference(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestUrlPathOnly(t *testing.T) {

	in := "/documents/document.json"

	r1, err := NewJsonReference(in)
	if err != nil {
		t.Errorf("NewJsonReference(%v) error %s", in, err.Error())
	}

	if in != r1.String() {
		t.Errorf("NewJsonReference(%v) = %v, expect %v", in, r1.String(), in)
	}

	if r1.HasFragmentOnly != false {
		t.Errorf("NewJsonReference(%v)::HasFragmentOnly %v expect %v", in, r1.HasFragmentOnly, false)
	}

	if r1.HasFullUrl != false {
		t.Errorf("NewJsonReference(%v)::HasFullUrl %v expect %v", in, r1.HasFullUrl, false)
	}

	if r1.HasUrlPathOnly != true {
		t.Errorf("NewJsonReference(%v)::HasUrlPathOnly %v expect %v", in, r1.HasUrlPathOnly, true)
	}

	if r1.HasFileScheme != false {
		t.Errorf("NewJsonReference(%v)::HasFileScheme %v expect %v", in, r1.HasFileScheme, false)
	}
}

func TestInheritsValid(t *testing.T) {

	in1 := "http://www.test.com/doc.json"
	in2 := "#/a/b"
	out := in1 + in2

	r1, _ := NewJsonReference(in1)
	r2, _ := NewJsonReference(in2)

	result, err := r1.Inherits(r2)
	if err != nil {
		t.Errorf("Inherits(%s,%s) error %s", r1.String(), r2.String(), err.Error())
	}

	if result.String() != out {
		t.Errorf("Inherits(%s,%s) = %s, expect %s", r1.String(), r2.String(), result.String(), out)
	}
}

func TestInheritsInvalid(t *testing.T) {

	in1 := "http://www.test.com/doc.json"
	in2 := "http://www.test2.com/doc.json#bla"

	r1, _ := NewJsonReference(in1)
	r2, _ := NewJsonReference(in2)

	_, err := r1.Inherits(r2)
	if err == nil {
		t.Errorf("Inherits(%s,%s) should fail", r1.String(), r2.String())
	}

}

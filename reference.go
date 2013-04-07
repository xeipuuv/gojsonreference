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
// description		Main and unique file.
// 
// created      	26-02-2013

package gojsonreference

import (
	"errors"
	"github.com/sigu-399/gojsonpointer"
	"net/url"
	"strings"
)

const (
	const_fragment_char = `#`
)

func NewJsonReference(jsonReferenceString string) (JsonReference, error) {

	var r JsonReference
	err := r.parse(jsonReferenceString)
	return r, err

}

type JsonReference struct {
	referenceUrl     *url.URL
	referencePointer gojsonpointer.JsonPointer

	HasFullUrl      bool
	HasUrlPathOnly  bool
	HasFragmentOnly bool
	HasFileScheme   bool
	HasFullFilePath bool
}

func (r *JsonReference) GetUrl() *url.URL {
	return r.referenceUrl
}

func (r *JsonReference) GetPointer() *gojsonpointer.JsonPointer {
	return &r.referencePointer
}

func (r *JsonReference) String() string {

	if r.referenceUrl != nil {
		return r.referenceUrl.String()
	}

	if r.HasFragmentOnly {
		return const_fragment_char + r.referencePointer.String()
	}

	return r.referencePointer.String()
}

func (r *JsonReference) IsCanonical() bool {
	return (r.HasFileScheme && r.HasFullFilePath) || (!r.HasFileScheme && r.HasFullUrl)
}

// "Constructor", parses the given string JSON reference
func (r *JsonReference) parse(jsonReferenceString string) error {

	var err error

	// fragment only
	if strings.HasPrefix(jsonReferenceString, const_fragment_char) {
		r.referencePointer, err = gojsonpointer.NewJsonPointer(jsonReferenceString[1:])
		if err != nil {
			return nil
		}
		r.HasFragmentOnly = true
	} else {

		r.referenceUrl, err = url.Parse(jsonReferenceString)
		if err != nil {
			return nil
		}

		if r.referenceUrl.Scheme != "" && r.referenceUrl.Host != "" {
			r.HasFullUrl = true
		} else {
			r.HasUrlPathOnly = true
		}

		r.HasFileScheme = r.referenceUrl.Scheme == "file"

		r.HasFullFilePath = strings.HasPrefix(r.GetUrl().Path, "//")

		r.referencePointer, err = gojsonpointer.NewJsonPointer(r.referenceUrl.Fragment)
		if err != nil {
			return nil
		}
	}

	return nil
}

// Creates a new reference from a parent and a child
// If the child cannot inherit from the parent, an error is returned
func (r *JsonReference) Inherits(child JsonReference) (*JsonReference, error) {

	if !r.HasFileScheme && !child.HasFileScheme {
		return r.inheritsImplHttp(child)
	}

	if r.HasFileScheme && child.HasFileScheme {
		return r.inheritsImplFile(child)
	}

	return nil, errors.New("References are not compatible")

}

func (r *JsonReference) inheritsImplFile(child JsonReference) (*JsonReference, error) {

	if !r.HasFullFilePath {
		return nil, errors.New("Parent reference must be canonical")
	}

	if !child.HasFullFilePath {
		return nil, errors.New("Child reference must be canonical")
	}

	inheritedReference, err := NewJsonReference(child.String())
	if err != nil {
		return nil, err
	}

	return &inheritedReference, nil
}

func (r *JsonReference) inheritsImplHttp(child JsonReference) (*JsonReference, error) {

	if !r.HasFullUrl {
		return nil, errors.New("Parent reference must be canonical")
	}

	if r.HasFullUrl && child.HasFullUrl {
		if r.referenceUrl.Scheme != child.referenceUrl.Scheme {
			return nil, errors.New("References have different schemes")
		}
		if r.referenceUrl.Host != child.referenceUrl.Host {
			return nil, errors.New("References have different hosts")
		}
	}

	inheritedReference, err := NewJsonReference(r.String())
	if err != nil {
		return nil, err
	}

	if child.HasFragmentOnly {
		inheritedReference.referenceUrl.Fragment = child.referencePointer.String()
		inheritedReference.referencePointer = child.referencePointer
	}
	if child.HasUrlPathOnly {
		inheritedReference.referenceUrl.Path = child.referenceUrl.Path
	}
	if child.HasFullUrl {
		inheritedReference.referenceUrl.Fragment = child.referenceUrl.Fragment
		inheritedReference.referenceUrl.Path = child.referenceUrl.Path
	}
	return &inheritedReference, nil

	return nil, nil
}

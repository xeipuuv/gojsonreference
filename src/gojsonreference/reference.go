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
	"gojsonpointer"
	"net/url"
	"strings"
)

const (
	FRAGMENT_CHAR = `#`
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
		return FRAGMENT_CHAR + r.referencePointer.String()
	}

	return r.referencePointer.String()
}

// "Constructor", parses the given string JSON reference
func (r *JsonReference) parse(jsonReferenceString string) error {

	var err error

	// fragment only
	if strings.HasPrefix(jsonReferenceString, FRAGMENT_CHAR) {
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

		r.referencePointer, err = gojsonpointer.NewJsonPointer(r.referenceUrl.Fragment)
		if err != nil {
			return nil
		}
	}

	return nil
}

// Creates a new reference from a parent and a child
// If the child cannot inherit from the parent, an error is returned
func Inherits(parent JsonReference, child JsonReference) (*JsonReference, error) {

	if !parent.HasFullUrl {
		return nil, errors.New("Parent reference must be canonical")
	}

	if parent.HasFullUrl && child.HasFullUrl {
		if parent.referenceUrl.Scheme != child.referenceUrl.Scheme {
			return nil, errors.New("References have different schemes")
		}
		if parent.referenceUrl.Host != child.referenceUrl.Host {
			return nil, errors.New("References have different hosts")
		}
	}

	inheritedReference, err := NewJsonReference(parent.String())
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

// @author       sigu-399
// @description  An implementation of JSON Reference - Go language
// @created      26-02-2013

package gojsonreference

import (
	"gojsonpointer"
	"net/url"
	"strings"
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

	return r.referencePointer.String()
}

func (r *JsonReference) parse(jsonReferenceString string) error {

	var err error

	// fragment only
	if strings.HasPrefix(jsonReferenceString, "#") {
		r.referencePointer, err = gojsonpointer.NewJsonPointer(jsonReferenceString)
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

		r.referencePointer, err = gojsonpointer.NewJsonPointer(r.referenceUrl.Fragment)
		if err != nil {
			return nil
		}
	}

	return nil
}

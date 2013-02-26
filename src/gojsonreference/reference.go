// @author       sigu-399
// @description  An implementation of JSON Reference - Go language
// @created      26-02-2013

package gojsonreference

import (
	"gojsonpointer"
	"net/url"
)

func NewJsonReference(jsonReferenceString string) (JsonReference, error) {

	var r JsonReference
	err := r.parse(jsonReferenceString)
	return r, err

}

type JsonReference struct {
	referenceUrl     *url.URL
	referencePointer gojsonpointer.JsonPointer
}

func (r *JsonReference) GetUrl() *url.URL {
	return r.referenceUrl
}

func (r *JsonReference) GetPointer() *gojsonpointer.JsonPointer {
	return &r.referencePointer
}

func (r *JsonReference) parse(jsonReferenceString string) error {

	var err error

	r.referenceUrl, err = url.Parse(jsonReferenceString)
	if err == nil {
		r.referencePointer, err = gojsonpointer.NewJsonPointer(r.referenceUrl.Fragment)
	}

	return err
}
// @author       sigu-399
// @description  An implementation of JSON Reference - Go language
// @created      26-02-2013

package gojsonreference

import ()

func NewJsonReference(jsonReferenceString string) (JsonReference, error) {

	var r JsonReference
	err := r.parse(jsonReferenceString)
	return r, err

}

type JsonReference struct {
}

func (r *JsonReference) parse(JsonReferenceString string) error {

	var err error

	return err
}

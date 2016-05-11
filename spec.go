package main

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

type specData struct {
	OperationIDs                   map[string]*spec.Operation
	OperationIDByTag               map[string]map[string]*spec.Operation
	OperationByTagAndPathAndMethod map[string]map[string]map[string]*spec.Operation

	doc *loads.Document
}

func newSpecData(doc *loads.Document) (*specData, error) {
	s := &specData{
		OperationIDs:                   make(map[string]*spec.Operation),
		OperationIDByTag:               make(map[string]map[string]*spec.Operation),
		OperationByTagAndPathAndMethod: make(map[string]map[string]map[string]*spec.Operation),

		doc: doc,
	}

	err := s.Validate()
	if err != nil {
		return nil, fmt.Errorf("validation error: %s", err)
	}

	return s, nil
}

func (s *specData) Validate() error {
	err := s.UniqueAndRequiredOperationID()
	if err != nil {
		return err
	}

	return nil
}

// UniqueAndRequiredOperationID check for non unique and not specified
// operationId. It also fills the OperationIDs map.
func (s *specData) UniqueAndRequiredOperationID() error {
	m := make(map[string]string)
	for method, av := range s.doc.Analyzer.Operations() {
		for path, bv := range av {
			if bv.ID == "" {
				return fmt.Errorf("required operationId: '%s %s'", method, path)
			}
			oid, ok := m[bv.ID]
			if !ok {
				m[bv.ID] = method + " " + path
				s.OperationIDs[bv.ID] = bv
			} else {
				return fmt.Errorf("duplicate operationId: %s '%s %s' and '%s'", bv.ID, oid, method, path)
			}

			tags := bv.Tags
			var tag string
			if len(tags) != 0 {
				tag = tags[0]
			}

			s.putOperationIDByTag(tag, bv.ID, bv)
			s.putOperationByTagAndPathAndMethod(tag, path, method, bv)
		}
	}

	return nil
}

func (s *specData) putOperationIDByTag(tag, id string, op *spec.Operation) {
	_t, ok := s.OperationIDByTag[tag]
	if !ok {
		_t = make(map[string]*spec.Operation)
		s.OperationIDByTag[tag] = _t
	}

	_t[id] = op
}

func (s *specData) putOperationByTagAndPathAndMethod(tag, path, method string, op *spec.Operation) {
	_t, ok := s.OperationByTagAndPathAndMethod[tag]
	if !ok {
		_t = make(map[string]map[string]*spec.Operation)
		s.OperationByTagAndPathAndMethod[tag] = _t
	}

	_m, ok := _t[path]
	if !ok {
		_m = make(map[string]*spec.Operation)
		_t[path] = _m
	}
	_m[method] = op
}

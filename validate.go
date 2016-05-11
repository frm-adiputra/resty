package main

import (
	"fmt"

	"github.com/go-openapi/loads"
)

type doc loads.Document

func (v doc) Validate() error {
	err := v.UniqueAndRequiredOperationID()
	if err != nil {
		return err
	}

	return nil
}

func (v doc) UniqueAndRequiredOperationID() error {
	m := make(map[string]string)
	for method, av := range loads.Document(v).Analyzer.Operations() {
		for path, bv := range av {
			if bv.ID == "" {
				return fmt.Errorf("required operationId: '%s %s'", method, path)
			}
			oid, ok := m[bv.ID]
			if !ok {
				m[bv.ID] = method + " " + path
			} else {
				return fmt.Errorf("duplicate operationId: %s '%s %s' and '%s'", bv.ID, oid, method, path)
			}
		}
	}

	return nil
}

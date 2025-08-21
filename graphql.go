package forest

import (
	"bytes"
	"encoding/json"
	"io"
)

type Map map[string]interface{}

// GraphQLRequest is used to model both a query or a mutation request
type GraphQLRequest struct {
	Query         string                 `json:"query,omitempty"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// NewGraphQLRequest returns a new Request (for query or mutation) without any variables.
func NewGraphQLRequest(query, operation string, vars ...Map) GraphQLRequest {
	initVars := map[string]interface{}{}
	for _, each := range vars {
		for k, v := range each {
			initVars[k] = v
		}
	}
	return GraphQLRequest{Query: query, OperationName: operation, Variables: initVars}
}

// WithVariablesFromString returns a copy of the request with decoded variables. Returns an error if the jsonhash cannot be converted.
func (r GraphQLRequest) WithVariablesFromString(jsonhash string) (GraphQLRequest, error) {
	vars := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonhash), &vars)
	if err != nil {
		return r, err
	}
	r.Variables = vars
	return r, nil
}

// Reader returns a new reader for sending it using a HTTP request.
// Deprecated: Use ReaderWithError() to handle errors.
func (r GraphQLRequest) Reader() io.Reader {
	reader, err := r.ReaderWithError()
	if err != nil {
		// For backward compatibility, panic on error.
		panic(err)
	}
	return reader
}

// ReaderWithError returns a new reader for sending it using a HTTP request, and returns an error if marshalling fails.
func (r GraphQLRequest) ReaderWithError() (io.Reader, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

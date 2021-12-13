package forest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

type Map map[string]interface{}

// GraphQLRequest is used to model both a query or a mutation request
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// NewGraphQLRequest returns a new Request without any variables.
func NewGraphQLRequest(query, operation string) (GraphQLRequest, error) {
	if query == "" {
		return GraphQLRequest{}, errors.New("query parameter cannot be empty")
	}
	if operation == "" {
		return GraphQLRequest{}, errors.New("operation parameter cannot be empty")
	}
	return GraphQLRequest{Query: query, OperationName: operation, Variables: map[string]interface{}{}}, nil
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
func (r GraphQLRequest) Reader() io.Reader {
	data, _ := json.Marshal(r)
	return bytes.NewReader(data)
}

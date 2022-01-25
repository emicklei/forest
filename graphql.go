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
	if len(vars) > 0 {
		initVars = vars[0] // merge all?
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
func (r GraphQLRequest) Reader() io.Reader {
	data, _ := json.Marshal(r)
	return bytes.NewReader(data)
}

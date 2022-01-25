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
	Mutation      string                 `json:"mutation,omitempty"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// NewGraphQLQuery returns a new Request without any variables.
func NewGraphQLQuery(query, operation string) GraphQLRequest {
	return GraphQLRequest{Query: query, OperationName: operation, Variables: map[string]interface{}{}}
}

// NewGraphQLMutation returns a new Request without any variables.
func NewGraphQLMutation(mutation, operation string) GraphQLRequest {
	return GraphQLRequest{Mutation: mutation, OperationName: operation, Variables: map[string]interface{}{}}
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

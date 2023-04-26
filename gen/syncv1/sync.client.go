// Package syncv1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package syncv1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateSyncJob request with any body
	CreateSyncJobWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateSyncJob(ctx context.Context, body CreateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateSyncJob request with any body
	UpdateSyncJobWithBody(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateSyncJob(ctx context.Context, jobId string, body UpdateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateJobIssue request with any body
	CreateJobIssueWithBody(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateJobIssue(ctx context.Context, jobId string, body CreateJobIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateSyncJobWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSyncJobRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSyncJob(ctx context.Context, body CreateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSyncJobRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateSyncJobWithBody(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateSyncJobRequestWithBody(c.Server, jobId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateSyncJob(ctx context.Context, jobId string, body UpdateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateSyncJobRequest(c.Server, jobId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateJobIssueWithBody(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateJobIssueRequestWithBody(c.Server, jobId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateJobIssue(ctx context.Context, jobId string, body CreateJobIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateJobIssueRequest(c.Server, jobId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateSyncJobRequest calls the generic CreateSyncJob builder with application/json body
func NewCreateSyncJobRequest(server string, body CreateSyncJobJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateSyncJobRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateSyncJobRequestWithBody generates requests for CreateSyncJob with any type of body
func NewCreateSyncJobRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewUpdateSyncJobRequest calls the generic UpdateSyncJob builder with application/json body
func NewUpdateSyncJobRequest(server string, jobId string, body UpdateSyncJobJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateSyncJobRequestWithBody(server, jobId, "application/json", bodyReader)
}

// NewUpdateSyncJobRequestWithBody generates requests for UpdateSyncJob with any type of body
func NewUpdateSyncJobRequestWithBody(server string, jobId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "job_id", runtime.ParamLocationPath, jobId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewCreateJobIssueRequest calls the generic CreateJobIssue builder with application/json body
func NewCreateJobIssueRequest(server string, jobId string, body CreateJobIssueJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateJobIssueRequestWithBody(server, jobId, "application/json", bodyReader)
}

// NewCreateJobIssueRequestWithBody generates requests for CreateJobIssue with any type of body
func NewCreateJobIssueRequestWithBody(server string, jobId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "job_id", runtime.ParamLocationPath, jobId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs/%s/issues", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateSyncJob request with any body
	CreateSyncJobWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSyncJobResponse, error)

	CreateSyncJobWithResponse(ctx context.Context, body CreateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSyncJobResponse, error)

	// UpdateSyncJob request with any body
	UpdateSyncJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSyncJobResponse, error)

	UpdateSyncJobWithResponse(ctx context.Context, jobId string, body UpdateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSyncJobResponse, error)

	// CreateJobIssue request with any body
	CreateJobIssueWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateJobIssueResponse, error)

	CreateJobIssueWithResponse(ctx context.Context, jobId string, body CreateJobIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateJobIssueResponse, error)
}

type CreateSyncJobResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *AnyJobResponse
	JSON403      *ApiError
	JSON429      *ApiError
	JSON500      *ApiError
}

// Status returns HTTPResponse.Status
func (r CreateSyncJobResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateSyncJobResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateSyncJobResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnyJobResponse
	JSON403      *ApiError
	JSON429      *ApiError
	JSON500      *ApiError
}

// Status returns HTTPResponse.Status
func (r UpdateSyncJobResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateSyncJobResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateJobIssueResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *CreateIssueResponse
	JSON403      *ApiError
	JSON429      *ApiError
	JSON500      *ApiError
}

// Status returns HTTPResponse.Status
func (r CreateJobIssueResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateJobIssueResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateSyncJobWithBodyWithResponse request with arbitrary body returning *CreateSyncJobResponse
func (c *ClientWithResponses) CreateSyncJobWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSyncJobResponse, error) {
	rsp, err := c.CreateSyncJobWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSyncJobResponse(rsp)
}

func (c *ClientWithResponses) CreateSyncJobWithResponse(ctx context.Context, body CreateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSyncJobResponse, error) {
	rsp, err := c.CreateSyncJob(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSyncJobResponse(rsp)
}

// UpdateSyncJobWithBodyWithResponse request with arbitrary body returning *UpdateSyncJobResponse
func (c *ClientWithResponses) UpdateSyncJobWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSyncJobResponse, error) {
	rsp, err := c.UpdateSyncJobWithBody(ctx, jobId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateSyncJobResponse(rsp)
}

func (c *ClientWithResponses) UpdateSyncJobWithResponse(ctx context.Context, jobId string, body UpdateSyncJobJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSyncJobResponse, error) {
	rsp, err := c.UpdateSyncJob(ctx, jobId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateSyncJobResponse(rsp)
}

// CreateJobIssueWithBodyWithResponse request with arbitrary body returning *CreateJobIssueResponse
func (c *ClientWithResponses) CreateJobIssueWithBodyWithResponse(ctx context.Context, jobId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateJobIssueResponse, error) {
	rsp, err := c.CreateJobIssueWithBody(ctx, jobId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateJobIssueResponse(rsp)
}

func (c *ClientWithResponses) CreateJobIssueWithResponse(ctx context.Context, jobId string, body CreateJobIssueJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateJobIssueResponse, error) {
	rsp, err := c.CreateJobIssue(ctx, jobId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateJobIssueResponse(rsp)
}

// ParseCreateSyncJobResponse parses an HTTP response from a CreateSyncJobWithResponse call
func ParseCreateSyncJobResponse(rsp *http.Response) (*CreateSyncJobResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateSyncJobResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest AnyJobResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseUpdateSyncJobResponse parses an HTTP response from a UpdateSyncJobWithResponse call
func ParseUpdateSyncJobResponse(rsp *http.Response) (*UpdateSyncJobResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateSyncJobResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnyJobResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseCreateJobIssueResponse parses an HTTP response from a CreateJobIssueWithResponse call
func ParseCreateJobIssueResponse(rsp *http.Response) (*CreateJobIssueResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateJobIssueResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest CreateIssueResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for CategoryName.
const (
	Food   CategoryName = "food"
	Music  CategoryName = "music"
	Sports CategoryName = "sports"
)

// AlbumCreateRequest defines model for AlbumCreateRequest.
type AlbumCreateRequest struct {
	ReleaseDate *ReleaseDate `json:"ReleaseDate,omitempty"`
	Category    Category     `json:"category"`
	Title       string       `json:"title"`
}

// AlbumResponse defines model for AlbumResponse.
type AlbumResponse struct {
	ReleaseDate *ReleaseDate `json:"ReleaseDate,omitempty"`
	Anniversary Anniversary  `json:"anniversary"`
	Category    Category     `json:"category"`
	Id          int          `json:"id"`
	Title       string       `json:"title"`
}

// AlbumUpdateRequest defines model for AlbumUpdateRequest.
type AlbumUpdateRequest struct {
	Category *Category `json:"category,omitempty"`
	Title    *string   `json:"title,omitempty"`
}

// Anniversary defines model for Anniversary.
type Anniversary = int

// Category defines model for Category.
type Category struct {
	Id   *int         `json:"id,omitempty"`
	Name CategoryName `json:"name"`
}

// CategoryName defines model for Category.Name.
type CategoryName string

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// ReleaseDate defines model for ReleaseDate.
type ReleaseDate = openapi_types.Date

// CreateAlbumJSONRequestBody defines body for CreateAlbum for application/json ContentType.
type CreateAlbumJSONRequestBody = AlbumCreateRequest

// UpdateAlbumIdJSONRequestBody defines body for UpdateAlbumId for application/json ContentType.
type UpdateAlbumIdJSONRequestBody = AlbumUpdateRequest

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
	// CreateAlbumWithBody request with any body
	CreateAlbumWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAlbum(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteAlbumId request
	DeleteAlbumId(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAlbumId request
	GetAlbumId(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateAlbumIdWithBody request with any body
	UpdateAlbumIdWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateAlbumId(ctx context.Context, id int, body UpdateAlbumIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateAlbumWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAlbumRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAlbum(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAlbumRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteAlbumId(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteAlbumIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAlbumId(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAlbumIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAlbumIdWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAlbumIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateAlbumId(ctx context.Context, id int, body UpdateAlbumIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateAlbumIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateAlbumRequest calls the generic CreateAlbum builder with application/json body
func NewCreateAlbumRequest(server string, body CreateAlbumJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAlbumRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateAlbumRequestWithBody generates requests for CreateAlbum with any type of body
func NewCreateAlbumRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album")
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

// NewDeleteAlbumIdRequest generates requests for DeleteAlbumId
func NewDeleteAlbumIdRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetAlbumIdRequest generates requests for GetAlbumId
func NewGetAlbumIdRequest(server string, id int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateAlbumIdRequest calls the generic UpdateAlbumId builder with application/json body
func NewUpdateAlbumIdRequest(server string, id int, body UpdateAlbumIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateAlbumIdRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateAlbumIdRequestWithBody generates requests for UpdateAlbumId with any type of body
func NewUpdateAlbumIdRequestWithBody(server string, id int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/album/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
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
	// CreateAlbumWithBodyWithResponse request with any body
	CreateAlbumWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error)

	CreateAlbumWithResponse(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error)

	// DeleteAlbumIdWithResponse request
	DeleteAlbumIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteAlbumIdResponse, error)

	// GetAlbumIdWithResponse request
	GetAlbumIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAlbumIdResponse, error)

	// UpdateAlbumIdWithBodyWithResponse request with any body
	UpdateAlbumIdWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAlbumIdResponse, error)

	UpdateAlbumIdWithResponse(ctx context.Context, id int, body UpdateAlbumIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAlbumIdResponse, error)
}

type CreateAlbumResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *AlbumResponse
	JSON400      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r CreateAlbumResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAlbumResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteAlbumIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r DeleteAlbumIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteAlbumIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAlbumIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AlbumResponse
	JSON400      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetAlbumIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAlbumIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateAlbumIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AlbumResponse
	JSON400      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r UpdateAlbumIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateAlbumIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateAlbumWithBodyWithResponse request with arbitrary body returning *CreateAlbumResponse
func (c *ClientWithResponses) CreateAlbumWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error) {
	rsp, err := c.CreateAlbumWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAlbumResponse(rsp)
}

func (c *ClientWithResponses) CreateAlbumWithResponse(ctx context.Context, body CreateAlbumJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAlbumResponse, error) {
	rsp, err := c.CreateAlbum(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAlbumResponse(rsp)
}

// DeleteAlbumIdWithResponse request returning *DeleteAlbumIdResponse
func (c *ClientWithResponses) DeleteAlbumIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*DeleteAlbumIdResponse, error) {
	rsp, err := c.DeleteAlbumId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteAlbumIdResponse(rsp)
}

// GetAlbumIdWithResponse request returning *GetAlbumIdResponse
func (c *ClientWithResponses) GetAlbumIdWithResponse(ctx context.Context, id int, reqEditors ...RequestEditorFn) (*GetAlbumIdResponse, error) {
	rsp, err := c.GetAlbumId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAlbumIdResponse(rsp)
}

// UpdateAlbumIdWithBodyWithResponse request with arbitrary body returning *UpdateAlbumIdResponse
func (c *ClientWithResponses) UpdateAlbumIdWithBodyWithResponse(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateAlbumIdResponse, error) {
	rsp, err := c.UpdateAlbumIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAlbumIdResponse(rsp)
}

func (c *ClientWithResponses) UpdateAlbumIdWithResponse(ctx context.Context, id int, body UpdateAlbumIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateAlbumIdResponse, error) {
	rsp, err := c.UpdateAlbumId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateAlbumIdResponse(rsp)
}

// ParseCreateAlbumResponse parses an HTTP response from a CreateAlbumWithResponse call
func ParseCreateAlbumResponse(rsp *http.Response) (*CreateAlbumResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAlbumResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest AlbumResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseDeleteAlbumIdResponse parses an HTTP response from a DeleteAlbumIdWithResponse call
func ParseDeleteAlbumIdResponse(rsp *http.Response) (*DeleteAlbumIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteAlbumIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetAlbumIdResponse parses an HTTP response from a GetAlbumIdWithResponse call
func ParseGetAlbumIdResponse(rsp *http.Response) (*GetAlbumIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAlbumIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AlbumResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseUpdateAlbumIdResponse parses an HTTP response from a UpdateAlbumIdWithResponse call
func ParseUpdateAlbumIdResponse(rsp *http.Response) (*UpdateAlbumIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateAlbumIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AlbumResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new album
	// (POST /album)
	CreateAlbum(c *gin.Context)
	// Delete a album by ID
	// (DELETE /album/{id})
	DeleteAlbumId(c *gin.Context, id int)
	// Find album by ID
	// (GET /album/{id})
	GetAlbumId(c *gin.Context, id int)
	// Update a album by ID
	// (PATCH /album/{id})
	UpdateAlbumId(c *gin.Context, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// CreateAlbum operation middleware
func (siw *ServerInterfaceWrapper) CreateAlbum(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateAlbum(c)
}

// DeleteAlbumId operation middleware
func (siw *ServerInterfaceWrapper) DeleteAlbumId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteAlbumId(c, id)
}

// GetAlbumId operation middleware
func (siw *ServerInterfaceWrapper) GetAlbumId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAlbumId(c, id)
}

// UpdateAlbumId operation middleware
func (siw *ServerInterfaceWrapper) UpdateAlbumId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateAlbumId(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/album", wrapper.CreateAlbum)
	router.DELETE(options.BaseURL+"/album/:id", wrapper.DeleteAlbumId)
	router.GET(options.BaseURL+"/album/:id", wrapper.GetAlbumId)
	router.PATCH(options.BaseURL+"/album/:id", wrapper.UpdateAlbumId)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWTU/jPBD+K9G87zFqUkBa5BuUBVUrsSu0e0IcTDxtjeIP7AkrVOW/r2y3kDTlY1cF",
	"LqiXNh37medjpl1CZZQ1GjV5YEvw1QIVj2+P6utGTRxywgu8bdBTeGqdsehIYqy5wBq5xxNOGD7+73AG",
	"DP4rHi8tVjcW3dI2h4oTzo27f+nYZF3X5kCS6ohD9xaBgScn9RzaNgeHt410KIBdrso6EOHrR/CrfH3e",
	"XN9gReHmyPUCvTXa4y5pcq3lHTrPX2Z61Cn9R4Gk6KgjNeEc3V8IJwXk29Trknillr+seC43O3Z/2ENf",
	"9qEik04D/dae0lBzFdFRNyqINTMmyKUaLyvIwVvjyHcEeULjeM022b46Z9zTEVToPZ+/wsR14TaMjRzP",
	"jFOcgEHwCoadh0TpmYmYSXw4M9nRj2n2E5Wt06EgszQaGIxH5agMMMai5lYCg/1ROdqHHCynRaRR8JCO",
	"SM+kYASSnKTRUwEM0r6JEYLECz0dGxGNqowm1PEUt7aWVTxX3PgAv95dL47ZcK21fQ3JNRgfJC9i33vl",
	"eLcdPDgdwQX6yklLScjUnQhSHpTlznD7CduCe8xF9qhJDr5RKk7QqqOMZxp/Z8nCUJDcLJZStAFcYI0p",
	"WX1TT+LzSHsqYhocV0joPLDLJciAHRIC6ylLm6jvSN5huTmd7dXAroPUUJfeuckmKx0/VNmAffB+2OeG",
	"slPTaLHhaXIl48nP7Po+m56E5ua4ZS7PkN7Xv/L9xu37t888AINTqcVmFiynajFMQ/p1f/NAvNHy7/83",
	"aVfb/6PSl7oRnxF80GJzJcUadHfreDWuBgYLIsuKohzFFzssD8uCW1ncjaHNN4pqU/F6YTw9Xzbe+xJv",
	"G/fLrto/AQAA//+hyv5VJg0AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

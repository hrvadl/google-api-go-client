// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package workflowexecutions provides access to the Workflow Executions API.
//
// For product documentation, see: https://cloud.google.com/workflows
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/workflowexecutions/v1beta"
//	...
//	ctx := context.Background()
//	workflowexecutionsService, err := workflowexecutions.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	workflowexecutionsService, err := workflowexecutions.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	workflowexecutionsService, err := workflowexecutions.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package workflowexecutions // import "google.golang.org/api/workflowexecutions/v1beta"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "workflowexecutions:v1beta"
const apiName = "workflowexecutions"
const apiVersion = "v1beta"
const basePath = "https://workflowexecutions.googleapis.com/"
const basePathTemplate = "https://workflowexecutions.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://workflowexecutions.mtls.googleapis.com/"
const defaultUniverseDomain = "googleapis.com"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the email
	// address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultEndpointTemplate(basePathTemplate))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	opts = append(opts, internaloption.WithDefaultUniverseDomain(defaultUniverseDomain))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.Workflows = NewProjectsLocationsWorkflowsService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	Workflows *ProjectsLocationsWorkflowsService
}

func NewProjectsLocationsWorkflowsService(s *Service) *ProjectsLocationsWorkflowsService {
	rs := &ProjectsLocationsWorkflowsService{s: s}
	rs.Executions = NewProjectsLocationsWorkflowsExecutionsService(s)
	return rs
}

type ProjectsLocationsWorkflowsService struct {
	s *Service

	Executions *ProjectsLocationsWorkflowsExecutionsService
}

func NewProjectsLocationsWorkflowsExecutionsService(s *Service) *ProjectsLocationsWorkflowsExecutionsService {
	rs := &ProjectsLocationsWorkflowsExecutionsService{s: s}
	return rs
}

type ProjectsLocationsWorkflowsExecutionsService struct {
	s *Service
}

// CancelExecutionRequest: Request for the CancelExecution method.
type CancelExecutionRequest struct {
}

// Error: Error describes why the execution was abnormally terminated.
type Error struct {
	// Context: Human-readable stack trace string.
	Context string `json:"context,omitempty"`
	// Payload: Error message and data returned represented as a JSON string.
	Payload string `json:"payload,omitempty"`
	// StackTrace: Stack trace with detailed information of where error was
	// generated.
	StackTrace *StackTrace `json:"stackTrace,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Context") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Context") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Error) MarshalJSON() ([]byte, error) {
	type NoMethod Error
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Execution: A running instance of a Workflow
// (/workflows/docs/reference/rest/v1beta/projects.locations.workflows).
type Execution struct {
	// Argument: Input parameters of the execution represented as a JSON string.
	// The size limit is 32KB. *Note*: If you are using the REST API directly to
	// run your workflow, you must escape any JSON string value of `argument`.
	// Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
	Argument string `json:"argument,omitempty"`
	// CallLogLevel: The call logging level associated to this execution.
	//
	// Possible values:
	//   "CALL_LOG_LEVEL_UNSPECIFIED" - No call logging level specified.
	//   "LOG_ALL_CALLS" - Log all call steps within workflows, all call returns,
	// and all exceptions raised.
	//   "LOG_ERRORS_ONLY" - Log only exceptions that are raised from call steps
	// within workflows.
	CallLogLevel string `json:"callLogLevel,omitempty"`
	// EndTime: Output only. Marks the end of execution, successful or not.
	EndTime string `json:"endTime,omitempty"`
	// Error: Output only. The error which caused the execution to finish
	// prematurely. The value is only present if the execution's state is `FAILED`
	// or `CANCELLED`.
	Error *Error `json:"error,omitempty"`
	// Name: Output only. The resource name of the execution. Format:
	// projects/{project}/locations/{location}/workflows/{workflow}/executions/{exec
	// ution}
	Name string `json:"name,omitempty"`
	// Result: Output only. Output of the execution represented as a JSON string.
	// The value can only be present if the execution's state is `SUCCEEDED`.
	Result string `json:"result,omitempty"`
	// StartTime: Output only. Marks the beginning of execution.
	StartTime string `json:"startTime,omitempty"`
	// State: Output only. Current state of the execution.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Invalid state.
	//   "ACTIVE" - The execution is in progress.
	//   "SUCCEEDED" - The execution finished successfully.
	//   "FAILED" - The execution failed with an error.
	//   "CANCELLED" - The execution was stopped intentionally.
	//   "UNAVAILABLE" - Reserved for future use.
	//   "QUEUED" - Request has been placed in the backlog for processing at a
	// later time.
	State string `json:"state,omitempty"`
	// Status: Output only. Status tracks the current steps and progress data of
	// this execution.
	Status *Status `json:"status,omitempty"`
	// WorkflowRevisionId: Output only. Revision of the workflow this execution is
	// using.
	WorkflowRevisionId string `json:"workflowRevisionId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Argument") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Argument") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Execution) MarshalJSON() ([]byte, error) {
	type NoMethod Execution
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// ListExecutionsResponse: Response for the ListExecutions method.
type ListExecutionsResponse struct {
	// Executions: The executions which match the request.
	Executions []*Execution `json:"executions,omitempty"`
	// NextPageToken: A token, which can be sent as `page_token` to retrieve the
	// next page. If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Executions") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Executions") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *ListExecutionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListExecutionsResponse
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Position: Position contains source position information about the stack
// trace element such as line number, column number and length of the code
// block in bytes.
type Position struct {
	// Column: The source code column position (of the line) the current
	// instruction was generated from.
	Column int64 `json:"column,omitempty,string"`
	// Length: The number of bytes of source code making up this stack trace
	// element.
	Length int64 `json:"length,omitempty,string"`
	// Line: The source code line number the current instruction was generated
	// from.
	Line int64 `json:"line,omitempty,string"`
	// ForceSendFields is a list of field names (e.g. "Column") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Column") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Position) MarshalJSON() ([]byte, error) {
	type NoMethod Position
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// StackTrace: A collection of stack elements (frames) where an error occurred.
type StackTrace struct {
	// Elements: An array of stack elements.
	Elements []*StackTraceElement `json:"elements,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Elements") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Elements") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *StackTrace) MarshalJSON() ([]byte, error) {
	type NoMethod StackTrace
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// StackTraceElement: A single stack element (frame) where an error occurred.
type StackTraceElement struct {
	// Position: The source position information of the stack trace element.
	Position *Position `json:"position,omitempty"`
	// Routine: The routine where the error occurred.
	Routine string `json:"routine,omitempty"`
	// Step: The step the error occurred at.
	Step string `json:"step,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Position") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Position") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *StackTraceElement) MarshalJSON() ([]byte, error) {
	type NoMethod StackTraceElement
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Status: Represents the current status of this execution.
type Status struct {
	// CurrentSteps: A list of currently executing or last executed step names for
	// the workflow execution currently running. If the workflow has succeeded or
	// failed, this is the last attempted or executed step. Presently, if the
	// current step is inside a subworkflow, the list only includes that step. In
	// the future, the list will contain items for each step in the call stack,
	// starting with the outermost step in the `main` subworkflow, and ending with
	// the most deeply nested step.
	CurrentSteps []*Step `json:"currentSteps,omitempty"`
	// ForceSendFields is a list of field names (e.g. "CurrentSteps") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "CurrentSteps") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Step: Represents a step of the workflow this execution is running.
type Step struct {
	// Routine: Name of a routine within the workflow.
	Routine string `json:"routine,omitempty"`
	// Step: Name of a step within the routine.
	Step string `json:"step,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Routine") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Routine") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Step) MarshalJSON() ([]byte, error) {
	type NoMethod Step
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

type ProjectsLocationsWorkflowsExecutionsCancelCall struct {
	s                      *Service
	name                   string
	cancelexecutionrequest *CancelExecutionRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Cancel: Cancels an execution of the given name.
//
//   - name: Name of the execution to be cancelled. Format:
//     projects/{project}/locations/{location}/workflows/{workflow}/executions/{ex
//     ecution}.
func (r *ProjectsLocationsWorkflowsExecutionsService) Cancel(name string, cancelexecutionrequest *CancelExecutionRequest) *ProjectsLocationsWorkflowsExecutionsCancelCall {
	c := &ProjectsLocationsWorkflowsExecutionsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.cancelexecutionrequest = cancelexecutionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cancelexecutionrequest)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.cancel" call.
// Any non-2xx status code is an error. Response headers are in either
// *Execution.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsCancelCall) Do(opts ...googleapi.CallOption) (*Execution, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Execution{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type ProjectsLocationsWorkflowsExecutionsCreateCall struct {
	s          *Service
	parent     string
	execution  *Execution
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new execution using the latest revision of the given
// workflow.
//
//   - parent: Name of the workflow for which an execution should be created.
//     Format: projects/{project}/locations/{location}/workflows/{workflow} The
//     latest revision of the workflow will be used.
func (r *ProjectsLocationsWorkflowsExecutionsService) Create(parent string, execution *Execution) *ProjectsLocationsWorkflowsExecutionsCreateCall {
	c := &ProjectsLocationsWorkflowsExecutionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.execution = execution
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.execution)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/executions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.create" call.
// Any non-2xx status code is an error. Response headers are in either
// *Execution.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsCreateCall) Do(opts ...googleapi.CallOption) (*Execution, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Execution{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type ProjectsLocationsWorkflowsExecutionsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns an execution of the given name.
//
//   - name: Name of the execution to be retrieved. Format:
//     projects/{project}/locations/{location}/workflows/{workflow}/executions/{ex
//     ecution}.
func (r *ProjectsLocationsWorkflowsExecutionsService) Get(name string) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c := &ProjectsLocationsWorkflowsExecutionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// View sets the optional parameter "view": A view defining which fields should
// be filled in the returned execution. The API will default to the FULL view.
//
// Possible values:
//
//	"EXECUTION_VIEW_UNSPECIFIED" - The default / unset value.
//	"BASIC" - Includes only basic metadata about the execution. Following
//
// fields are returned: name, start_time, end_time, state and
// workflow_revision_id.
//
//	"FULL" - Includes all data.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) View(view string) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.get" call.
// Any non-2xx status code is an error. Response headers are in either
// *Execution.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsLocationsWorkflowsExecutionsGetCall) Do(opts ...googleapi.CallOption) (*Execution, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Execution{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type ProjectsLocationsWorkflowsExecutionsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns a list of executions which belong to the workflow with the
// given name. The method returns executions of all workflow revisions.
// Returned executions are ordered by their start time (newest first).
//
//   - parent: Name of the workflow for which the executions should be listed.
//     Format: projects/{project}/locations/{location}/workflows/{workflow}.
func (r *ProjectsLocationsWorkflowsExecutionsService) List(parent string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c := &ProjectsLocationsWorkflowsExecutionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// executions to return per call. Max supported value depends on the selected
// Execution view: it's 10000 for BASIC and 100 for FULL. The default value
// used if the field is not specified is 100, regardless of the selected view.
// Values greater than the max value will be coerced down to it.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) PageSize(pageSize int64) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token, received
// from a previous `ListExecutions` call. Provide this to retrieve the
// subsequent page. When paginating, all other parameters provided to
// `ListExecutions` must match the call that provided the page token.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) PageToken(pageToken string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// View sets the optional parameter "view": A view defining which fields should
// be filled in the returned executions. The API will default to the BASIC
// view.
//
// Possible values:
//
//	"EXECUTION_VIEW_UNSPECIFIED" - The default / unset value.
//	"BASIC" - Includes only basic metadata about the execution. Following
//
// fields are returned: name, start_time, end_time, state and
// workflow_revision_id.
//
//	"FULL" - Includes all data.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) View(view string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Context(ctx context.Context) *ProjectsLocationsWorkflowsExecutionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsWorkflowsExecutionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/executions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "workflowexecutions.projects.locations.workflows.executions.list" call.
// Any non-2xx status code is an error. Response headers are in either
// *ListExecutionsResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Do(opts ...googleapi.CallOption) (*ListExecutionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &ListExecutionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsWorkflowsExecutionsListCall) Pages(ctx context.Context, f func(*ListExecutionsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken"))
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

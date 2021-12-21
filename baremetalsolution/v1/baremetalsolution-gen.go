// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package baremetalsolution provides access to the Bare Metal Solution API.
//
// For product documentation, see: https://cloud.google.com/bare-metal
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/baremetalsolution/v1"
//   ...
//   ctx := context.Background()
//   baremetalsolutionService, err := baremetalsolution.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   baremetalsolutionService, err := baremetalsolution.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   baremetalsolutionService, err := baremetalsolution.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package baremetalsolution // import "google.golang.org/api/baremetalsolution/v1"

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

const apiId = "baremetalsolution:v1"
const apiName = "baremetalsolution"
const apiVersion = "v1"
const basePath = "https://baremetalsolution.googleapis.com/"
const mtlsBasePath = "https://baremetalsolution.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the
	// email address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
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
	rs.Instances = NewProjectsLocationsInstancesService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	Instances *ProjectsLocationsInstancesService
}

func NewProjectsLocationsInstancesService(s *Service) *ProjectsLocationsInstancesService {
	rs := &ProjectsLocationsInstancesService{s: s}
	return rs
}

type ProjectsLocationsInstancesService struct {
	s *Service
}

// ResetInstanceRequest: Request for ResetInstance.
type ResetInstanceRequest struct {
}

// ResetInstanceResponse: Response for ResetInstance.
type ResetInstanceResponse struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// method id "baremetalsolution.projects.locations.instances.resetInstance":

type ProjectsLocationsInstancesResetInstanceCall struct {
	s                    *Service
	instance             string
	resetinstancerequest *ResetInstanceRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// ResetInstance: Perform an ungraceful, hard reset on a machine
// (equivalent to shutting the power off, and then turning it back on).
//
// - instance: Name of the instance to reset.
func (r *ProjectsLocationsInstancesService) ResetInstance(instance string, resetinstancerequest *ResetInstanceRequest) *ProjectsLocationsInstancesResetInstanceCall {
	c := &ProjectsLocationsInstancesResetInstanceCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.instance = instance
	c.resetinstancerequest = resetinstancerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsInstancesResetInstanceCall) Fields(s ...googleapi.Field) *ProjectsLocationsInstancesResetInstanceCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsInstancesResetInstanceCall) Context(ctx context.Context) *ProjectsLocationsInstancesResetInstanceCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsInstancesResetInstanceCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsInstancesResetInstanceCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211220")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resetinstancerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+instance}:resetInstance")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"instance": c.instance,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "baremetalsolution.projects.locations.instances.resetInstance" call.
// Exactly one of *ResetInstanceResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ResetInstanceResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsInstancesResetInstanceCall) Do(opts ...googleapi.CallOption) (*ResetInstanceResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ResetInstanceResponse{
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
	// {
	//   "description": "Perform an ungraceful, hard reset on a machine (equivalent to shutting the power off, and then turning it back on).",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}/instances/{instancesId}:resetInstance",
	//   "httpMethod": "POST",
	//   "id": "baremetalsolution.projects.locations.instances.resetInstance",
	//   "parameterOrder": [
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Required. Name of the instance to reset.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+instance}:resetInstance",
	//   "request": {
	//     "$ref": "ResetInstanceRequest"
	//   },
	//   "response": {
	//     "$ref": "ResetInstanceResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

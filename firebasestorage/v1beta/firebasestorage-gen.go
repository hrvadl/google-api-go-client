// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package firebasestorage provides access to the Cloud Storage for Firebase API.
//
// For product documentation, see: https://firebase.google.com/docs/storage
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
//	import "google.golang.org/api/firebasestorage/v1beta"
//	...
//	ctx := context.Background()
//	firebasestorageService, err := firebasestorage.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// By default, all available scopes (see "Constants") are used to authenticate.
// To restrict scopes, use [google.golang.org/api/option.WithScopes]:
//
//	firebasestorageService, err := firebasestorage.NewService(ctx, option.WithScopes(firebasestorage.FirebaseScope))
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	firebasestorageService, err := firebasestorage.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	firebasestorageService, err := firebasestorage.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package firebasestorage // import "google.golang.org/api/firebasestorage/v1beta"

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

const apiId = "firebasestorage:v1beta"
const apiName = "firebasestorage"
const apiVersion = "v1beta"
const basePath = "https://firebasestorage.googleapis.com/"
const basePathTemplate = "https://firebasestorage.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://firebasestorage.mtls.googleapis.com/"
const defaultUniverseDomain = "googleapis.com"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the email
	// address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and administer all your Firebase data and settings
	FirebaseScope = "https://www.googleapis.com/auth/firebase"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/firebase",
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
	rs.Buckets = NewProjectsBucketsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Buckets *ProjectsBucketsService
}

func NewProjectsBucketsService(s *Service) *ProjectsBucketsService {
	rs := &ProjectsBucketsService{s: s}
	return rs
}

type ProjectsBucketsService struct {
	s *Service
}

// AddFirebaseRequest: The request used to link a Google Cloud Storage bucket
// to a Firebase project.
type AddFirebaseRequest struct {
}

// Bucket: A storage bucket and its relation to a parent Firebase project.
type Bucket struct {
	// Name: Output only. Resource name of the bucket.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Name") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Name") to include in API requests
	// with the JSON null value. By default, fields with empty values are omitted
	// from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Bucket) MarshalJSON() ([]byte, error) {
	type NoMethod Bucket
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use it as
// the request or the response type of an API method. For instance: service Foo
// { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
}

// GoogleFirebaseStorageControlplaneV1alphaMigrateLocationDestructivelyMetadata:
//
//	Metadata for MigrateLocationDestructively LRO.
type GoogleFirebaseStorageControlplaneV1alphaMigrateLocationDestructivelyMetadata struct {
	// CreateTime: The time the LRO was created.
	CreateTime string `json:"createTime,omitempty"`
	// LastUpdateTime: The time the LRO was last updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// State: The current state of the migration.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unspecified state. Should not be used.
	//   "PENDING" - The MigrateLocationDestructively request has passed
	// precondition checks and the bucket migration will begin soon.
	//   "CREATING_TEMP_BUCKET" - Generating a unique bucket name, storing the
	// source -> temp mapping in Spanner, and actually creating the temporary
	// bucket via Bigstore.
	//   "TRANSFERRING_TO_TEMP" - The first STS transfer to move all objects from
	// the source bucket to the temp bucket is underway.
	//   "DELETING_SOURCE_BUCKET" - The source bucket is being emptied and deleted.
	//   "CREATING_DESTINATION_BUCKET" - The source bucket is being recreated in
	// the new location.
	//   "TRANSFERRING_TO_DESTINATION" - The second STS transfer to move all
	// objects from the temp bucket to the final bucket is underway.
	//   "DELETING_TEMP_BUCKET" - The temp bucket is being emptied and deleted.
	//   "SUCCEEDED" - All stages of the migration have completed and the operation
	// has been marked done and updated with a response.
	//   "FAILED" - The migration failed at some stage and it is not possible to
	// continue retrying that stage. Manual recovery may be needed. Rollback is
	// either impossible at this stage, or has been attempted and failed.
	//   "ROLLING_BACK" - The migration has encountered a permanent failure and is
	// now being rolled back so that the source bucket is restored to its original
	// state.
	//   "ROLLED_BACK" - The migration has been successfully rolled back.
	State string `json:"state,omitempty"`
	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "CreateTime") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseStorageControlplaneV1alphaMigrateLocationDestructivelyMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseStorageControlplaneV1alphaMigrateLocationDestructivelyMetadata
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// GoogleFirebaseStorageControlplaneV1betaMigrateLocationDestructivelyMetadata:
// Metadata for MigrateLocationDestructively LRO.
type GoogleFirebaseStorageControlplaneV1betaMigrateLocationDestructivelyMetadata struct {
	// CreateTime: The time the LRO was created.
	CreateTime string `json:"createTime,omitempty"`
	// LastUpdateTime: The time the LRO was last updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// State: The current state of the migration.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unspecified state. Should not be used.
	//   "PENDING" - The MigrateLocationDestructively request has passed
	// precondition checks and the bucket migration will begin soon.
	//   "CREATING_TEMP_BUCKET" - Generating a unique bucket name, storing the
	// source -> temp mapping in Spanner, and actually creating the temporary
	// bucket via Bigstore.
	//   "TRANSFERRING_TO_TEMP" - The first STS transfer to move all objects from
	// the source bucket to the temp bucket is underway.
	//   "DELETING_SOURCE_BUCKET" - The source bucket is being emptied and deleted.
	//   "CREATING_DESTINATION_BUCKET" - The source bucket is being recreated in
	// the new location.
	//   "TRANSFERRING_TO_DESTINATION" - The second STS transfer to move all
	// objects from the temp bucket to the final bucket is underway.
	//   "DELETING_TEMP_BUCKET" - The temp bucket is being emptied and deleted.
	//   "SUCCEEDED" - All stages of the migration have completed and the operation
	// has been marked done and updated with a response.
	//   "FAILED" - The migration failed at some stage and it is not possible to
	// continue retrying that stage. Manual recovery may be needed. Rollback is
	// either impossible at this stage, or has been attempted and failed.
	//   "ROLLING_BACK" - The migration has encountered a permanent failure and is
	// now being rolled back so that the source bucket is restored to its original
	// state.
	//   "ROLLED_BACK" - The migration has been successfully rolled back.
	State string `json:"state,omitempty"`
	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "CreateTime") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *GoogleFirebaseStorageControlplaneV1betaMigrateLocationDestructivelyMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirebaseStorageControlplaneV1betaMigrateLocationDestructivelyMetadata
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// ListBucketsResponse: The response returned by `ListBuckets`.
type ListBucketsResponse struct {
	// Buckets: The list of linked buckets.
	Buckets []*Bucket `json:"buckets,omitempty"`
	// NextPageToken: A token that can be sent as `page_token` to retrieve the next
	// page. If this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Buckets") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Buckets") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *ListBucketsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBucketsResponse
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// RemoveFirebaseRequest: The request used to unlink a Google Cloud Storage
// bucket from a Firebase project.
type RemoveFirebaseRequest struct {
}

type ProjectsBucketsAddFirebaseCall struct {
	s                  *Service
	bucket             string
	addfirebaserequest *AddFirebaseRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// AddFirebase: Links a Google Cloud Storage bucket to a Firebase project.
//
//   - bucket: Resource name of the bucket, mirrors the ID of the underlying
//     Google Cloud Storage bucket,
//     `projects/{project_id_or_number}/buckets/{bucket_id}`.
func (r *ProjectsBucketsService) AddFirebase(bucket string, addfirebaserequest *AddFirebaseRequest) *ProjectsBucketsAddFirebaseCall {
	c := &ProjectsBucketsAddFirebaseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.addfirebaserequest = addfirebaserequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsBucketsAddFirebaseCall) Fields(s ...googleapi.Field) *ProjectsBucketsAddFirebaseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsBucketsAddFirebaseCall) Context(ctx context.Context) *ProjectsBucketsAddFirebaseCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsBucketsAddFirebaseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBucketsAddFirebaseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.addfirebaserequest)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+bucket}:addFirebase")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasestorage.projects.buckets.addFirebase" call.
// Any non-2xx status code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsBucketsAddFirebaseCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
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
	ret := &Bucket{
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

type ProjectsBucketsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single linked storage bucket.
//
//   - name: Resource name of the bucket, mirrors the ID of the underlying Google
//     Cloud Storage bucket,
//     `projects/{project_id_or_number}/buckets/{bucket_id}`.
func (r *ProjectsBucketsService) Get(name string) *ProjectsBucketsGetCall {
	c := &ProjectsBucketsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsBucketsGetCall) Fields(s ...googleapi.Field) *ProjectsBucketsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProjectsBucketsGetCall) IfNoneMatch(entityTag string) *ProjectsBucketsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsBucketsGetCall) Context(ctx context.Context) *ProjectsBucketsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsBucketsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBucketsGetCall) doRequest(alt string) (*http.Response, error) {
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

// Do executes the "firebasestorage.projects.buckets.get" call.
// Any non-2xx status code is an error. Response headers are in either
// *Bucket.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsBucketsGetCall) Do(opts ...googleapi.CallOption) (*Bucket, error) {
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
	ret := &Bucket{
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

type ProjectsBucketsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the linked storage buckets for a project.
//
//   - parent: Resource name of the parent Firebase project,
//     `projects/{project_id_or_number}`.
func (r *ProjectsBucketsService) List(parent string) *ProjectsBucketsListCall {
	c := &ProjectsBucketsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number of
// buckets to return. If not set, the server will use a reasonable default.
func (c *ProjectsBucketsListCall) PageSize(pageSize int64) *ProjectsBucketsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token, received
// from a previous `ListBuckets` call. Provide this to retrieve the subsequent
// page. When paginating, all other parameters provided to `ListBuckets` must
// match the call that provided the page token.
func (c *ProjectsBucketsListCall) PageToken(pageToken string) *ProjectsBucketsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsBucketsListCall) Fields(s ...googleapi.Field) *ProjectsBucketsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProjectsBucketsListCall) IfNoneMatch(entityTag string) *ProjectsBucketsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsBucketsListCall) Context(ctx context.Context) *ProjectsBucketsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsBucketsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBucketsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+parent}/buckets")
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

// Do executes the "firebasestorage.projects.buckets.list" call.
// Any non-2xx status code is an error. Response headers are in either
// *ListBucketsResponse.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsBucketsListCall) Do(opts ...googleapi.CallOption) (*ListBucketsResponse, error) {
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
	ret := &ListBucketsResponse{
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
func (c *ProjectsBucketsListCall) Pages(ctx context.Context, f func(*ListBucketsResponse) error) error {
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

type ProjectsBucketsRemoveFirebaseCall struct {
	s                     *Service
	bucket                string
	removefirebaserequest *RemoveFirebaseRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// RemoveFirebase: Unlinks a linked Google Cloud Storage bucket from a Firebase
// project.
//
//   - bucket: Resource name of the bucket, mirrors the ID of the underlying
//     Google Cloud Storage bucket,
//     `projects/{project_id_or_number}/buckets/{bucket_id}`.
func (r *ProjectsBucketsService) RemoveFirebase(bucket string, removefirebaserequest *RemoveFirebaseRequest) *ProjectsBucketsRemoveFirebaseCall {
	c := &ProjectsBucketsRemoveFirebaseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.bucket = bucket
	c.removefirebaserequest = removefirebaserequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsBucketsRemoveFirebaseCall) Fields(s ...googleapi.Field) *ProjectsBucketsRemoveFirebaseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsBucketsRemoveFirebaseCall) Context(ctx context.Context) *ProjectsBucketsRemoveFirebaseCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsBucketsRemoveFirebaseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBucketsRemoveFirebaseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.removefirebaserequest)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta/{+bucket}:removeFirebase")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"bucket": c.bucket,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firebasestorage.projects.buckets.removeFirebase" call.
// Any non-2xx status code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsBucketsRemoveFirebaseCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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

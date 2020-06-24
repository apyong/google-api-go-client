// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package androidpublisher provides access to the Google Play Developer API.
//
// For product documentation, see: https://developers.google.com/android-publisher
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/androidpublisher/v2"
//   ...
//   ctx := context.Background()
//   androidpublisherService, err := androidpublisher.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   androidpublisherService, err := androidpublisher.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   androidpublisherService, err := androidpublisher.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package androidpublisher // import "google.golang.org/api/androidpublisher/v2"

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

const apiId = "androidpublisher:v2"
const apiName = "androidpublisher"
const apiVersion = "v2"
const basePath = "https://www.googleapis.com/androidpublisher/v2/applications/"

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Play Developer account
	AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/androidpublisher",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
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
	s.Purchases = NewPurchasesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Purchases *PurchasesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewPurchasesService(s *Service) *PurchasesService {
	rs := &PurchasesService{s: s}
	rs.Products = NewPurchasesProductsService(s)
	rs.Voidedpurchases = NewPurchasesVoidedpurchasesService(s)
	return rs
}

type PurchasesService struct {
	s *Service

	Products *PurchasesProductsService

	Voidedpurchases *PurchasesVoidedpurchasesService
}

func NewPurchasesProductsService(s *Service) *PurchasesProductsService {
	rs := &PurchasesProductsService{s: s}
	return rs
}

type PurchasesProductsService struct {
	s *Service
}

func NewPurchasesVoidedpurchasesService(s *Service) *PurchasesVoidedpurchasesService {
	rs := &PurchasesVoidedpurchasesService{s: s}
	return rs
}

type PurchasesVoidedpurchasesService struct {
	s *Service
}

type PageInfo struct {
	ResultPerPage int64 `json:"resultPerPage,omitempty"`

	StartIndex int64 `json:"startIndex,omitempty"`

	TotalResults int64 `json:"totalResults,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResultPerPage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResultPerPage") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PageInfo) MarshalJSON() ([]byte, error) {
	type NoMethod PageInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ProductPurchase: A ProductPurchase resource indicates the status of a
// user's inapp product purchase.
type ProductPurchase struct {
	// ConsumptionState: The consumption state of the inapp product.
	// Possible values are:
	// - Yet to be consumed
	// - Consumed
	ConsumptionState int64 `json:"consumptionState,omitempty"`

	// DeveloperPayload: A developer-specified string that contains
	// supplemental information about an order.
	DeveloperPayload string `json:"developerPayload,omitempty"`

	// Kind: This kind represents an inappPurchase object in the
	// androidpublisher service.
	Kind string `json:"kind,omitempty"`

	// OrderId: The order id associated with the purchase of the inapp
	// product.
	OrderId string `json:"orderId,omitempty"`

	// PurchaseState: The purchase state of the order. Possible values are:
	//
	// - Purchased
	// - Canceled
	// - Pending
	PurchaseState int64 `json:"purchaseState,omitempty"`

	// PurchaseTimeMillis: The time the product was purchased, in
	// milliseconds since the epoch (Jan 1, 1970).
	PurchaseTimeMillis int64 `json:"purchaseTimeMillis,omitempty,string"`

	// PurchaseType: The type of purchase of the inapp product. This field
	// is only set if this purchase was not made using the standard in-app
	// billing flow. Possible values are:
	// - Test (i.e. purchased from a license testing account)
	// - Promo (i.e. purchased using a promo code)
	// - Rewarded (i.e. from watching a video ad instead of paying)
	PurchaseType *int64 `json:"purchaseType,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ConsumptionState") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConsumptionState") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ProductPurchase) MarshalJSON() ([]byte, error) {
	type NoMethod ProductPurchase
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TokenPagination struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	PreviousPageToken string `json:"previousPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TokenPagination) MarshalJSON() ([]byte, error) {
	type NoMethod TokenPagination
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VoidedPurchase: A VoidedPurchase resource indicates a purchase that
// was either canceled/refunded/charged-back.
type VoidedPurchase struct {
	// Kind: This kind represents a voided purchase object in the
	// androidpublisher service.
	Kind string `json:"kind,omitempty"`

	// PurchaseTimeMillis: The time at which the purchase was made, in
	// milliseconds since the epoch (Jan 1, 1970).
	PurchaseTimeMillis int64 `json:"purchaseTimeMillis,omitempty,string"`

	// PurchaseToken: The token which uniquely identifies a one-time
	// purchase or subscription. To uniquely identify subscription renewals
	// use order_id (available starting from version 3 of the API).
	PurchaseToken string `json:"purchaseToken,omitempty"`

	// VoidedTimeMillis: The time at which the purchase was
	// canceled/refunded/charged-back, in milliseconds since the epoch (Jan
	// 1, 1970).
	VoidedTimeMillis int64 `json:"voidedTimeMillis,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VoidedPurchase) MarshalJSON() ([]byte, error) {
	type NoMethod VoidedPurchase
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type VoidedPurchasesListResponse struct {
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	TokenPagination *TokenPagination `json:"tokenPagination,omitempty"`

	VoidedPurchases []*VoidedPurchase `json:"voidedPurchases,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "PageInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PageInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VoidedPurchasesListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod VoidedPurchasesListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "androidpublisher.purchases.products.get":

type PurchasesProductsGetCall struct {
	s            *Service
	packageName  string
	productId    string
	token        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Checks the purchase and consumption status of an inapp item.
func (r *PurchasesProductsService) Get(packageName string, productId string, token string) *PurchasesProductsGetCall {
	c := &PurchasesProductsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.packageName = packageName
	c.productId = productId
	c.token = token
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PurchasesProductsGetCall) Fields(s ...googleapi.Field) *PurchasesProductsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PurchasesProductsGetCall) IfNoneMatch(entityTag string) *PurchasesProductsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PurchasesProductsGetCall) Context(ctx context.Context) *PurchasesProductsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PurchasesProductsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PurchasesProductsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200623")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "{packageName}/purchases/products/{productId}/tokens/{token}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"packageName": c.packageName,
		"productId":   c.productId,
		"token":       c.token,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "androidpublisher.purchases.products.get" call.
// Exactly one of *ProductPurchase or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ProductPurchase.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PurchasesProductsGetCall) Do(opts ...googleapi.CallOption) (*ProductPurchase, error) {
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
	ret := &ProductPurchase{
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
	//   "description": "Checks the purchase and consumption status of an inapp item.",
	//   "httpMethod": "GET",
	//   "id": "androidpublisher.purchases.products.get",
	//   "parameterOrder": [
	//     "packageName",
	//     "productId",
	//     "token"
	//   ],
	//   "parameters": {
	//     "packageName": {
	//       "description": "The package name of the application the inapp product was sold in (for example, 'com.some.thing').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "The inapp product SKU (for example, 'com.some.thing.inapp1').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "token": {
	//       "description": "The token provided to the user's device when the inapp product was purchased.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{packageName}/purchases/products/{productId}/tokens/{token}",
	//   "response": {
	//     "$ref": "ProductPurchase"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "androidpublisher.purchases.voidedpurchases.list":

type PurchasesVoidedpurchasesListCall struct {
	s            *Service
	packageName  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the purchases that were canceled, refunded or
// charged-back.
func (r *PurchasesVoidedpurchasesService) List(packageName string) *PurchasesVoidedpurchasesListCall {
	c := &PurchasesVoidedpurchasesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.packageName = packageName
	return c
}

// EndTime sets the optional parameter "endTime": The time, in
// milliseconds since the Epoch, of the newest voided purchase that you
// want to see in the response. The value of this parameter cannot be
// greater than the current time and is ignored if a pagination token is
// set. Default value is current time. Note: This filter is applied on
// the time at which the record is seen as voided by our systems and not
// the actual voided time returned in the response.
func (c *PurchasesVoidedpurchasesListCall) EndTime(endTime int64) *PurchasesVoidedpurchasesListCall {
	c.urlParams_.Set("endTime", fmt.Sprint(endTime))
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *PurchasesVoidedpurchasesListCall) MaxResults(maxResults int64) *PurchasesVoidedpurchasesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// StartIndex sets the optional parameter "startIndex":
func (c *PurchasesVoidedpurchasesListCall) StartIndex(startIndex int64) *PurchasesVoidedpurchasesListCall {
	c.urlParams_.Set("startIndex", fmt.Sprint(startIndex))
	return c
}

// StartTime sets the optional parameter "startTime": The time, in
// milliseconds since the Epoch, of the oldest voided purchase that you
// want to see in the response. The value of this parameter cannot be
// older than 30 days and is ignored if a pagination token is set.
// Default value is current time minus 30 days. Note: This filter is
// applied on the time at which the record is seen as voided by our
// systems and not the actual voided time returned in the response.
func (c *PurchasesVoidedpurchasesListCall) StartTime(startTime int64) *PurchasesVoidedpurchasesListCall {
	c.urlParams_.Set("startTime", fmt.Sprint(startTime))
	return c
}

// Token sets the optional parameter "token":
func (c *PurchasesVoidedpurchasesListCall) Token(token string) *PurchasesVoidedpurchasesListCall {
	c.urlParams_.Set("token", token)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PurchasesVoidedpurchasesListCall) Fields(s ...googleapi.Field) *PurchasesVoidedpurchasesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *PurchasesVoidedpurchasesListCall) IfNoneMatch(entityTag string) *PurchasesVoidedpurchasesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *PurchasesVoidedpurchasesListCall) Context(ctx context.Context) *PurchasesVoidedpurchasesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *PurchasesVoidedpurchasesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *PurchasesVoidedpurchasesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20200623")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "{packageName}/purchases/voidedpurchases")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"packageName": c.packageName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "androidpublisher.purchases.voidedpurchases.list" call.
// Exactly one of *VoidedPurchasesListResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *VoidedPurchasesListResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *PurchasesVoidedpurchasesListCall) Do(opts ...googleapi.CallOption) (*VoidedPurchasesListResponse, error) {
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
	ret := &VoidedPurchasesListResponse{
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
	//   "description": "Lists the purchases that were canceled, refunded or charged-back.",
	//   "httpMethod": "GET",
	//   "id": "androidpublisher.purchases.voidedpurchases.list",
	//   "parameterOrder": [
	//     "packageName"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "description": "The time, in milliseconds since the Epoch, of the newest voided purchase that you want to see in the response. The value of this parameter cannot be greater than the current time and is ignored if a pagination token is set. Default value is current time. Note: This filter is applied on the time at which the record is seen as voided by our systems and not the actual voided time returned in the response.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "packageName": {
	//       "description": "The package name of the application for which voided purchases need to be returned (for example, 'com.some.thing').",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "startTime": {
	//       "description": "The time, in milliseconds since the Epoch, of the oldest voided purchase that you want to see in the response. The value of this parameter cannot be older than 30 days and is ignored if a pagination token is set. Default value is current time minus 30 days. Note: This filter is applied on the time at which the record is seen as voided by our systems and not the actual voided time returned in the response.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "token": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{packageName}/purchases/voidedpurchases",
	//   "response": {
	//     "$ref": "VoidedPurchasesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

/*
NiFi Rest API

The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.

API version: 2.0.0
Contact: dev@nifi.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifiapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PoliciesAPIService PoliciesAPI service
type PoliciesAPIService service

type PoliciesAPICreateAccessPolicyRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	body *AccessPolicyEntity
}

// The access policy configuration details.
func (r PoliciesAPICreateAccessPolicyRequest) Body(body AccessPolicyEntity) PoliciesAPICreateAccessPolicyRequest {
	r.body = &body
	return r
}

func (r PoliciesAPICreateAccessPolicyRequest) Execute() (*AccessPolicyEntity, *http.Response, error) {
	return r.ApiService.CreateAccessPolicyExecute(r)
}

/*
CreateAccessPolicy Creates an access policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PoliciesAPICreateAccessPolicyRequest
*/
func (a *PoliciesAPIService) CreateAccessPolicy(ctx context.Context) PoliciesAPICreateAccessPolicyRequest {
	return PoliciesAPICreateAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccessPolicyEntity
func (a *PoliciesAPIService) CreateAccessPolicyExecute(r PoliciesAPICreateAccessPolicyRequest) (*AccessPolicyEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.CreateAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PoliciesAPIGetAccessPolicyRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	id string
}

func (r PoliciesAPIGetAccessPolicyRequest) Execute() (*AccessPolicyEntity, *http.Response, error) {
	return r.ApiService.GetAccessPolicyExecute(r)
}

/*
GetAccessPolicy Gets an access policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The access policy id.
 @return PoliciesAPIGetAccessPolicyRequest
*/
func (a *PoliciesAPIService) GetAccessPolicy(ctx context.Context, id string) PoliciesAPIGetAccessPolicyRequest {
	return PoliciesAPIGetAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccessPolicyEntity
func (a *PoliciesAPIService) GetAccessPolicyExecute(r PoliciesAPIGetAccessPolicyRequest) (*AccessPolicyEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.GetAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PoliciesAPIGetAccessPolicyForResourceRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	action string
	resource string
}

func (r PoliciesAPIGetAccessPolicyForResourceRequest) Execute() (*AccessPolicyEntity, *http.Response, error) {
	return r.ApiService.GetAccessPolicyForResourceExecute(r)
}

/*
GetAccessPolicyForResource Gets an access policy for the specified action and resource

Will return the effective policy if no component specific policy exists for the specified action and resource. Must have Read permissions to the policy with the desired action and resource. Permissions for the policy that is returned will be indicated in the response. This means the client could be authorized to get the policy for a given component but the effective policy may be inherited from an ancestor Process Group. If the client does not have permissions to that policy, the response will not include the policy and the permissions in the response will be marked accordingly. If the client does not have permissions to the policy of the desired action and resource a 403 response will be returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param action The request action.
 @param resource The resource of the policy.
 @return PoliciesAPIGetAccessPolicyForResourceRequest
*/
func (a *PoliciesAPIService) GetAccessPolicyForResource(ctx context.Context, action string, resource string) PoliciesAPIGetAccessPolicyForResourceRequest {
	return PoliciesAPIGetAccessPolicyForResourceRequest{
		ApiService: a,
		ctx: ctx,
		action: action,
		resource: resource,
	}
}

// Execute executes the request
//  @return AccessPolicyEntity
func (a *PoliciesAPIService) GetAccessPolicyForResourceExecute(r PoliciesAPIGetAccessPolicyForResourceRequest) (*AccessPolicyEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.GetAccessPolicyForResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{action}/{resource}"
	localVarPath = strings.Replace(localVarPath, "{"+"action"+"}", url.PathEscape(parameterValueToString(r.action, "action")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource"+"}", url.PathEscape(parameterValueToString(r.resource, "resource")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PoliciesAPIRemoveAccessPolicyRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	id string
	version *string
	clientId *string
	disconnectedNodeAcknowledged *bool
}

// The revision is used to verify the client is working with the latest version of the flow.
func (r PoliciesAPIRemoveAccessPolicyRequest) Version(version string) PoliciesAPIRemoveAccessPolicyRequest {
	r.version = &version
	return r
}

// If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
func (r PoliciesAPIRemoveAccessPolicyRequest) ClientId(clientId string) PoliciesAPIRemoveAccessPolicyRequest {
	r.clientId = &clientId
	return r
}

// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
func (r PoliciesAPIRemoveAccessPolicyRequest) DisconnectedNodeAcknowledged(disconnectedNodeAcknowledged bool) PoliciesAPIRemoveAccessPolicyRequest {
	r.disconnectedNodeAcknowledged = &disconnectedNodeAcknowledged
	return r
}

func (r PoliciesAPIRemoveAccessPolicyRequest) Execute() (*AccessPolicyEntity, *http.Response, error) {
	return r.ApiService.RemoveAccessPolicyExecute(r)
}

/*
RemoveAccessPolicy Deletes an access policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The access policy id.
 @return PoliciesAPIRemoveAccessPolicyRequest
*/
func (a *PoliciesAPIService) RemoveAccessPolicy(ctx context.Context, id string) PoliciesAPIRemoveAccessPolicyRequest {
	return PoliciesAPIRemoveAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccessPolicyEntity
func (a *PoliciesAPIService) RemoveAccessPolicyExecute(r PoliciesAPIRemoveAccessPolicyRequest) (*AccessPolicyEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.RemoveAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "")
	}
	if r.disconnectedNodeAcknowledged != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disconnectedNodeAcknowledged", r.disconnectedNodeAcknowledged, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PoliciesAPIUpdateAccessPolicyRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	id string
	body *AccessPolicyEntity
}

// The access policy configuration details.
func (r PoliciesAPIUpdateAccessPolicyRequest) Body(body AccessPolicyEntity) PoliciesAPIUpdateAccessPolicyRequest {
	r.body = &body
	return r
}

func (r PoliciesAPIUpdateAccessPolicyRequest) Execute() (*AccessPolicyEntity, *http.Response, error) {
	return r.ApiService.UpdateAccessPolicyExecute(r)
}

/*
UpdateAccessPolicy Updates a access policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The access policy id.
 @return PoliciesAPIUpdateAccessPolicyRequest
*/
func (a *PoliciesAPIService) UpdateAccessPolicy(ctx context.Context, id string) PoliciesAPIUpdateAccessPolicyRequest {
	return PoliciesAPIUpdateAccessPolicyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccessPolicyEntity
func (a *PoliciesAPIService) UpdateAccessPolicyExecute(r PoliciesAPIUpdateAccessPolicyRequest) (*AccessPolicyEntity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessPolicyEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.UpdateAccessPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

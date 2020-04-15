// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package client

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// Linger please
var (
	_ _context.Context
)

// ConstructionAPIService ConstructionAPI service
type ConstructionAPIService service

// ConstructionMetadata Get any information required to construct a transaction for a specific
// network. Metadata returned here could be a recent hash to use, an account sequence number, or
// even arbitrary chain state. It is up to the client to correctly populate the options object with
// any network-specific details to ensure the correct metadata is retrieved.  It is important to
// clarify that this endpoint should not pre-construct any transactions for the client (this should
// happen in the SDK). This endpoint is left purposely unstructured because of the wide scope of
// metadata that could be required.  In a future version of the spec, we plan to pass an array of
// Rosetta Operations to specify which metadata should be received and to create a transaction in an
// accompanying SDK. This will help to insulate the client from chain-specific details that are
// currently required here.
func (a *ConstructionAPIService) ConstructionMetadata(
	ctx _context.Context,
	constructionMetadataRequest *types.ConstructionMetadataRequest,
) (*types.ConstructionMetadataResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/construction/metadata"
	localVarHeaderParams := make(map[string]string)

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
	localVarPostBody = constructionMetadataRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	if localVarHTTPResponse.StatusCode != _nethttp.StatusOK {
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, err
		}

		return nil, &v, fmt.Errorf("%+v", v)
	}

	var v types.ConstructionMetadataResponse
	err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, nil, err
	}

	return &v, nil, nil
}

// ConstructionSubmit Submit a pre-signed transaction to the node. This call should not block on the
// transaction being included in a block. Rather, it should return immediately with an indication of
// whether or not the transaction was included in the mempool.  The transaction submission response
// should only return a 200 status if the submitted transaction could be included in the mempool.
// Otherwise, it should return an error.
func (a *ConstructionAPIService) ConstructionSubmit(
	ctx _context.Context,
	constructionSubmitRequest *types.ConstructionSubmitRequest,
) (*types.ConstructionSubmitResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/construction/submit"
	localVarHeaderParams := make(map[string]string)

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
	localVarPostBody = constructionSubmitRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	if localVarHTTPResponse.StatusCode != _nethttp.StatusOK {
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, err
		}

		return nil, &v, fmt.Errorf("%+v", v)
	}

	var v types.ConstructionSubmitResponse
	err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return nil, nil, err
	}

	return &v, nil, nil
}

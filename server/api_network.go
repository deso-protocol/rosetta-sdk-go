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

package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// A NetworkAPIController binds http requests to an api service and writes the service results to
// the http response
type NetworkAPIController struct {
	service NetworkAPIServicer
}

// NewNetworkAPIController creates a default api controller
func NewNetworkAPIController(s NetworkAPIServicer) Router {
	return &NetworkAPIController{service: s}
}

// Routes returns all of the api route for the NetworkAPIController
func (c *NetworkAPIController) Routes() Routes {
	return Routes{
		{
			"NetworkList",
			strings.ToUpper("Post"),
			"/network/list",
			c.NetworkList,
		},
		{
			"NetworkOptions",
			strings.ToUpper("Post"),
			"/network/options",
			c.NetworkOptions,
		},
		{
			"NetworkStatus",
			strings.ToUpper("Post"),
			"/network/status",
			c.NetworkStatus,
		},
	}
}

// NetworkList - Get List of Available Networks
func (c *NetworkAPIController) NetworkList(w http.ResponseWriter, r *http.Request) {
	metadataRequest := &types.MetadataRequest{}
	if err := json.NewDecoder(r.Body).Decode(&metadataRequest); err != nil {
		err = EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	// Assert that MetadataRequest is correct
	if err := asserter.MetadataRequest(metadataRequest); err != nil {
		err = EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	result, serviceErr := c.service.NetworkList(metadataRequest)
	if serviceErr != nil {
		err := EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := EncodeJSONResponse(result, http.StatusOK, w); err != nil {
		log.Fatal(err)
	}
}

// NetworkOptions - Get Network Options
func (c *NetworkAPIController) NetworkOptions(w http.ResponseWriter, r *http.Request) {
	networkRequest := &types.NetworkRequest{}
	if err := json.NewDecoder(r.Body).Decode(&networkRequest); err != nil {
		err = EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	// Assert that NetworkRequest is correct
	if err := asserter.NetworkRequest(networkRequest); err != nil {
		err = EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	result, serviceErr := c.service.NetworkOptions(networkRequest)
	if serviceErr != nil {
		err := EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := EncodeJSONResponse(result, http.StatusOK, w); err != nil {
		log.Fatal(err)
	}
}

// NetworkStatus - Get Network Status
func (c *NetworkAPIController) NetworkStatus(w http.ResponseWriter, r *http.Request) {
	networkRequest := &types.NetworkRequest{}
	if err := json.NewDecoder(r.Body).Decode(&networkRequest); err != nil {
		err = EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	// Assert that NetworkRequest is correct
	if err := asserter.NetworkRequest(networkRequest); err != nil {
		err = EncodeJSONResponse(&types.Error{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	result, serviceErr := c.service.NetworkStatus(networkRequest)
	if serviceErr != nil {
		err := EncodeJSONResponse(serviceErr, http.StatusInternalServerError, w)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := EncodeJSONResponse(result, http.StatusOK, w); err != nil {
		log.Fatal(err)
	}
}

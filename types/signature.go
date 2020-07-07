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

package types

// Signature Signature contains the payload that was signed, the public keys of the keypairs used to
// produce the signature, the signature (encoded in hex), and the SignatureType. PublicKey is often
// times not known during construction of the signing payloads but may be needed to combine
// signatures properly.
type Signature struct {
	SigningPayload *SigningPayload `json:"signing_payload"`
	PublicKey      *PublicKey      `json:"public_key"`
	SignatureType  SignatureType   `json:"signature_type"`
	Bytes          []byte          `json:"hex_bytes"`
}

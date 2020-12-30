/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// IncidentUpdateRequest Update request for an incident.
type IncidentUpdateRequest struct {
	Data IncidentUpdateData `json:"data"`
}

// NewIncidentUpdateRequest instantiates a new IncidentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentUpdateRequest(data IncidentUpdateData) *IncidentUpdateRequest {
	this := IncidentUpdateRequest{}
	this.Data = data
	return &this
}

// NewIncidentUpdateRequestWithDefaults instantiates a new IncidentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentUpdateRequestWithDefaults() *IncidentUpdateRequest {
	this := IncidentUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *IncidentUpdateRequest) GetData() IncidentUpdateData {
	if o == nil {
		var ret IncidentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentUpdateRequest) GetDataOk() (*IncidentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IncidentUpdateRequest) SetData(v IncidentUpdateData) {
	o.Data = v
}

func (o IncidentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableIncidentUpdateRequest struct {
	value *IncidentUpdateRequest
	isSet bool
}

func (v NullableIncidentUpdateRequest) Get() *IncidentUpdateRequest {
	return v.value
}

func (v *NullableIncidentUpdateRequest) Set(val *IncidentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentUpdateRequest(val *IncidentUpdateRequest) *NullableIncidentUpdateRequest {
	return &NullableIncidentUpdateRequest{value: val, isSet: true}
}

func (v NullableIncidentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
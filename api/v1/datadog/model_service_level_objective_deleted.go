/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// ServiceLevelObjectiveDeleted struct for ServiceLevelObjectiveDeleted
type ServiceLevelObjectiveDeleted struct {
	// An array containing the ID of the deleted service level objective object.
	Data []string `json:"data"`
}

// NewServiceLevelObjectiveDeleted instantiates a new ServiceLevelObjectiveDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectiveDeleted(data []string) *ServiceLevelObjectiveDeleted {
	this := ServiceLevelObjectiveDeleted{}
	this.Data = data
	return &this
}

// NewServiceLevelObjectiveDeletedWithDefaults instantiates a new ServiceLevelObjectiveDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectiveDeletedWithDefaults() *ServiceLevelObjectiveDeleted {
	this := ServiceLevelObjectiveDeleted{}
	return &this
}

// GetData returns the Data field value
func (o *ServiceLevelObjectiveDeleted) GetData() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *ServiceLevelObjectiveDeleted) SetData(v []string) {
	o.Data = v
}

type NullableServiceLevelObjectiveDeleted struct {
	Value        ServiceLevelObjectiveDeleted
	ExplicitNull bool
}

func (v NullableServiceLevelObjectiveDeleted) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableServiceLevelObjectiveDeleted) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
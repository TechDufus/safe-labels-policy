// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// ResourceQuotaSpec ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
//
// swagger:model ResourceQuotaSpec
type ResourceQuotaSpec struct {

	// hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard map[string]*apimachinery_pkg_api_resource.Quantity `json:"hard,omitempty"`

	// scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.
	ScopeSelector *ScopeSelector `json:"scopeSelector,omitempty"`

	// A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
	Scopes []string `json:"scopes,omitempty"`
}

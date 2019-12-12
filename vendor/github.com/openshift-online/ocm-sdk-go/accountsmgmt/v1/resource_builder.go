/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// ResourceBuilder contains the data and logic needed to build 'resource' objects.
//
// Identifies computing resources
type ResourceBuilder struct {
	id           *string
	href         *string
	link         bool
	allowed      *int
	resourceName *string
	resourceType *string
}

// NewResource creates a new builder of 'resource' objects.
func NewResource() *ResourceBuilder {
	return new(ResourceBuilder)
}

// ID sets the identifier of the object.
func (b *ResourceBuilder) ID(value string) *ResourceBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *ResourceBuilder) HREF(value string) *ResourceBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *ResourceBuilder) Link(value bool) *ResourceBuilder {
	b.link = value
	return b
}

// Allowed sets the value of the 'allowed' attribute
// to the given value.
//
//
func (b *ResourceBuilder) Allowed(value int) *ResourceBuilder {
	b.allowed = &value
	return b
}

// ResourceName sets the value of the 'resource_name' attribute
// to the given value.
//
//
func (b *ResourceBuilder) ResourceName(value string) *ResourceBuilder {
	b.resourceName = &value
	return b
}

// ResourceType sets the value of the 'resource_type' attribute
// to the given value.
//
//
func (b *ResourceBuilder) ResourceType(value string) *ResourceBuilder {
	b.resourceType = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ResourceBuilder) Copy(object *Resource) *ResourceBuilder {
	if object == nil {
		return b
	}
	b.id = object.id
	b.href = object.href
	b.link = object.link
	b.allowed = object.allowed
	b.resourceName = object.resourceName
	b.resourceType = object.resourceType
	return b
}

// Build creates a 'resource' object using the configuration stored in the builder.
func (b *ResourceBuilder) Build() (object *Resource, err error) {
	object = new(Resource)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.allowed != nil {
		object.allowed = b.allowed
	}
	if b.resourceName != nil {
		object.resourceName = b.resourceName
	}
	if b.resourceType != nil {
		object.resourceType = b.resourceType
	}
	return
}

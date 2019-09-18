// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcawsprovisioner

import (
	"encoding/json"
	"errors"

	tcclient "github.com/taskcluster/taskcluster/clients/client-go/v17"
)

type (
	// Availability zone configuration
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/availability-zone-configuration.json#
	AvailabilityZoneConfiguration struct {

		// The AWS availability zone being configured.  Example: eu-central-1b
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/availability-zone-configuration.json#/properties/availabilityZone
		AvailabilityZone string `json:"availabilityZone"`

		// LaunchSpecification entries unique to this AZ
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/availability-zone-configuration.json#/properties/launchSpec
		LaunchSpec json.RawMessage `json:"launchSpec"`

		// The AWS region containing this availability zone.  Example: eu-central-1
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/availability-zone-configuration.json#/properties/region
		Region string `json:"region"`

		// Static Secrets unique to this AZ
		//
		// Default:    {}
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/availability-zone-configuration.json#/properties/secrets
		Secrets json.RawMessage `json:"secrets,omitempty"`

		// UserData entries unique to this AZ
		//
		// Default:    {}
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/availability-zone-configuration.json#/properties/userData
		UserData json.RawMessage `json:"userData,omitempty"`
	}

	// Backend Status Response
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/backend-status-response.json#
	BackendStatusResponse struct {

		// A date when the provisioner backend process last completed an iteration.
		// This does not imply success, rather it is to make sure that the process
		// is alive
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/backend-status-response.json#/properties/lastCheckedIn
		LastCheckedIn tcclient.Time `json:"lastCheckedIn"`

		// A string from Deadman's Snitch which describes the status.  See
		// https://deadmanssnitch.com/docs/api/v1#listing-your-snitches for an
		// explanation of this value
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/backend-status-response.json#/properties/status
		Status string `json:"status"`
	}

	// A worker launchSpecification and required metadata
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#
	CreateWorkerTypeRequest struct {

		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/availabilityZones
		AvailabilityZones []AvailabilityZoneConfiguration `json:"availabilityZones,omitempty"`

		// True if this worker type is allowed on demand instances.  Currently
		// ignored
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/canUseOndemand
		CanUseOndemand bool `json:"canUseOndemand,omitempty"`

		// True if this worker type is allowed spot instances.  Currently ignored
		// as all instances are Spot
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/canUseSpot
		CanUseSpot bool `json:"canUseSpot,omitempty"`

		// A string which describes what this image is for and hints on using it
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/description
		Description string `json:"description"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/instanceTypes
		InstanceTypes []InstanceTypeConfiguration `json:"instanceTypes"`

		// Launch Specification entries which are used in all regions and all instance types
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/launchSpec
		LaunchSpec json.RawMessage `json:"launchSpec"`

		// Maximum number of capacity units to be provisioned.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/maxCapacity
		MaxCapacity float64 `json:"maxCapacity"`

		// Maximum price we'll pay.  Like minPrice, this takes into account the
		// utility factor when figuring out what the actual SpotPrice submitted
		// to Amazon will be
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/maxPrice
		MaxPrice float64 `json:"maxPrice"`

		// Minimum number of capacity units to be provisioned.  A capacity unit
		// is an abstract unit of capacity, where one capacity unit is roughly
		// one task which should be taken off the queue
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/minCapacity
		MinCapacity float64 `json:"minCapacity,omitempty"`

		// Minimum price to pay for an instance.  A Price is considered to be the
		// Amazon Spot Price multiplied by the utility factor of the InstantType
		// as specified in the instanceTypes list.  For example, if the minPrice
		// is set to $0.5 and the utility factor is 2, the actual minimum bid
		// used will be $0.25
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/minPrice
		MinPrice float64 `json:"minPrice"`

		// A string which identifies the owner of this worker type
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/owner
		Owner string `json:"owner"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/regions
		Regions []RegionConfiguration `json:"regions"`

		// A scaling ratio of `0.2` means that the provisioner will attempt to keep
		// the number of pending tasks around 20% of the provisioned capacity.
		// This results in pending tasks waiting 20% of the average task execution
		// time before starting to run.
		// A higher scaling ratio often results in better utilization and longer
		// waiting times. For workerTypes running long tasks a short scaling ratio
		// may be preferred, but for workerTypes running quick tasks a higher scaling
		// ratio may increase utilization without major delays.
		// If using a scaling ratio of 0, the provisioner will attempt to keep the
		// capacity of pending spot requests equal to the number of pending tasks.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/scalingRatio
		ScalingRatio float64 `json:"scalingRatio"`

		// Scopes to issue credentials to for all regions Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/scopes
		Scopes []string `json:"scopes"`

		// Static secrets entries which are used in all regions and all instance types
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/secrets
		Secrets json.RawMessage `json:"secrets"`

		// UserData entries which are used in all regions and all instance types
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-worker-type-request.json#/properties/userData
		UserData json.RawMessage `json:"userData"`
	}

	// Generated Temporary credentials from the Provisioner
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/credentials
	Credentials struct {

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/credentials/properties/accessToken
		AccessToken string `json:"accessToken"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/credentials/properties/certificate
		Certificate string `json:"certificate"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/credentials/properties/clientId
		ClientID string `json:"clientId"`
	}

	// Instance Type configuration
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#
	InstanceTypeConfiguration struct {

		// This number represents the number of tasks that this instance type
		// is capable of running concurrently.  This is used by the provisioner
		// to know how many pending tasks to offset a pending instance of this
		// type by
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/capacity
		Capacity float64 `json:"capacity"`

		// InstanceType name for Amazon.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/instanceType
		InstanceType string `json:"instanceType"`

		// LaunchSpecification entries unique to this InstanceType
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/launchSpec
		LaunchSpec json.RawMessage `json:"launchSpec"`

		// Scopes which should be included for this InstanceType.  Scopes must
		// be composed of printable ASCII characters and spaces.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/scopes
		Scopes []string `json:"scopes"`

		// Static Secrets unique to this InstanceType
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/secrets
		Secrets json.RawMessage `json:"secrets"`

		// UserData entries unique to this InstanceType
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/userData
		UserData json.RawMessage `json:"userData"`

		// This number is a relative measure of performance between two instance
		// types.  It is multiplied by the spot price from Amazon to figure out
		// which instance type is the cheapest one
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/instance-type-configuration.json#/properties/utility
		Utility float64 `json:"utility"`
	}

	// All of the launch specifications for a worker type
	//
	// Additional properties allowed
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/get-launch-specs-response.json#
	LaunchSpecsResponse json.RawMessage

	// See http://schemas.taskcluster.net/aws-provisioner/v1/list-worker-types-summaries-response.json#
	ListWorkerTypeSummariesResponse []WorkerTypeSummary

	// See http://schemas.taskcluster.net/aws-provisioner/v1/list-worker-types-response.json#
	ListWorkerTypes []string

	// Region configuration
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/region-configuration.json#
	RegionConfiguration struct {

		// LaunchSpecification entries unique to this Region
		//
		// Defined properties:
		//
		//  struct {
		//
		//  	// Per-region AMI ImageId
		//  	//
		//  	// See http://schemas.taskcluster.net/aws-provisioner/v1/region-launch-spec.json#/properties/ImageId
		//  	ImageID string `json:"ImageId"`
		//  }
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/region-launch-spec.json#
		LaunchSpec json.RawMessage `json:"launchSpec"`

		// The Amazon AWS Region being configured.  Example: us-west-1
		//
		// Possible values:
		//   * "us-west-2"
		//   * "us-east-1"
		//   * "us-east-2"
		//   * "us-west-1"
		//   * "eu-central-1"
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/region-configuration.json#/properties/region
		Region string `json:"region"`

		// Scopes which should be included for this Region.  Scopes must be
		// composed of printable ASCII characters and spaces.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/region-configuration.json#/properties/scopes
		Scopes []string `json:"scopes"`

		// Static Secrets unique to this Region
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/region-configuration.json#/properties/secrets
		Secrets json.RawMessage `json:"secrets"`

		// UserData entries unique to this Region
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/region-configuration.json#/properties/userData
		UserData json.RawMessage `json:"userData"`
	}

	// LaunchSpecification entries unique to this Region
	//
	// Defined properties:
	//
	//  struct {
	//
	//  	// Per-region AMI ImageId
	//  	//
	//  	// See http://schemas.taskcluster.net/aws-provisioner/v1/region-launch-spec.json#/properties/ImageId
	//  	ImageID string `json:"ImageId"`
	//  }
	//
	// Additional properties allowed
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/region-launch-spec.json#
	RegionLaunchSpec json.RawMessage

	// A Secret
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/create-secret-request.json#
	SecretRequest struct {

		// The date at which the secret is no longer guaranteed to exist
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-secret-request.json#/properties/expiration
		Expiration tcclient.Time `json:"expiration"`

		// List of strings which are scopes for temporary credentials to give
		// to the worker through the secret system.  Scopes must be composed of
		// printable ASCII characters and spaces.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-secret-request.json#/properties/scopes
		Scopes []string `json:"scopes"`

		// Free form object which contains the secrets stored
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-secret-request.json#/properties/secrets
		Secrets json.RawMessage `json:"secrets"`

		// A Slug ID which is the uniquely addressable token to access this
		// set of secrets
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-secret-request.json#/properties/token
		Token string `json:"token"`

		// A string describing what the secret will be used for
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/create-secret-request.json#/properties/workerType
		WorkerType string `json:"workerType"`
	}

	// Secrets from the provisioner
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#
	SecretResponse struct {

		// Generated Temporary credentials from the Provisioner
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/credentials
		Credentials Credentials `json:"credentials"`

		// Free-form object which contains secrets from the worker type definition
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/data
		Data json.RawMessage `json:"data"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-secret-response.json#/properties/scopes
		Scopes []string `json:"scopes"`
	}

	// The last modified date of a workerType
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-last-modified.json#
	WorkerTypeLastModified struct {

		// ISO Date string (e.g. new Date().toISOString()) which represents the time
		// when this worker type definition was last altered (inclusive of creation)
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-last-modified.json#/properties/lastModified
		LastModified tcclient.Time `json:"lastModified"`

		// The ID of the workerType
		//
		// Syntax:     ^[A-Za-z0-9+/=_-]{1,22}$
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-last-modified.json#/properties/workerType
		WorkerType string `json:"workerType"`
	}

	// A worker launchSpecification and required metadata
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#
	WorkerTypeResponse struct {

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/availabilityZones
		AvailabilityZones []AvailabilityZoneConfiguration `json:"availabilityZones,omitempty"`

		// True if this worker type is allowed on demand instances.  Currently
		// ignored
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/canUseOndemand
		CanUseOndemand bool `json:"canUseOndemand,omitempty"`

		// True if this worker type is allowed spot instances.  Currently ignored
		// as all instances are Spot
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/canUseSpot
		CanUseSpot bool `json:"canUseSpot,omitempty"`

		// A string which describes what this image is for and hints on using it
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/description
		Description string `json:"description"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/instanceTypes
		InstanceTypes []InstanceTypeConfiguration `json:"instanceTypes"`

		// ISO Date string (e.g. new Date().toISOString()) which represents the time
		// when this worker type definition was last altered (inclusive of creation)
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/lastModified
		LastModified tcclient.Time `json:"lastModified"`

		// Launch Specification entries which are used in all regions and all instance types
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/launchSpec
		LaunchSpec json.RawMessage `json:"launchSpec"`

		// Maximum number of capacity units to be provisioned.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/maxCapacity
		MaxCapacity float64 `json:"maxCapacity"`

		// Maximum price we'll pay.  Like minPrice, this takes into account the
		// utility factor when figuring out what the actual SpotPrice submitted
		// to Amazon will be
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/maxPrice
		MaxPrice float64 `json:"maxPrice"`

		// Minimum number of capacity units to be provisioned.  A capacity unit
		// is an abstract unit of capacity, where one capacity unit is roughly
		// one task which should be taken off the queue
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/minCapacity
		MinCapacity float64 `json:"minCapacity"`

		// Minimum price to pay for an instance.  A Price is considered to be the
		// Amazon Spot Price multiplied by the utility factor of the InstantType
		// as specified in the instanceTypes list.  For example, if the minPrice
		// is set to $0.5 and the utility factor is 2, the actual minimum bid
		// used will be $0.25
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/minPrice
		MinPrice float64 `json:"minPrice"`

		// A string which identifies the owner of this worker type
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/owner
		Owner string `json:"owner"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/regions
		Regions []RegionConfiguration `json:"regions"`

		// A scaling ratio of `0.2` means that the provisioner will attempt to keep
		// the number of pending tasks around 20% of the provisioned capacity.
		// This results in pending tasks waiting 20% of the average task execution
		// time before starting to run.
		// A higher scaling ratio often results in better utilization and longer
		// waiting times. For workerTypes running long tasks a short scaling ratio
		// may be preferred, but for workerTypes running quick tasks a higher scaling
		// ratio may increase utilization without major delays.
		// If using a scaling ratio of 0, the provisioner will attempt to keep the
		// capacity of pending spot requests equal to the number of pending tasks.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/scalingRatio
		ScalingRatio float64 `json:"scalingRatio"`

		// Scopes to issue credentials to for all regions.  Scopes must be composed
		// of printable ASCII characters and spaces.
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/scopes
		Scopes []string `json:"scopes"`

		// Static secrets entries which are used in all regions and all instance types
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/secrets
		Secrets json.RawMessage `json:"secrets"`

		// UserData entries which are used in all regions and all instance types
		//
		// Additional properties allowed
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/userData
		UserData json.RawMessage `json:"userData"`

		// The ID of the workerType
		//
		// Syntax:     ^[A-Za-z0-9+/=_-]{1,22}$
		//
		// See http://schemas.taskcluster.net/aws-provisioner/v1/get-worker-type-response.json#/properties/workerType
		WorkerType string `json:"workerType"`
	}

	// A summary of a worker type's current state, expresed in terms of capacity.
	//
	// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#
	WorkerTypeSummary struct {

		// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#/properties/maxCapacity
		MaxCapacity int64 `json:"maxCapacity"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#/properties/minCapacity
		MinCapacity int64 `json:"minCapacity"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#/properties/pendingCapacity
		PendingCapacity int64 `json:"pendingCapacity"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#/properties/requestedCapacity
		RequestedCapacity int64 `json:"requestedCapacity"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#/properties/runningCapacity
		RunningCapacity int64 `json:"runningCapacity"`

		// See http://schemas.taskcluster.net/aws-provisioner/v1/worker-type-summary.json#/properties/workerType
		WorkerType string `json:"workerType"`
	}
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// LaunchSpecsResponse is of type json.RawMessage...
func (this *LaunchSpecsResponse) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *LaunchSpecsResponse) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("LaunchSpecsResponse: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

// MarshalJSON calls json.RawMessage method of the same name. Required since
// RegionLaunchSpec is of type json.RawMessage...
func (this *RegionLaunchSpec) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *RegionLaunchSpec) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("RegionLaunchSpec: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}

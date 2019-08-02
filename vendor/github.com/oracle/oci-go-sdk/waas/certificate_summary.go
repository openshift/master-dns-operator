// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CertificateSummary A summary of the SSL certificate's information.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CertificateSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SSL certificate.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SSL certificate's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name of the SSL certificate.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
	TimeNotValidAfter *common.SDKTime `mandatory:"false" json:"timeNotValidAfter"`

	// A simple key-value pair without any defined schema.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A key-value pair with a defined schema that restricts the values of tags. These predefined keys are scoped to namespaces.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current lifecycle state of the certificate.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the certificate was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m CertificateSummary) String() string {
	return common.PointerString(m)
}

// CertificateSummaryLifecycleStateEnum is an alias to type: LifecycleStatesEnum
// Consider using LifecycleStatesEnum instead
// Deprecated
type CertificateSummaryLifecycleStateEnum = LifecycleStatesEnum

// Set of constants representing the allowable values for LifecycleStatesEnum
// Deprecated
const (
	CertificateSummaryLifecycleStateCreating LifecycleStatesEnum = "CREATING"
	CertificateSummaryLifecycleStateActive   LifecycleStatesEnum = "ACTIVE"
	CertificateSummaryLifecycleStateFailed   LifecycleStatesEnum = "FAILED"
	CertificateSummaryLifecycleStateUpdating LifecycleStatesEnum = "UPDATING"
	CertificateSummaryLifecycleStateDeleting LifecycleStatesEnum = "DELETING"
	CertificateSummaryLifecycleStateDeleted  LifecycleStatesEnum = "DELETED"
)

// GetCertificateSummaryLifecycleStateEnumValues Enumerates the set of values for LifecycleStatesEnum
// Consider using GetLifecycleStatesEnumValue
// Deprecated
var GetCertificateSummaryLifecycleStateEnumValues = GetLifecycleStatesEnumValues
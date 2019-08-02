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

// Certificate The details of the SSL certificate.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Certificate struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the certificate.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the certificate's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name of the certificate.
	DisplayName *string `mandatory:"false" json:"displayName"`

	IssuedBy *string `mandatory:"false" json:"issuedBy"`

	SubjectName *CertificateSubjectName `mandatory:"false" json:"subjectName"`

	IssuerName *CertificateSubjectName `mandatory:"false" json:"issuerName"`

	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	Version *int `mandatory:"false" json:"version"`

	SignatureAlgorithm *string `mandatory:"false" json:"signatureAlgorithm"`

	TimeNotValidBefore *common.SDKTime `mandatory:"false" json:"timeNotValidBefore"`

	// The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
	TimeNotValidAfter *common.SDKTime `mandatory:"false" json:"timeNotValidAfter"`

	PublicKeyInfo *CertificatePublicKeyInfo `mandatory:"false" json:"publicKeyInfo"`

	Extensions []CertificateExtensions `mandatory:"false" json:"extensions"`

	// A simple key-value pair without any defined schema.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A key-value pair with a defined schema that restricts the values of tags. These predefined keys are scoped to namespaces.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current lifecycle state of the SSL certificate.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the certificate was created, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Certificate) String() string {
	return common.PointerString(m)
}

// CertificateLifecycleStateEnum is an alias to type: LifecycleStatesEnum
// Consider using LifecycleStatesEnum instead
// Deprecated
type CertificateLifecycleStateEnum = LifecycleStatesEnum

// Set of constants representing the allowable values for LifecycleStatesEnum
// Deprecated
const (
	CertificateLifecycleStateCreating LifecycleStatesEnum = "CREATING"
	CertificateLifecycleStateActive   LifecycleStatesEnum = "ACTIVE"
	CertificateLifecycleStateFailed   LifecycleStatesEnum = "FAILED"
	CertificateLifecycleStateUpdating LifecycleStatesEnum = "UPDATING"
	CertificateLifecycleStateDeleting LifecycleStatesEnum = "DELETING"
	CertificateLifecycleStateDeleted  LifecycleStatesEnum = "DELETED"
)

// GetCertificateLifecycleStateEnumValues Enumerates the set of values for LifecycleStatesEnum
// Consider using GetLifecycleStatesEnumValue
// Deprecated
var GetCertificateLifecycleStateEnumValues = GetLifecycleStatesEnumValues

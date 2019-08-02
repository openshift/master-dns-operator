package storsimple8000series

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BackupSchedulesClient is the client for the BackupSchedules methods of the Storsimple8000series service.
type BackupSchedulesClient struct {
	ManagementClient
}

// NewBackupSchedulesClient creates an instance of the BackupSchedulesClient client.
func NewBackupSchedulesClient(subscriptionID string) BackupSchedulesClient {
	return NewBackupSchedulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupSchedulesClientWithBaseURI creates an instance of the BackupSchedulesClient client.
func NewBackupSchedulesClientWithBaseURI(baseURI string, subscriptionID string) BackupSchedulesClient {
	return BackupSchedulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the backup schedule. This method may poll for completion. Polling can be canceled
// by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// deviceName is the device name backupPolicyName is the backup policy name. backupScheduleName is the backup schedule
// name. parameters is the backup schedule. resourceGroupName is the resource group name managerName is the manager
// name
func (client BackupSchedulesClient) CreateOrUpdate(deviceName string, backupPolicyName string, backupScheduleName string, parameters BackupSchedule, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan BackupSchedule, <-chan error) {
	resultChan := make(chan BackupSchedule, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BackupScheduleProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.BackupScheduleProperties.ScheduleRecurrence", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.BackupScheduleProperties.ScheduleRecurrence.RecurrenceValue", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.BackupScheduleProperties.RetentionCount", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.BackupScheduleProperties.StartTime", Name: validation.Null, Rule: true, Chain: nil},
				}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple8000series.BackupSchedulesClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result BackupSchedule
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(deviceName, backupPolicyName, backupScheduleName, parameters, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupSchedulesClient) CreateOrUpdatePreparer(deviceName string, backupPolicyName string, backupScheduleName string, parameters BackupSchedule, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":   backupPolicyName,
		"backupScheduleName": backupScheduleName,
		"deviceName":         deviceName,
		"managerName":        managerName,
		"resourceGroupName":  resourceGroupName,
		"subscriptionId":     client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupSchedulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupSchedulesClient) CreateOrUpdateResponder(resp *http.Response) (result BackupSchedule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the backup schedule. This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// deviceName is the device name backupPolicyName is the backup policy name. backupScheduleName is the name the backup
// schedule. resourceGroupName is the resource group name managerName is the manager name
func (client BackupSchedulesClient) Delete(deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple8000series.BackupSchedulesClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client BackupSchedulesClient) DeletePreparer(deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":   backupPolicyName,
		"backupScheduleName": backupScheduleName,
		"deviceName":         deviceName,
		"managerName":        managerName,
		"resourceGroupName":  resourceGroupName,
		"subscriptionId":     client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupSchedulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupSchedulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified backup schedule name.
//
// deviceName is the device name backupPolicyName is the backup policy name. backupScheduleName is the name of the
// backup schedule to be fetched resourceGroupName is the resource group name managerName is the manager name
func (client BackupSchedulesClient) Get(deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string) (result BackupSchedule, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.BackupSchedulesClient", "Get")
	}

	req, err := client.GetPreparer(deviceName, backupPolicyName, backupScheduleName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupSchedulesClient) GetPreparer(deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":   backupPolicyName,
		"backupScheduleName": backupScheduleName,
		"deviceName":         deviceName,
		"managerName":        managerName,
		"resourceGroupName":  resourceGroupName,
		"subscriptionId":     client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupSchedulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupSchedulesClient) GetResponder(resp *http.Response) (result BackupSchedule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBackupPolicy gets all the backup schedules in a backup policy.
//
// deviceName is the device name backupPolicyName is the backup policy name. resourceGroupName is the resource group
// name managerName is the manager name
func (client BackupSchedulesClient) ListByBackupPolicy(deviceName string, backupPolicyName string, resourceGroupName string, managerName string) (result BackupScheduleList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple8000series.BackupSchedulesClient", "ListByBackupPolicy")
	}

	req, err := client.ListByBackupPolicyPreparer(deviceName, backupPolicyName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "ListByBackupPolicy", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBackupPolicySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "ListByBackupPolicy", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBackupPolicyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple8000series.BackupSchedulesClient", "ListByBackupPolicy", resp, "Failure responding to request")
	}

	return
}

// ListByBackupPolicyPreparer prepares the ListByBackupPolicy request.
func (client BackupSchedulesClient) ListByBackupPolicyPreparer(deviceName string, backupPolicyName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":  backupPolicyName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByBackupPolicySender sends the ListByBackupPolicy request. The method will close the
// http.Response Body if it receives an error.
func (client BackupSchedulesClient) ListByBackupPolicySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByBackupPolicyResponder handles the response to the ListByBackupPolicy request. The method always
// closes the http.Response Body.
func (client BackupSchedulesClient) ListByBackupPolicyResponder(resp *http.Response) (result BackupScheduleList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

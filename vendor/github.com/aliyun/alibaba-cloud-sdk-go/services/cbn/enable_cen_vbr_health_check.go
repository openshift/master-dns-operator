package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// EnableCenVbrHealthCheck invokes the cbn.EnableCenVbrHealthCheck API synchronously
// api document: https://help.aliyun.com/api/cbn/enablecenvbrhealthcheck.html
func (client *Client) EnableCenVbrHealthCheck(request *EnableCenVbrHealthCheckRequest) (response *EnableCenVbrHealthCheckResponse, err error) {
	response = CreateEnableCenVbrHealthCheckResponse()
	err = client.DoAction(request, response)
	return
}

// EnableCenVbrHealthCheckWithChan invokes the cbn.EnableCenVbrHealthCheck API asynchronously
// api document: https://help.aliyun.com/api/cbn/enablecenvbrhealthcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableCenVbrHealthCheckWithChan(request *EnableCenVbrHealthCheckRequest) (<-chan *EnableCenVbrHealthCheckResponse, <-chan error) {
	responseChan := make(chan *EnableCenVbrHealthCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableCenVbrHealthCheck(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// EnableCenVbrHealthCheckWithCallback invokes the cbn.EnableCenVbrHealthCheck API asynchronously
// api document: https://help.aliyun.com/api/cbn/enablecenvbrhealthcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableCenVbrHealthCheckWithCallback(request *EnableCenVbrHealthCheckRequest, callback func(response *EnableCenVbrHealthCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableCenVbrHealthCheckResponse
		var err error
		defer close(result)
		response, err = client.EnableCenVbrHealthCheck(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// EnableCenVbrHealthCheckRequest is the request struct for api EnableCenVbrHealthCheck
type EnableCenVbrHealthCheckRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string           `position:"Query" name:"CenId"`
	HealthCheckSourceIp  string           `position:"Query" name:"HealthCheckSourceIp"`
	VbrInstanceOwnerId   requests.Integer `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string           `position:"Query" name:"VbrInstanceId"`
	HealthCheckTargetIp  string           `position:"Query" name:"HealthCheckTargetIp"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string           `position:"Query" name:"VbrInstanceRegionId"`
}

// EnableCenVbrHealthCheckResponse is the response struct for api EnableCenVbrHealthCheck
type EnableCenVbrHealthCheckResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableCenVbrHealthCheckRequest creates a request to invoke EnableCenVbrHealthCheck API
func CreateEnableCenVbrHealthCheckRequest() (request *EnableCenVbrHealthCheckRequest) {
	request = &EnableCenVbrHealthCheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "EnableCenVbrHealthCheck", "cbn", "openAPI")
	return
}

// CreateEnableCenVbrHealthCheckResponse creates a response to parse from EnableCenVbrHealthCheck response
func CreateEnableCenVbrHealthCheckResponse() (response *EnableCenVbrHealthCheckResponse) {
	response = &EnableCenVbrHealthCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

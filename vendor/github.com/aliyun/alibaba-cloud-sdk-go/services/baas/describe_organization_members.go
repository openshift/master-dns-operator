package baas

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

// DescribeOrganizationMembers invokes the baas.DescribeOrganizationMembers API synchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationmembers.html
func (client *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
	response = CreateDescribeOrganizationMembersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOrganizationMembersWithChan invokes the baas.DescribeOrganizationMembers API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationmembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationMembersWithChan(request *DescribeOrganizationMembersRequest) (<-chan *DescribeOrganizationMembersResponse, <-chan error) {
	responseChan := make(chan *DescribeOrganizationMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOrganizationMembers(request)
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

// DescribeOrganizationMembersWithCallback invokes the baas.DescribeOrganizationMembers API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationmembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationMembersWithCallback(request *DescribeOrganizationMembersRequest, callback func(response *DescribeOrganizationMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOrganizationMembersResponse
		var err error
		defer close(result)
		response, err = client.DescribeOrganizationMembers(request)
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

// DescribeOrganizationMembersRequest is the request struct for api DescribeOrganizationMembers
type DescribeOrganizationMembersRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Body" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
}

// DescribeOrganizationMembersResponse is the response struct for api DescribeOrganizationMembers
type DescribeOrganizationMembersResponse struct {
	*responses.BaseResponse
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	Success   bool                          `json:"Success" xml:"Success"`
	ErrorCode int                           `json:"ErrorCode" xml:"ErrorCode"`
	Result    []DescribeOrganizationMembers `json:"Result" xml:"Result"`
}

// CreateDescribeOrganizationMembersRequest creates a request to invoke DescribeOrganizationMembers API
func CreateDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
	request = &DescribeOrganizationMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "DescribeOrganizationMembers", "", "")
	return
}

// CreateDescribeOrganizationMembersResponse creates a response to parse from DescribeOrganizationMembers response
func CreateDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
	response = &DescribeOrganizationMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

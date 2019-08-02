package imm

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

// ListTagSets invokes the imm.ListTagSets API synchronously
// api document: https://help.aliyun.com/api/imm/listtagsets.html
func (client *Client) ListTagSets(request *ListTagSetsRequest) (response *ListTagSetsResponse, err error) {
	response = CreateListTagSetsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagSetsWithChan invokes the imm.ListTagSets API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagsets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagSetsWithChan(request *ListTagSetsRequest) (<-chan *ListTagSetsResponse, <-chan error) {
	responseChan := make(chan *ListTagSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagSets(request)
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

// ListTagSetsWithCallback invokes the imm.ListTagSets API asynchronously
// api document: https://help.aliyun.com/api/imm/listtagsets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTagSetsWithCallback(request *ListTagSetsRequest, callback func(response *ListTagSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagSetsResponse
		var err error
		defer close(result)
		response, err = client.ListTagSets(request)
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

// ListTagSetsRequest is the request struct for api ListTagSets
type ListTagSetsRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Marker  string           `position:"Query" name:"Marker"`
	Project string           `position:"Query" name:"Project"`
}

// ListTagSetsResponse is the response struct for api ListTagSets
type ListTagSetsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	NextMarker string     `json:"NextMarker" xml:"NextMarker"`
	Sets       []SetsItem `json:"Sets" xml:"Sets"`
}

// CreateListTagSetsRequest creates a request to invoke ListTagSets API
func CreateListTagSetsRequest() (request *ListTagSetsRequest) {
	request = &ListTagSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListTagSets", "imm", "openAPI")
	return
}

// CreateListTagSetsResponse creates a response to parse from ListTagSets response
func CreateListTagSetsResponse() (response *ListTagSetsResponse) {
	response = &ListTagSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

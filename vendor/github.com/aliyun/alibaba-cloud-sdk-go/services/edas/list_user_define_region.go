package edas

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

// ListUserDefineRegion invokes the edas.ListUserDefineRegion API synchronously
// api document: https://help.aliyun.com/api/edas/listuserdefineregion.html
func (client *Client) ListUserDefineRegion(request *ListUserDefineRegionRequest) (response *ListUserDefineRegionResponse, err error) {
	response = CreateListUserDefineRegionResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserDefineRegionWithChan invokes the edas.ListUserDefineRegion API asynchronously
// api document: https://help.aliyun.com/api/edas/listuserdefineregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUserDefineRegionWithChan(request *ListUserDefineRegionRequest) (<-chan *ListUserDefineRegionResponse, <-chan error) {
	responseChan := make(chan *ListUserDefineRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserDefineRegion(request)
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

// ListUserDefineRegionWithCallback invokes the edas.ListUserDefineRegion API asynchronously
// api document: https://help.aliyun.com/api/edas/listuserdefineregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUserDefineRegionWithCallback(request *ListUserDefineRegionRequest, callback func(response *ListUserDefineRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserDefineRegionResponse
		var err error
		defer close(result)
		response, err = client.ListUserDefineRegion(request)
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

// ListUserDefineRegionRequest is the request struct for api ListUserDefineRegion
type ListUserDefineRegionRequest struct {
	*requests.RoaRequest
	DebugEnable requests.Boolean `position:"Query" name:"DebugEnable"`
}

// ListUserDefineRegionResponse is the response struct for api ListUserDefineRegion
type ListUserDefineRegionResponse struct {
	*responses.BaseResponse
	Code                 int                  `json:"Code" xml:"Code"`
	Message              string               `json:"Message" xml:"Message"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	UserDefineRegionList UserDefineRegionList `json:"UserDefineRegionList" xml:"UserDefineRegionList"`
}

// CreateListUserDefineRegionRequest creates a request to invoke ListUserDefineRegion API
func CreateListUserDefineRegionRequest() (request *ListUserDefineRegionRequest) {
	request = &ListUserDefineRegionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListUserDefineRegion", "/pop/v5/user_region_defs", "", "")
	request.Method = requests.POST
	return
}

// CreateListUserDefineRegionResponse creates a response to parse from ListUserDefineRegion response
func CreateListUserDefineRegionResponse() (response *ListUserDefineRegionResponse) {
	response = &ListUserDefineRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

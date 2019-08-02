package push

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

// BindAlias invokes the push.BindAlias API synchronously
// api document: https://help.aliyun.com/api/push/bindalias.html
func (client *Client) BindAlias(request *BindAliasRequest) (response *BindAliasResponse, err error) {
	response = CreateBindAliasResponse()
	err = client.DoAction(request, response)
	return
}

// BindAliasWithChan invokes the push.BindAlias API asynchronously
// api document: https://help.aliyun.com/api/push/bindalias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindAliasWithChan(request *BindAliasRequest) (<-chan *BindAliasResponse, <-chan error) {
	responseChan := make(chan *BindAliasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindAlias(request)
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

// BindAliasWithCallback invokes the push.BindAlias API asynchronously
// api document: https://help.aliyun.com/api/push/bindalias.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindAliasWithCallback(request *BindAliasRequest, callback func(response *BindAliasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindAliasResponse
		var err error
		defer close(result)
		response, err = client.BindAlias(request)
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

// BindAliasRequest is the request struct for api BindAlias
type BindAliasRequest struct {
	*requests.RpcRequest
	AliasName string           `position:"Query" name:"AliasName"`
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	DeviceId  string           `position:"Query" name:"DeviceId"`
}

// BindAliasResponse is the response struct for api BindAlias
type BindAliasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindAliasRequest creates a request to invoke BindAlias API
func CreateBindAliasRequest() (request *BindAliasRequest) {
	request = &BindAliasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "BindAlias", "push", "openAPI")
	return
}

// CreateBindAliasResponse creates a response to parse from BindAlias response
func CreateBindAliasResponse() (response *BindAliasResponse) {
	response = &BindAliasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

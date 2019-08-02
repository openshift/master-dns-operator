package dms_enterprise

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

// RegisterUser invokes the dms_enterprise.RegisterUser API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/registeruser.html
func (client *Client) RegisterUser(request *RegisterUserRequest) (response *RegisterUserResponse, err error) {
	response = CreateRegisterUserResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterUserWithChan invokes the dms_enterprise.RegisterUser API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/registeruser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterUserWithChan(request *RegisterUserRequest) (<-chan *RegisterUserResponse, <-chan error) {
	responseChan := make(chan *RegisterUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterUser(request)
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

// RegisterUserWithCallback invokes the dms_enterprise.RegisterUser API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/registeruser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterUserWithCallback(request *RegisterUserRequest, callback func(response *RegisterUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterUserResponse
		var err error
		defer close(result)
		response, err = client.RegisterUser(request)
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

// RegisterUserRequest is the request struct for api RegisterUser
type RegisterUserRequest struct {
	*requests.RpcRequest
	RoleNames string           `position:"Query" name:"RoleNames"`
	Uid       requests.Integer `position:"Query" name:"Uid"`
	UserNick  string           `position:"Query" name:"UserNick"`
	Tid       requests.Integer `position:"Query" name:"Tid"`
}

// RegisterUserResponse is the response struct for api RegisterUser
type RegisterUserResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateRegisterUserRequest creates a request to invoke RegisterUser API
func CreateRegisterUserRequest() (request *RegisterUserRequest) {
	request = &RegisterUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "RegisterUser", "dmsenterprise", "openAPI")
	return
}

// CreateRegisterUserResponse creates a response to parse from RegisterUser response
func CreateRegisterUserResponse() (response *RegisterUserResponse) {
	response = &RegisterUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

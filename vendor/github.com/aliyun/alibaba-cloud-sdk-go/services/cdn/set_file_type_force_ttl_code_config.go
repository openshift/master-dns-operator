package cdn

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

// SetFileTypeForceTtlCodeConfig invokes the cdn.SetFileTypeForceTtlCodeConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setfiletypeforcettlcodeconfig.html
func (client *Client) SetFileTypeForceTtlCodeConfig(request *SetFileTypeForceTtlCodeConfigRequest) (response *SetFileTypeForceTtlCodeConfigResponse, err error) {
	response = CreateSetFileTypeForceTtlCodeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetFileTypeForceTtlCodeConfigWithChan invokes the cdn.SetFileTypeForceTtlCodeConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setfiletypeforcettlcodeconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetFileTypeForceTtlCodeConfigWithChan(request *SetFileTypeForceTtlCodeConfigRequest) (<-chan *SetFileTypeForceTtlCodeConfigResponse, <-chan error) {
	responseChan := make(chan *SetFileTypeForceTtlCodeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetFileTypeForceTtlCodeConfig(request)
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

// SetFileTypeForceTtlCodeConfigWithCallback invokes the cdn.SetFileTypeForceTtlCodeConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setfiletypeforcettlcodeconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetFileTypeForceTtlCodeConfigWithCallback(request *SetFileTypeForceTtlCodeConfigRequest, callback func(response *SetFileTypeForceTtlCodeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetFileTypeForceTtlCodeConfigResponse
		var err error
		defer close(result)
		response, err = client.SetFileTypeForceTtlCodeConfig(request)
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

// SetFileTypeForceTtlCodeConfigRequest is the request struct for api SetFileTypeForceTtlCodeConfig
type SetFileTypeForceTtlCodeConfigRequest struct {
	*requests.RpcRequest
	FileType      string           `position:"Query" name:"FileType"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Version       string           `position:"Query" name:"Version"`
	CodeString    string           `position:"Query" name:"CodeString"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// SetFileTypeForceTtlCodeConfigResponse is the response struct for api SetFileTypeForceTtlCodeConfig
type SetFileTypeForceTtlCodeConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetFileTypeForceTtlCodeConfigRequest creates a request to invoke SetFileTypeForceTtlCodeConfig API
func CreateSetFileTypeForceTtlCodeConfigRequest() (request *SetFileTypeForceTtlCodeConfigRequest) {
	request = &SetFileTypeForceTtlCodeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetFileTypeForceTtlCodeConfig", "", "")
	return
}

// CreateSetFileTypeForceTtlCodeConfigResponse creates a response to parse from SetFileTypeForceTtlCodeConfig response
func CreateSetFileTypeForceTtlCodeConfigResponse() (response *SetFileTypeForceTtlCodeConfigResponse) {
	response = &SetFileTypeForceTtlCodeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
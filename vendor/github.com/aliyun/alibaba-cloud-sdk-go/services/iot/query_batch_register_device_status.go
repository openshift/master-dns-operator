package iot

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

// QueryBatchRegisterDeviceStatus invokes the iot.QueryBatchRegisterDeviceStatus API synchronously
// api document: https://help.aliyun.com/api/iot/querybatchregisterdevicestatus.html
func (client *Client) QueryBatchRegisterDeviceStatus(request *QueryBatchRegisterDeviceStatusRequest) (response *QueryBatchRegisterDeviceStatusResponse, err error) {
	response = CreateQueryBatchRegisterDeviceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBatchRegisterDeviceStatusWithChan invokes the iot.QueryBatchRegisterDeviceStatus API asynchronously
// api document: https://help.aliyun.com/api/iot/querybatchregisterdevicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBatchRegisterDeviceStatusWithChan(request *QueryBatchRegisterDeviceStatusRequest) (<-chan *QueryBatchRegisterDeviceStatusResponse, <-chan error) {
	responseChan := make(chan *QueryBatchRegisterDeviceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBatchRegisterDeviceStatus(request)
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

// QueryBatchRegisterDeviceStatusWithCallback invokes the iot.QueryBatchRegisterDeviceStatus API asynchronously
// api document: https://help.aliyun.com/api/iot/querybatchregisterdevicestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBatchRegisterDeviceStatusWithCallback(request *QueryBatchRegisterDeviceStatusRequest, callback func(response *QueryBatchRegisterDeviceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBatchRegisterDeviceStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryBatchRegisterDeviceStatus(request)
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

// QueryBatchRegisterDeviceStatusRequest is the request struct for api QueryBatchRegisterDeviceStatus
type QueryBatchRegisterDeviceStatusRequest struct {
	*requests.RpcRequest
	ApplyId       requests.Integer `position:"Query" name:"ApplyId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
}

// QueryBatchRegisterDeviceStatusResponse is the response struct for api QueryBatchRegisterDeviceStatus
type QueryBatchRegisterDeviceStatusResponse struct {
	*responses.BaseResponse
	RequestId    string                               `json:"RequestId" xml:"RequestId"`
	Success      bool                                 `json:"Success" xml:"Success"`
	Code         string                               `json:"Code" xml:"Code"`
	ErrorMessage string                               `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryBatchRegisterDeviceStatus `json:"Data" xml:"Data"`
}

// CreateQueryBatchRegisterDeviceStatusRequest creates a request to invoke QueryBatchRegisterDeviceStatus API
func CreateQueryBatchRegisterDeviceStatusRequest() (request *QueryBatchRegisterDeviceStatusRequest) {
	request = &QueryBatchRegisterDeviceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryBatchRegisterDeviceStatus", "iot", "openAPI")
	return
}

// CreateQueryBatchRegisterDeviceStatusResponse creates a response to parse from QueryBatchRegisterDeviceStatus response
func CreateQueryBatchRegisterDeviceStatusResponse() (response *QueryBatchRegisterDeviceStatusResponse) {
	response = &QueryBatchRegisterDeviceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// QueryDeviceListByDeviceGroup invokes the iot.QueryDeviceListByDeviceGroup API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicelistbydevicegroup.html
func (client *Client) QueryDeviceListByDeviceGroup(request *QueryDeviceListByDeviceGroupRequest) (response *QueryDeviceListByDeviceGroupResponse, err error) {
	response = CreateQueryDeviceListByDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceListByDeviceGroupWithChan invokes the iot.QueryDeviceListByDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicelistbydevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceListByDeviceGroupWithChan(request *QueryDeviceListByDeviceGroupRequest) (<-chan *QueryDeviceListByDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceListByDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceListByDeviceGroup(request)
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

// QueryDeviceListByDeviceGroupWithCallback invokes the iot.QueryDeviceListByDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicelistbydevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceListByDeviceGroupWithCallback(request *QueryDeviceListByDeviceGroupRequest, callback func(response *QueryDeviceListByDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceListByDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceListByDeviceGroup(request)
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

// QueryDeviceListByDeviceGroupRequest is the request struct for api QueryDeviceListByDeviceGroup
type QueryDeviceListByDeviceGroupRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	GroupId       string           `position:"Query" name:"GroupId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
}

// QueryDeviceListByDeviceGroupResponse is the response struct for api QueryDeviceListByDeviceGroup
type QueryDeviceListByDeviceGroupResponse struct {
	*responses.BaseResponse
	RequestId    string                             `json:"RequestId" xml:"RequestId"`
	Success      bool                               `json:"Success" xml:"Success"`
	Code         string                             `json:"Code" xml:"Code"`
	ErrorMessage string                             `json:"ErrorMessage" xml:"ErrorMessage"`
	Page         int                                `json:"Page" xml:"Page"`
	PageSize     int                                `json:"PageSize" xml:"PageSize"`
	PageCount    int                                `json:"PageCount" xml:"PageCount"`
	Total        int                                `json:"Total" xml:"Total"`
	Data         DataInQueryDeviceListByDeviceGroup `json:"Data" xml:"Data"`
}

// CreateQueryDeviceListByDeviceGroupRequest creates a request to invoke QueryDeviceListByDeviceGroup API
func CreateQueryDeviceListByDeviceGroupRequest() (request *QueryDeviceListByDeviceGroupRequest) {
	request = &QueryDeviceListByDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceListByDeviceGroup", "iot", "openAPI")
	return
}

// CreateQueryDeviceListByDeviceGroupResponse creates a response to parse from QueryDeviceListByDeviceGroup response
func CreateQueryDeviceListByDeviceGroupResponse() (response *QueryDeviceListByDeviceGroupResponse) {
	response = &QueryDeviceListByDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

package smartag

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

// DescribeUserOnlineClientStatistics invokes the smartag.DescribeUserOnlineClientStatistics API synchronously
// api document: https://help.aliyun.com/api/smartag/describeuseronlineclientstatistics.html
func (client *Client) DescribeUserOnlineClientStatistics(request *DescribeUserOnlineClientStatisticsRequest) (response *DescribeUserOnlineClientStatisticsResponse, err error) {
	response = CreateDescribeUserOnlineClientStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserOnlineClientStatisticsWithChan invokes the smartag.DescribeUserOnlineClientStatistics API asynchronously
// api document: https://help.aliyun.com/api/smartag/describeuseronlineclientstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserOnlineClientStatisticsWithChan(request *DescribeUserOnlineClientStatisticsRequest) (<-chan *DescribeUserOnlineClientStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserOnlineClientStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserOnlineClientStatistics(request)
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

// DescribeUserOnlineClientStatisticsWithCallback invokes the smartag.DescribeUserOnlineClientStatistics API asynchronously
// api document: https://help.aliyun.com/api/smartag/describeuseronlineclientstatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserOnlineClientStatisticsWithCallback(request *DescribeUserOnlineClientStatisticsRequest, callback func(response *DescribeUserOnlineClientStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserOnlineClientStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserOnlineClientStatistics(request)
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

// DescribeUserOnlineClientStatisticsRequest is the request struct for api DescribeUserOnlineClientStatistics
type DescribeUserOnlineClientStatisticsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	UserNames            *[]string        `position:"Query" name:"UserNames"  type:"Repeated"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUserOnlineClientStatisticsResponse is the response struct for api DescribeUserOnlineClientStatistics
type DescribeUserOnlineClientStatisticsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	UserStatistics UserStatistics `json:"UserStatistics" xml:"UserStatistics"`
}

// CreateDescribeUserOnlineClientStatisticsRequest creates a request to invoke DescribeUserOnlineClientStatistics API
func CreateDescribeUserOnlineClientStatisticsRequest() (request *DescribeUserOnlineClientStatisticsRequest) {
	request = &DescribeUserOnlineClientStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeUserOnlineClientStatistics", "smartag", "openAPI")
	return
}

// CreateDescribeUserOnlineClientStatisticsResponse creates a response to parse from DescribeUserOnlineClientStatistics response
func CreateDescribeUserOnlineClientStatisticsResponse() (response *DescribeUserOnlineClientStatisticsResponse) {
	response = &DescribeUserOnlineClientStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

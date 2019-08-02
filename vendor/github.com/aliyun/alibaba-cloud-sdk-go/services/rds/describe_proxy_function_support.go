package rds

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

// DescribeProxyFunctionSupport invokes the rds.DescribeProxyFunctionSupport API synchronously
// api document: https://help.aliyun.com/api/rds/describeproxyfunctionsupport.html
func (client *Client) DescribeProxyFunctionSupport(request *DescribeProxyFunctionSupportRequest) (response *DescribeProxyFunctionSupportResponse, err error) {
	response = CreateDescribeProxyFunctionSupportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProxyFunctionSupportWithChan invokes the rds.DescribeProxyFunctionSupport API asynchronously
// api document: https://help.aliyun.com/api/rds/describeproxyfunctionsupport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProxyFunctionSupportWithChan(request *DescribeProxyFunctionSupportRequest) (<-chan *DescribeProxyFunctionSupportResponse, <-chan error) {
	responseChan := make(chan *DescribeProxyFunctionSupportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProxyFunctionSupport(request)
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

// DescribeProxyFunctionSupportWithCallback invokes the rds.DescribeProxyFunctionSupport API asynchronously
// api document: https://help.aliyun.com/api/rds/describeproxyfunctionsupport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProxyFunctionSupportWithCallback(request *DescribeProxyFunctionSupportRequest, callback func(response *DescribeProxyFunctionSupportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProxyFunctionSupportResponse
		var err error
		defer close(result)
		response, err = client.DescribeProxyFunctionSupport(request)
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

// DescribeProxyFunctionSupportRequest is the request struct for api DescribeProxyFunctionSupport
type DescribeProxyFunctionSupportRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeProxyFunctionSupportResponse is the response struct for api DescribeProxyFunctionSupport
type DescribeProxyFunctionSupportResponse struct {
	*responses.BaseResponse
	RequestId                       string `json:"RequestId" xml:"RequestId"`
	IsProxySwitchEnable             bool   `json:"IsProxySwitchEnable" xml:"IsProxySwitchEnable"`
	IsRwsplitEnable                 bool   `json:"IsRwsplitEnable" xml:"IsRwsplitEnable"`
	IsRwsplitSupportReplicationLag  bool   `json:"IsRwsplitSupportReplicationLag" xml:"IsRwsplitSupportReplicationLag"`
	IsRwsplitSupportWeight          bool   `json:"IsRwsplitSupportWeight" xml:"IsRwsplitSupportWeight"`
	IsTransparentSwitchEnable       bool   `json:"IsTransparentSwitchEnable" xml:"IsTransparentSwitchEnable"`
	IsShortConnectionOptimizeEnable bool   `json:"IsShortConnectionOptimizeEnable" xml:"IsShortConnectionOptimizeEnable"`
	IsAntiBruteFroceEnable          bool   `json:"IsAntiBruteFroceEnable" xml:"IsAntiBruteFroceEnable"`
}

// CreateDescribeProxyFunctionSupportRequest creates a request to invoke DescribeProxyFunctionSupport API
func CreateDescribeProxyFunctionSupportRequest() (request *DescribeProxyFunctionSupportRequest) {
	request = &DescribeProxyFunctionSupportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeProxyFunctionSupport", "rds", "openAPI")
	return
}

// CreateDescribeProxyFunctionSupportResponse creates a response to parse from DescribeProxyFunctionSupport response
func CreateDescribeProxyFunctionSupportResponse() (response *DescribeProxyFunctionSupportResponse) {
	response = &DescribeProxyFunctionSupportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

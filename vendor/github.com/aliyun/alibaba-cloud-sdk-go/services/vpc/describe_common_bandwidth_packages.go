package vpc

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

// DescribeCommonBandwidthPackages invokes the vpc.DescribeCommonBandwidthPackages API synchronously
// api document: https://help.aliyun.com/api/vpc/describecommonbandwidthpackages.html
func (client *Client) DescribeCommonBandwidthPackages(request *DescribeCommonBandwidthPackagesRequest) (response *DescribeCommonBandwidthPackagesResponse, err error) {
	response = CreateDescribeCommonBandwidthPackagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCommonBandwidthPackagesWithChan invokes the vpc.DescribeCommonBandwidthPackages API asynchronously
// api document: https://help.aliyun.com/api/vpc/describecommonbandwidthpackages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCommonBandwidthPackagesWithChan(request *DescribeCommonBandwidthPackagesRequest) (<-chan *DescribeCommonBandwidthPackagesResponse, <-chan error) {
	responseChan := make(chan *DescribeCommonBandwidthPackagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCommonBandwidthPackages(request)
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

// DescribeCommonBandwidthPackagesWithCallback invokes the vpc.DescribeCommonBandwidthPackages API asynchronously
// api document: https://help.aliyun.com/api/vpc/describecommonbandwidthpackages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCommonBandwidthPackagesWithCallback(request *DescribeCommonBandwidthPackagesRequest, callback func(response *DescribeCommonBandwidthPackagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCommonBandwidthPackagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCommonBandwidthPackages(request)
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

// DescribeCommonBandwidthPackagesRequest is the request struct for api DescribeCommonBandwidthPackages
type DescribeCommonBandwidthPackagesRequest struct {
	*requests.RpcRequest
	ResourceGroupId        string           `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId     string           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	Name                   string           `position:"Query" name:"Name"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	IncludeReservationData requests.Boolean `position:"Query" name:"IncludeReservationData"`
	PageNumber             requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeCommonBandwidthPackagesResponse is the response struct for api DescribeCommonBandwidthPackages
type DescribeCommonBandwidthPackagesResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	TotalCount              int                     `json:"TotalCount" xml:"TotalCount"`
	PageNumber              int                     `json:"PageNumber" xml:"PageNumber"`
	PageSize                int                     `json:"PageSize" xml:"PageSize"`
	CommonBandwidthPackages CommonBandwidthPackages `json:"CommonBandwidthPackages" xml:"CommonBandwidthPackages"`
}

// CreateDescribeCommonBandwidthPackagesRequest creates a request to invoke DescribeCommonBandwidthPackages API
func CreateDescribeCommonBandwidthPackagesRequest() (request *DescribeCommonBandwidthPackagesRequest) {
	request = &DescribeCommonBandwidthPackagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCommonBandwidthPackages", "vpc", "openAPI")
	return
}

// CreateDescribeCommonBandwidthPackagesResponse creates a response to parse from DescribeCommonBandwidthPackages response
func CreateDescribeCommonBandwidthPackagesResponse() (response *DescribeCommonBandwidthPackagesResponse) {
	response = &DescribeCommonBandwidthPackagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

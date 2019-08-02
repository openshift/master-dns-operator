package r_kvstore

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

// DescribeCacheAnalysisReport invokes the r_kvstore.DescribeCacheAnalysisReport API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecacheanalysisreport.html
func (client *Client) DescribeCacheAnalysisReport(request *DescribeCacheAnalysisReportRequest) (response *DescribeCacheAnalysisReportResponse, err error) {
	response = CreateDescribeCacheAnalysisReportResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCacheAnalysisReportWithChan invokes the r_kvstore.DescribeCacheAnalysisReport API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecacheanalysisreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCacheAnalysisReportWithChan(request *DescribeCacheAnalysisReportRequest) (<-chan *DescribeCacheAnalysisReportResponse, <-chan error) {
	responseChan := make(chan *DescribeCacheAnalysisReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCacheAnalysisReport(request)
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

// DescribeCacheAnalysisReportWithCallback invokes the r_kvstore.DescribeCacheAnalysisReport API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describecacheanalysisreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCacheAnalysisReportWithCallback(request *DescribeCacheAnalysisReportRequest, callback func(response *DescribeCacheAnalysisReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCacheAnalysisReportResponse
		var err error
		defer close(result)
		response, err = client.DescribeCacheAnalysisReport(request)
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

// DescribeCacheAnalysisReportRequest is the request struct for api DescribeCacheAnalysisReport
type DescribeCacheAnalysisReportRequest struct {
	*requests.RpcRequest
	Date                 string           `position:"Query" name:"Date"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AnalysisType         string           `position:"Query" name:"AnalysisType"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumbers          requests.Integer `position:"Query" name:"PageNumbers"`
	NodeId               string           `position:"Query" name:"NodeId"`
}

// DescribeCacheAnalysisReportResponse is the response struct for api DescribeCacheAnalysisReport
type DescribeCacheAnalysisReportResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageSize         int    `json:"PageSize" xml:"PageSize"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	HotKeys          string `json:"HotKeys" xml:"HotKeys"`
	BigKeys          string `json:"BigKeys" xml:"BigKeys"`
}

// CreateDescribeCacheAnalysisReportRequest creates a request to invoke DescribeCacheAnalysisReport API
func CreateDescribeCacheAnalysisReportRequest() (request *DescribeCacheAnalysisReportRequest) {
	request = &DescribeCacheAnalysisReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeCacheAnalysisReport", "redisa", "openAPI")
	return
}

// CreateDescribeCacheAnalysisReportResponse creates a response to parse from DescribeCacheAnalysisReport response
func CreateDescribeCacheAnalysisReportResponse() (response *DescribeCacheAnalysisReportResponse) {
	response = &DescribeCacheAnalysisReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

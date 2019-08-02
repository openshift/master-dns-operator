package domain

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

// QueryDSRecord invokes the domain.QueryDSRecord API synchronously
// api document: https://help.aliyun.com/api/domain/querydsrecord.html
func (client *Client) QueryDSRecord(request *QueryDSRecordRequest) (response *QueryDSRecordResponse, err error) {
	response = CreateQueryDSRecordResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDSRecordWithChan invokes the domain.QueryDSRecord API asynchronously
// api document: https://help.aliyun.com/api/domain/querydsrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDSRecordWithChan(request *QueryDSRecordRequest) (<-chan *QueryDSRecordResponse, <-chan error) {
	responseChan := make(chan *QueryDSRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDSRecord(request)
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

// QueryDSRecordWithCallback invokes the domain.QueryDSRecord API asynchronously
// api document: https://help.aliyun.com/api/domain/querydsrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDSRecordWithCallback(request *QueryDSRecordRequest, callback func(response *QueryDSRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDSRecordResponse
		var err error
		defer close(result)
		response, err = client.QueryDSRecord(request)
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

// QueryDSRecordRequest is the request struct for api QueryDSRecord
type QueryDSRecordRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryDSRecordResponse is the response struct for api QueryDSRecord
type QueryDSRecordResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	DSRecordList []DSRecord `json:"DSRecordList" xml:"DSRecordList"`
}

// CreateQueryDSRecordRequest creates a request to invoke QueryDSRecord API
func CreateQueryDSRecordRequest() (request *QueryDSRecordRequest) {
	request = &QueryDSRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDSRecord", "", "")
	return
}

// CreateQueryDSRecordResponse creates a response to parse from QueryDSRecord response
func CreateQueryDSRecordResponse() (response *QueryDSRecordResponse) {
	response = &QueryDSRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

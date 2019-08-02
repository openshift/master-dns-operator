package scdn

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

// DescribeScdnDomainCname invokes the scdn.DescribeScdnDomainCname API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaincname.html
func (client *Client) DescribeScdnDomainCname(request *DescribeScdnDomainCnameRequest) (response *DescribeScdnDomainCnameResponse, err error) {
	response = CreateDescribeScdnDomainCnameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainCnameWithChan invokes the scdn.DescribeScdnDomainCname API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaincname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainCnameWithChan(request *DescribeScdnDomainCnameRequest) (<-chan *DescribeScdnDomainCnameResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainCnameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainCname(request)
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

// DescribeScdnDomainCnameWithCallback invokes the scdn.DescribeScdnDomainCname API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomaincname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainCnameWithCallback(request *DescribeScdnDomainCnameRequest, callback func(response *DescribeScdnDomainCnameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainCnameResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainCname(request)
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

// DescribeScdnDomainCnameRequest is the request struct for api DescribeScdnDomainCname
type DescribeScdnDomainCnameRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnDomainCnameResponse is the response struct for api DescribeScdnDomainCname
type DescribeScdnDomainCnameResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	CnameDatas CnameDatas `json:"CnameDatas" xml:"CnameDatas"`
}

// CreateDescribeScdnDomainCnameRequest creates a request to invoke DescribeScdnDomainCname API
func CreateDescribeScdnDomainCnameRequest() (request *DescribeScdnDomainCnameRequest) {
	request = &DescribeScdnDomainCnameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainCname", "scdn", "openAPI")
	return
}

// CreateDescribeScdnDomainCnameResponse creates a response to parse from DescribeScdnDomainCname response
func CreateDescribeScdnDomainCnameResponse() (response *DescribeScdnDomainCnameResponse) {
	response = &DescribeScdnDomainCnameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

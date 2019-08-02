package petadata

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

// ReleaseInstancePublicConnection invokes the petadata.ReleaseInstancePublicConnection API synchronously
// api document: https://help.aliyun.com/api/petadata/releaseinstancepublicconnection.html
func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (response *ReleaseInstancePublicConnectionResponse, err error) {
	response = CreateReleaseInstancePublicConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseInstancePublicConnectionWithChan invokes the petadata.ReleaseInstancePublicConnection API asynchronously
// api document: https://help.aliyun.com/api/petadata/releaseinstancepublicconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseInstancePublicConnectionWithChan(request *ReleaseInstancePublicConnectionRequest) (<-chan *ReleaseInstancePublicConnectionResponse, <-chan error) {
	responseChan := make(chan *ReleaseInstancePublicConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseInstancePublicConnection(request)
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

// ReleaseInstancePublicConnectionWithCallback invokes the petadata.ReleaseInstancePublicConnection API asynchronously
// api document: https://help.aliyun.com/api/petadata/releaseinstancepublicconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseInstancePublicConnectionWithCallback(request *ReleaseInstancePublicConnectionRequest, callback func(response *ReleaseInstancePublicConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseInstancePublicConnectionResponse
		var err error
		defer close(result)
		response, err = client.ReleaseInstancePublicConnection(request)
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

// ReleaseInstancePublicConnectionRequest is the request struct for api ReleaseInstancePublicConnection
type ReleaseInstancePublicConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken           string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId            string           `position:"Query" name:"DBInstanceId"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	CurrentConnectionString string           `position:"Query" name:"CurrentConnectionString"`
}

// ReleaseInstancePublicConnectionResponse is the response struct for api ReleaseInstancePublicConnection
type ReleaseInstancePublicConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseInstancePublicConnectionRequest creates a request to invoke ReleaseInstancePublicConnection API
func CreateReleaseInstancePublicConnectionRequest() (request *ReleaseInstancePublicConnectionRequest) {
	request = &ReleaseInstancePublicConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PetaData", "2016-01-01", "ReleaseInstancePublicConnection", "petadata", "openAPI")
	return
}

// CreateReleaseInstancePublicConnectionResponse creates a response to parse from ReleaseInstancePublicConnection response
func CreateReleaseInstancePublicConnectionResponse() (response *ReleaseInstancePublicConnectionResponse) {
	response = &ReleaseInstancePublicConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

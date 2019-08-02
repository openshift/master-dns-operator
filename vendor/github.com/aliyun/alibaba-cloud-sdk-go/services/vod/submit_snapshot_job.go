package vod

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

// SubmitSnapshotJob invokes the vod.SubmitSnapshotJob API synchronously
// api document: https://help.aliyun.com/api/vod/submitsnapshotjob.html
func (client *Client) SubmitSnapshotJob(request *SubmitSnapshotJobRequest) (response *SubmitSnapshotJobResponse, err error) {
	response = CreateSubmitSnapshotJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitSnapshotJobWithChan invokes the vod.SubmitSnapshotJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitsnapshotjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitSnapshotJobWithChan(request *SubmitSnapshotJobRequest) (<-chan *SubmitSnapshotJobResponse, <-chan error) {
	responseChan := make(chan *SubmitSnapshotJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSnapshotJob(request)
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

// SubmitSnapshotJobWithCallback invokes the vod.SubmitSnapshotJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitsnapshotjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitSnapshotJobWithCallback(request *SubmitSnapshotJobRequest, callback func(response *SubmitSnapshotJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSnapshotJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitSnapshotJob(request)
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

// SubmitSnapshotJobRequest is the request struct for api SubmitSnapshotJob
type SubmitSnapshotJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Count                requests.Integer `position:"Query" name:"Count"`
	VideoId              string           `position:"Query" name:"VideoId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UserData             string           `position:"Query" name:"UserData"`
	SpecifiedOffsetTime  requests.Integer `position:"Query" name:"SpecifiedOffsetTime"`
	Width                string           `position:"Query" name:"Width"`
	Interval             requests.Integer `position:"Query" name:"Interval"`
	SpriteSnapshotConfig string           `position:"Query" name:"SpriteSnapshotConfig"`
	SnapshotTemplateId   string           `position:"Query" name:"SnapshotTemplateId"`
	Height               string           `position:"Query" name:"Height"`
}

// SubmitSnapshotJobResponse is the response struct for api SubmitSnapshotJob
type SubmitSnapshotJobResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	SnapshotJob SnapshotJob `json:"SnapshotJob" xml:"SnapshotJob"`
}

// CreateSubmitSnapshotJobRequest creates a request to invoke SubmitSnapshotJob API
func CreateSubmitSnapshotJobRequest() (request *SubmitSnapshotJobRequest) {
	request = &SubmitSnapshotJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitSnapshotJob", "vod", "openAPI")
	return
}

// CreateSubmitSnapshotJobResponse creates a response to parse from SubmitSnapshotJob response
func CreateSubmitSnapshotJobResponse() (response *SubmitSnapshotJobResponse) {
	response = &SubmitSnapshotJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

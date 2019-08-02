package dbs

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

// ModifyBackupStrategy invokes the dbs.ModifyBackupStrategy API synchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupstrategy.html
func (client *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
	response = CreateModifyBackupStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupStrategyWithChan invokes the dbs.ModifyBackupStrategy API asynchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupStrategyWithChan(request *ModifyBackupStrategyRequest) (<-chan *ModifyBackupStrategyResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupStrategy(request)
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

// ModifyBackupStrategyWithCallback invokes the dbs.ModifyBackupStrategy API asynchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupStrategyWithCallback(request *ModifyBackupStrategyRequest, callback func(response *ModifyBackupStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupStrategyResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupStrategy(request)
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

// ModifyBackupStrategyRequest is the request struct for api ModifyBackupStrategy
type ModifyBackupStrategyRequest struct {
	*requests.RpcRequest
	BackupPeriod    string `position:"Query" name:"BackupPeriod"`
	BackupStartTime string `position:"Query" name:"BackupStartTime"`
	ClientToken     string `position:"Query" name:"ClientToken"`
	BackupPlanId    string `position:"Query" name:"BackupPlanId"`
	OwnerId         string `position:"Query" name:"OwnerId"`
}

// ModifyBackupStrategyResponse is the response struct for api ModifyBackupStrategy
type ModifyBackupStrategyResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	NeedPrecheck   bool   `json:"NeedPrecheck" xml:"NeedPrecheck"`
}

// CreateModifyBackupStrategyRequest creates a request to invoke ModifyBackupStrategy API
func CreateModifyBackupStrategyRequest() (request *ModifyBackupStrategyRequest) {
	request = &ModifyBackupStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ModifyBackupStrategy", "cbs", "openAPI")
	return
}

// CreateModifyBackupStrategyResponse creates a response to parse from ModifyBackupStrategy response
func CreateModifyBackupStrategyResponse() (response *ModifyBackupStrategyResponse) {
	response = &ModifyBackupStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

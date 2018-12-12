package uis

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

// ModifyUisNodeAttribute invokes the uis.ModifyUisNodeAttribute API synchronously
// api document: https://help.aliyun.com/api/uis/modifyuisnodeattribute.html
func (client *Client) ModifyUisNodeAttribute(request *ModifyUisNodeAttributeRequest) (response *ModifyUisNodeAttributeResponse, err error) {
	response = CreateModifyUisNodeAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUisNodeAttributeWithChan invokes the uis.ModifyUisNodeAttribute API asynchronously
// api document: https://help.aliyun.com/api/uis/modifyuisnodeattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUisNodeAttributeWithChan(request *ModifyUisNodeAttributeRequest) (<-chan *ModifyUisNodeAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyUisNodeAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUisNodeAttribute(request)
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

// ModifyUisNodeAttributeWithCallback invokes the uis.ModifyUisNodeAttribute API asynchronously
// api document: https://help.aliyun.com/api/uis/modifyuisnodeattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUisNodeAttributeWithCallback(request *ModifyUisNodeAttributeRequest, callback func(response *ModifyUisNodeAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUisNodeAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyUisNodeAttribute(request)
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

// ModifyUisNodeAttributeRequest is the request struct for api ModifyUisNodeAttribute
type ModifyUisNodeAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UisNodeBandwidth     requests.Integer `position:"Query" name:"UisNodeBandwidth"`
	UisNodeId            string           `position:"Query" name:"UisNodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UisId                string           `position:"Query" name:"UisId"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyUisNodeAttributeResponse is the response struct for api ModifyUisNodeAttribute
type ModifyUisNodeAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyUisNodeAttributeRequest creates a request to invoke ModifyUisNodeAttribute API
func CreateModifyUisNodeAttributeRequest() (request *ModifyUisNodeAttributeRequest) {
	request = &ModifyUisNodeAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "ModifyUisNodeAttribute", "uis", "openAPI")
	return
}

// CreateModifyUisNodeAttributeResponse creates a response to parse from ModifyUisNodeAttribute response
func CreateModifyUisNodeAttributeResponse() (response *ModifyUisNodeAttributeResponse) {
	response = &ModifyUisNodeAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
package qualitycheck

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

// UploadAudioData4Pre invokes the qualitycheck.UploadAudioData4Pre API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadaudiodata4pre.html
func (client *Client) UploadAudioData4Pre(request *UploadAudioData4PreRequest) (response *UploadAudioData4PreResponse, err error) {
	response = CreateUploadAudioData4PreResponse()
	err = client.DoAction(request, response)
	return
}

// UploadAudioData4PreWithChan invokes the qualitycheck.UploadAudioData4Pre API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadaudiodata4pre.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadAudioData4PreWithChan(request *UploadAudioData4PreRequest) (<-chan *UploadAudioData4PreResponse, <-chan error) {
	responseChan := make(chan *UploadAudioData4PreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadAudioData4Pre(request)
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

// UploadAudioData4PreWithCallback invokes the qualitycheck.UploadAudioData4Pre API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/uploadaudiodata4pre.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadAudioData4PreWithCallback(request *UploadAudioData4PreRequest, callback func(response *UploadAudioData4PreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadAudioData4PreResponse
		var err error
		defer close(result)
		response, err = client.UploadAudioData4Pre(request)
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

// UploadAudioData4PreRequest is the request struct for api UploadAudioData4Pre
type UploadAudioData4PreRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UploadAudioData4PreResponse is the response struct for api UploadAudioData4Pre
type UploadAudioData4PreResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateUploadAudioData4PreRequest creates a request to invoke UploadAudioData4Pre API
func CreateUploadAudioData4PreRequest() (request *UploadAudioData4PreRequest) {
	request = &UploadAudioData4PreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UploadAudioData4Pre", "", "")
	return
}

// CreateUploadAudioData4PreResponse creates a response to parse from UploadAudioData4Pre response
func CreateUploadAudioData4PreResponse() (response *UploadAudioData4PreResponse) {
	response = &UploadAudioData4PreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

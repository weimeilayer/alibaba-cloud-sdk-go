package cloudauth

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

// DescribeFaceVerify invokes the cloudauth.DescribeFaceVerify API synchronously
// api document: https://help.aliyun.com/api/cloudauth/describefaceverify.html
func (client *Client) DescribeFaceVerify(request *DescribeFaceVerifyRequest) (response *DescribeFaceVerifyResponse, err error) {
	response = CreateDescribeFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFaceVerifyWithChan invokes the cloudauth.DescribeFaceVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describefaceverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFaceVerifyWithChan(request *DescribeFaceVerifyRequest) (<-chan *DescribeFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *DescribeFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFaceVerify(request)
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

// DescribeFaceVerifyWithCallback invokes the cloudauth.DescribeFaceVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describefaceverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFaceVerifyWithCallback(request *DescribeFaceVerifyRequest, callback func(response *DescribeFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.DescribeFaceVerify(request)
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

// DescribeFaceVerifyRequest is the request struct for api DescribeFaceVerify
type DescribeFaceVerifyRequest struct {
	*requests.RpcRequest
	SceneId   requests.Integer `position:"Query" name:"SceneId"`
	CertifyId string           `position:"Query" name:"CertifyId"`
}

// DescribeFaceVerifyResponse is the response struct for api DescribeFaceVerify
type DescribeFaceVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateDescribeFaceVerifyRequest creates a request to invoke DescribeFaceVerify API
func CreateDescribeFaceVerifyRequest() (request *DescribeFaceVerifyRequest) {
	request = &DescribeFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeFaceVerify", "", "")
	return
}

// CreateDescribeFaceVerifyResponse creates a response to parse from DescribeFaceVerify response
func CreateDescribeFaceVerifyResponse() (response *DescribeFaceVerifyResponse) {
	response = &DescribeFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

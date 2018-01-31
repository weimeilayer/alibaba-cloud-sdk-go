package mts

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

func (client *Client) QueryFpShotJobList(request *QueryFpShotJobListRequest) (response *QueryFpShotJobListResponse, err error) {
	response = CreateQueryFpShotJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryFpShotJobListWithChan(request *QueryFpShotJobListRequest) (<-chan *QueryFpShotJobListResponse, <-chan error) {
	responseChan := make(chan *QueryFpShotJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFpShotJobList(request)
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

func (client *Client) QueryFpShotJobListWithCallback(request *QueryFpShotJobListRequest, callback func(response *QueryFpShotJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFpShotJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryFpShotJobList(request)
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

type QueryFpShotJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryFpShotJobListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	NonExistIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistIds" xml:"NonExistIds"`
	FpShotJobList struct {
		FpShotJob []struct {
			Id           string `json:"Id" xml:"Id"`
			UserData     string `json:"UserData" xml:"UserData"`
			PipelineId   string `json:"PipelineId" xml:"PipelineId"`
			State        string `json:"State" xml:"State"`
			Code         string `json:"Code" xml:"Code"`
			Message      string `json:"Message" xml:"Message"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			InputFile    struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
			} `json:"InputFile" xml:"InputFile"`
			FpShotResult struct {
				FpShots struct {
					FpShot []struct {
						PrimaryKey   string `json:"PrimaryKey" xml:"PrimaryKey"`
						Similarity   string `json:"Similarity" xml:"Similarity"`
						FpShotSlices struct {
							FpShotSlice []struct {
								Input struct {
									Start    string `json:"Start" xml:"Start"`
									Duration string `json:"Duration" xml:"Duration"`
								} `json:"Input" xml:"Input"`
								Duplication struct {
									Start    string `json:"Start" xml:"Start"`
									Duration string `json:"Duration" xml:"Duration"`
								} `json:"Duplication" xml:"Duplication"`
							} `json:"FpShotSlice" xml:"FpShotSlice"`
						} `json:"FpShotSlices" xml:"FpShotSlices"`
					} `json:"FpShot" xml:"FpShot"`
				} `json:"FpShots" xml:"FpShots"`
			} `json:"FpShotResult" xml:"FpShotResult"`
		} `json:"FpShotJob" xml:"FpShotJob"`
	} `json:"FpShotJobList" xml:"FpShotJobList"`
}

func CreateQueryFpShotJobListRequest() (request *QueryFpShotJobListRequest) {
	request = &QueryFpShotJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryFpShotJobList", "mts", "openAPI")
	return
}

func CreateQueryFpShotJobListResponse() (response *QueryFpShotJobListResponse) {
	response = &QueryFpShotJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
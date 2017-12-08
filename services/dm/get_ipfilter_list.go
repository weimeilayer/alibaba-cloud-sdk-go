
package dm

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

func (client *Client) GetIpfilterList(request *GetIpfilterListRequest) (response *GetIpfilterListResponse, err error) {
response = CreateGetIpfilterListResponse()
err = client.DoAction(request, response)
return
}

func (client *Client) GetIpfilterListWithChan(request *GetIpfilterListRequest) (<-chan *GetIpfilterListResponse, <-chan error) {
responseChan := make(chan *GetIpfilterListResponse, 1)
errChan := make(chan error, 1)
err := client.AddAsyncTask(func() {
defer close(responseChan)
defer close(errChan)
response, err :=  client.GetIpfilterList(request)
responseChan <- response
errChan <- err
})
if err != nil {
errChan <- err
close(responseChan)
close(errChan)
}
return responseChan, errChan
}

func (client *Client) GetIpfilterListWithCallback(request *GetIpfilterListRequest, callback func(response *GetIpfilterListResponse, err error)) (<-chan int) {
result := make(chan int, 1)
err := client.AddAsyncTask(func() {
var response *GetIpfilterListResponse
var err error
defer close(result)
response, err = client.GetIpfilterList(request)
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

type GetIpfilterListRequest struct {
*requests.RpcRequest
                ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
                ResourceOwnerId  string `position:"Query" name:"ResourceOwnerId"`
                OwnerId  string `position:"Query" name:"OwnerId"`
}


type GetIpfilterListResponse struct {
*responses.BaseResponse
            RequestId     string `json:"RequestId" xml:"RequestId"`
            TotalCount     int `json:"TotalCount" xml:"TotalCount"`
            PageNumber     int `json:"PageNumber" xml:"PageNumber"`
            PageSize     int `json:"PageSize" xml:"PageSize"`
                Data struct {
                    Ipfilters []struct {
            Id     string `json:"Id" xml:"Id"`
            IpAddress     string `json:"IpAddress" xml:"IpAddress"`
            CreateTime     string `json:"CreateTime" xml:"CreateTime"`
                    }   `json:"ipfilters" xml:"ipfilters"`
                } `json:"data" xml:"data"`
}

func CreateGetIpfilterListRequest() (request *GetIpfilterListRequest) {
request = &GetIpfilterListRequest{
RpcRequest: &requests.RpcRequest{},
}
request.InitWithApiInfo("Dm", "2015-11-23", "GetIpfilterList", "", "")
return
}

func CreateGetIpfilterListResponse() (response *GetIpfilterListResponse) {
response = &GetIpfilterListResponse{
BaseResponse: &responses.BaseResponse{},
}
return
}

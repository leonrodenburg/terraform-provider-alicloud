package alidns

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

// DescribeGtmRecoveryPlans invokes the alidns.DescribeGtmRecoveryPlans API synchronously
// api document: https://help.aliyun.com/api/alidns/describegtmrecoveryplans.html
func (client *Client) DescribeGtmRecoveryPlans(request *DescribeGtmRecoveryPlansRequest) (response *DescribeGtmRecoveryPlansResponse, err error) {
	response = CreateDescribeGtmRecoveryPlansResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGtmRecoveryPlansWithChan invokes the alidns.DescribeGtmRecoveryPlans API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtmrecoveryplans.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmRecoveryPlansWithChan(request *DescribeGtmRecoveryPlansRequest) (<-chan *DescribeGtmRecoveryPlansResponse, <-chan error) {
	responseChan := make(chan *DescribeGtmRecoveryPlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGtmRecoveryPlans(request)
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

// DescribeGtmRecoveryPlansWithCallback invokes the alidns.DescribeGtmRecoveryPlans API asynchronously
// api document: https://help.aliyun.com/api/alidns/describegtmrecoveryplans.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGtmRecoveryPlansWithCallback(request *DescribeGtmRecoveryPlansRequest, callback func(response *DescribeGtmRecoveryPlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGtmRecoveryPlansResponse
		var err error
		defer close(result)
		response, err = client.DescribeGtmRecoveryPlans(request)
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

// DescribeGtmRecoveryPlansRequest is the request struct for api DescribeGtmRecoveryPlans
type DescribeGtmRecoveryPlansRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Keyword      string           `position:"Query" name:"Keyword"`
}

// DescribeGtmRecoveryPlansResponse is the response struct for api DescribeGtmRecoveryPlans
type DescribeGtmRecoveryPlansResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	TotalItems    int           `json:"TotalItems" xml:"TotalItems"`
	TotalPages    int           `json:"TotalPages" xml:"TotalPages"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	RecoveryPlans RecoveryPlans `json:"RecoveryPlans" xml:"RecoveryPlans"`
}

// CreateDescribeGtmRecoveryPlansRequest creates a request to invoke DescribeGtmRecoveryPlans API
func CreateDescribeGtmRecoveryPlansRequest() (request *DescribeGtmRecoveryPlansRequest) {
	request = &DescribeGtmRecoveryPlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeGtmRecoveryPlans", "alidns", "openAPI")
	return
}

// CreateDescribeGtmRecoveryPlansResponse creates a response to parse from DescribeGtmRecoveryPlans response
func CreateDescribeGtmRecoveryPlansResponse() (response *DescribeGtmRecoveryPlansResponse) {
	response = &DescribeGtmRecoveryPlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
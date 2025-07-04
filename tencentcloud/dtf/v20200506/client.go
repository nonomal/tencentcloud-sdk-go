// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200506

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-05-06"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeTransactionsRequest() (request *DescribeTransactionsRequest) {
    request = &DescribeTransactionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dtf", APIVersion, "DescribeTransactions")
    
    
    return
}

func NewDescribeTransactionsResponse() (response *DescribeTransactionsResponse) {
    response = &DescribeTransactionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTransactions
// 查询主事务列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeTransactions(request *DescribeTransactionsRequest) (response *DescribeTransactionsResponse, err error) {
    return c.DescribeTransactionsWithContext(context.Background(), request)
}

// DescribeTransactions
// 查询主事务列表
//
// 可能返回的错误码:
//  MISSINGPARAMETER_GROUPIDREQUIRED = "MissingParameter.GroupIdRequired"
//  RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"
func (c *Client) DescribeTransactionsWithContext(ctx context.Context, request *DescribeTransactionsRequest) (response *DescribeTransactionsResponse, err error) {
    if request == nil {
        request = NewDescribeTransactionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTransactions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTransactionsResponse()
    err = c.Send(request, response)
    return
}

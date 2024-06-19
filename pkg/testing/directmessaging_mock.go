/*
Copyright 2022 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

import (
	"context"
	"net/http"

	mock "github.com/stretchr/testify/mock"

	"github.com/dapr/dapr/pkg/channel"
	invokev1 "github.com/dapr/dapr/pkg/messaging/v1"
)

// MockDirectMessaging is a semi-autogenerated mock type for the MockDirectMessaging type.
// Note: This file is created by copy/pasting values and renaming to use MockDirectMessaging instead of DirectMessaging.
// You run "mockery --name directMessaging" in "pkg/messaging" and modify the corresponding values here.
type MockDirectMessaging struct {
	mock.Mock
}

// Invoke provides a mock function with given fields: ctx, targetAppID, req
func (_m *MockDirectMessaging) Invoke(ctx context.Context, targetAppID string, req *invokev1.InvokeMethodRequest) (*invokev1.InvokeMethodResponse, error) {
	ret := _m.Called(ctx, targetAppID, req)

	var r0 *invokev1.InvokeMethodResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *invokev1.InvokeMethodRequest) (*invokev1.InvokeMethodResponse, error)); ok {
		return rf(ctx, targetAppID, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *invokev1.InvokeMethodRequest) *invokev1.InvokeMethodResponse); ok {
		r0 = rf(ctx, targetAppID, req)
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).(*invokev1.InvokeMethodResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *invokev1.InvokeMethodRequest) error); ok {
		r1 = rf(ctx, targetAppID, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockDirectMessaging) SetAppChannel(appChannel channel.AppChannel) {
	// nop
}

// SetHTTPEndpointsAppChannel provides a mock function with given fields: appChannel
func (_m *MockDirectMessaging) SetHTTPEndpointsAppChannels(nonResourceChannel channel.HTTPEndpointAppChannel, resourceChannels map[string]channel.HTTPEndpointAppChannel) {
	// nop
}

func (_m *MockDirectMessaging) Close() error {
	return nil
}

type FailingDirectMessaging struct {
	Failure           Failure
	SuccessStatusCode int
}

func (_m *FailingDirectMessaging) Invoke(ctx context.Context, targetAppID string, req *invokev1.InvokeMethodRequest) (*invokev1.InvokeMethodResponse, error) {
	r, err := req.ProtoWithData()
	if err != nil {
		return &invokev1.InvokeMethodResponse{}, err
	}
	err = _m.Failure.PerformFailure(string(r.GetMessage().GetData().GetValue()))
	if err != nil {
		return &invokev1.InvokeMethodResponse{}, err
	}
	statusCode := _m.SuccessStatusCode
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	// Setting up the headers passed in request
	md := r.GetMetadata()
	headers := make(map[string][]string)
	for k, v := range md {
		headers[k] = v.GetValues()
	}
	contentType := r.GetMessage().GetContentType()
	resp := invokev1.
		NewInvokeMethodResponse(int32(statusCode), http.StatusText(statusCode), nil).
		WithRawDataBytes(r.GetMessage().GetData().GetValue()).
		WithHTTPHeaders(headers).
		WithContentType(contentType)
	return resp, nil
}

func (_m *FailingDirectMessaging) SetAppChannel(appChannel channel.AppChannel) {
	// nop
}

// SetHTTPEndpointsAppChannel provides a mock function with given fields: appChannel
func (_m *FailingDirectMessaging) SetHTTPEndpointsAppChannels(nonResourceChannel channel.HTTPEndpointAppChannel, resourceChannels map[string]channel.HTTPEndpointAppChannel) {
	// nop
}
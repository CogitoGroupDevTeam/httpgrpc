package proxy

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/LLKennedy/httpgrpc/v2/httpapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type exampleGRPCServer struct{}

type request struct {
	Arg1 string `json:"arg1"`
	Arg2 int    `json:"arg2"`
	Arg3 string `json:"arg3"`
}

func (r *request) GetArg1() string {
	if r == nil {
		return ""
	}
	return r.Arg1
}

func (r *request) GetArg2() int {
	if r == nil {
		return 0
	}
	return r.Arg2
}

func (r *request) GetArg3() string {
	if r == nil {
		return ""
	}
	return r.Arg3
}

type response struct {
	Arg3 bool `json:"arg3"`
}

func (s *exampleGRPCServer) DoThing(ctx context.Context, req *request) (*response, error) {
	return &response{
		Arg3: req.GetArg1() == "hello" && req.GetArg2() == 3 && req.GetArg3() == "goodbye",
	}, nil
}

type exampleService struct{}

func (e *exampleService) Example(ctx context.Context, req *ExampleRequest) (res *ExampleResponse, err error) {
	res = &ExampleResponse{
		FullResponseData: fmt.Sprintf("%d", req.GetOtherThing()),
		Done:             !req.GetToggle(),
		Output:           req.GetMainData(),
	}
	return res, nil
}

func TestServer_Proxy(t *testing.T) {
	tests := []struct {
		name        string
		s           *Server
		ctx         context.Context
		req         *httpapi.Request
		want        *httpapi.Response
		expectedErr string
	}{
		{
			name: "standard grpc call unimplemented",
			s: &Server{
				api: map[string]map[string]apiMethod{
					"POST": {
						"Proxy": apiMethod{pattern: apiMethodPatternStructStruct, reflection: reflect.TypeOf(&httpapi.UnimplementedExposedServiceServer{}).Method(0)},
					},
				},
				innerServer: &httpapi.UnimplementedExposedServiceServer{},
			},
			ctx: new(mockContext),
			req: &httpapi.Request{
				Method:    httpapi.Method_POST,
				Procedure: "Proxy",
			},
			want: &httpapi.Response{
				StatusCode: 500,
			},
			expectedErr: "rpc error: code = Unimplemented desc = method Proxy not implemented",
		},
		{
			name: "example grpc call success",
			s: &Server{
				api: map[string]map[string]apiMethod{
					"POST": {
						"Example": apiMethod{pattern: apiMethodPatternStructStruct, reflection: reflect.TypeOf(&exampleService{}).Method(0)},
					},
				},
				innerServer: &exampleService{},
			},
			ctx: new(mockContext),
			req: &httpapi.Request{
				Method:    httpapi.Method_POST,
				Procedure: "Example",
				Payload:   []byte(`{"toggle": false, "otherThing": 17}`),
				Params: map[string]*httpapi.MultiVal{
					"mainData": {Values: []string{"aGVsbG8="}},
				},
			},
			want: &httpapi.Response{
				StatusCode: 200,
				Payload:    []byte(`{"fullResponseData":"17","done":true,"output":"aGVsbG8="}`),
			},
		},
		{
			name:        "blank request",
			s:           &Server{},
			ctx:         new(mockContext),
			req:         &httpapi.Request{},
			want:        &httpapi.Response{},
			expectedErr: fmt.Sprintf("%v", status.Error(codes.Unimplemented, "httpgrpc: unknown HTTP method")),
		},
		{
			name: "unregistered method",
			s:    &Server{},
			ctx:  new(mockContext),
			req: &httpapi.Request{
				Method: httpapi.Method_POST,
			},
			want:        &httpapi.Response{},
			expectedErr: fmt.Sprintf("%v", status.Error(codes.Unimplemented, "httpgrpc: no POST methods defined in api")),
		},
		{
			name: "unregistered method",
			s: &Server{
				api: map[string]map[string]apiMethod{
					"POST": {},
				},
			},
			ctx: new(mockContext),
			req: &httpapi.Request{
				Method:    httpapi.Method_POST,
				Procedure: "DoThing",
			},
			want:        &httpapi.Response{},
			expectedErr: fmt.Sprintf("%v", status.Error(codes.Unimplemented, "httpgrpc: no procedure DoThing defined for POST method in api")),
		},
		// {
		// 	name: "non-standard grpc call",
		// 	s: &Server{
		// 		api: map[string]map[string]apiMethod{
		// 			"POST": {
		// 				"DoThing": apiMethod{pattern: apiMethodPatternStructStruct, reflection: reflect.TypeOf(&thingA{}).Method(0)},
		// 			},
		// 		},
		// 		innerServer: &thingA{},
		// 	},
		// 	ctx: new(mockContext),
		// 	req: &httpapi.Request{
		// 		Method:    httpapi.Method_POST,
		// 		Procedure: "DoThing",
		// 	},
		// 	want: &httpapi.Response{
		// 		StatusCode: 501,
		// 	},
		// 	expectedErr: "httpgrpc: nonstandard grpc signature not implemented",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Proxy(tt.ctx, tt.req)
			assert.Equal(t, tt.want, got)
			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedErr)
			}
		})
	}
}

// mockContext is an autogenerated mock type for the mockContext type
type mockContext struct {
	mock.Mock
}

// Deadline provides a mock function with given fields:
func (_m *mockContext) Deadline() (time.Time, bool) {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Done provides a mock function with given fields:
func (_m *mockContext) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *mockContext) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Value provides a mock function with given fields: key
func (_m *mockContext) Value(key interface{}) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

func Test_methodToString(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		in      httpapi.Method
		wantOut string
		wantErr string
	}{
		{
			name:    "invalid",
			in:      httpapi.Method(100),
			wantOut: "",
			wantErr: "unknown HTTP method",
		},
		{
			name:    "unknown",
			in:      httpapi.Method_UNKNOWN,
			wantOut: "",
			wantErr: "unknown HTTP method",
		},
		{
			name:    "CONNECT",
			in:      httpapi.Method_CONNECT,
			wantOut: "CONNECT",
		},
		{
			name:    "DELETE",
			in:      httpapi.Method_DELETE,
			wantOut: "DELETE",
		},
		{
			name:    "GET",
			in:      httpapi.Method_GET,
			wantOut: "GET",
		},
		{
			name:    "HEAD",
			in:      httpapi.Method_HEAD,
			wantOut: "HEAD",
		},
		{
			name:    "OPTIONS",
			in:      httpapi.Method_OPTIONS,
			wantOut: "OPTIONS",
		},
		{
			name:    "PATCH",
			in:      httpapi.Method_PATCH,
			wantOut: "PATCH",
		},
		{
			name:    "POST",
			in:      httpapi.Method_POST,
			wantOut: "POST",
		},
		{
			name:    "PUT",
			in:      httpapi.Method_PUT,
			wantOut: "PUT",
		},
		{
			name:    "TRACE",
			in:      httpapi.Method_TRACE,
			wantOut: "TRACE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := methodToString(tt.in)
			if tt.wantErr == "" {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantOut, gotOut)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}

func TestImpossiblePatterns(t *testing.T) {
	t.Run("struct-stream", func(t *testing.T) {
		res, err := callStructStream(context.Background(), nil, nil, reflect.Value{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Unimplemented desc = httpgrpc: Struct In, Stream Out is not yet supported, please manually implement exceptions for endpoint %!s(<nil>)")
	})
	t.Run("stream-struct", func(t *testing.T) {
		res, err := callStreamStruct(context.Background(), nil, nil, reflect.Value{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Unimplemented desc = httpgrpc: Stream In, Struct Out is not yet supported, please manually implement exceptions for endpoint %!s(<nil>)")
	})
	t.Run("stream-stream", func(t *testing.T) {
		res, err := callStreamStream(context.Background(), nil, nil, reflect.Value{})
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Unimplemented desc = httpgrpc: Stream In, Stream Out is not yet supported, please manually implement exceptions for endpoint %!s(<nil>)")
	})
}

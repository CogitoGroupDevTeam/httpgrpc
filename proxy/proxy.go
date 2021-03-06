package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/LLKennedy/httpgrpc/httpapi"
	"github.com/peterbourgon/mergemap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Proxy proxies connections through the server
func (s *Server) Proxy(ctx context.Context, req *httpapi.Request) (res *httpapi.Response, err error) {
	wrapErr := func(code codes.Code, err error) error {
		if err == nil {
			return nil
		}
		return status.Error(code, fmt.Sprintf("httpgrpc: %v", err))
	}
	defer func() {
		if r := recover(); r != nil {
			err = wrapErr(codes.Internal, fmt.Errorf("caught panic %v", r))
		}
	}()
	var handled bool
	handled, res, err = s.handleExceptions(ctx, req)
	if handled {
		// The user-defined exception handler already processes this request, we don't have to deal with it
		return
	}
	procType, caller, pattern, err := s.findProc(req.GetMethod(), req.GetProcedure())
	if err != nil {
		return &httpapi.Response{}, wrapErr(codes.Unimplemented, err)
	}
	if pattern == apiMethodPatternUnknown {
		return &httpapi.Response{}, wrapErr(codes.Unimplemented, fmt.Errorf("nonstandard grpc signature not implemented"))
	}
	return s.callProc(ctx, req, procType, caller, pattern)
}

func (s *Server) callProc(ctx context.Context, req *httpapi.Request, procType reflect.Type, caller reflect.Value, pattern apiMethodPattern) (res *httpapi.Response, err error) {
	wrapErr := func(code codes.Code, err error) error {
		if err == nil {
			return nil
		}
		return status.Error(code, fmt.Sprintf("httpgrpc: %v", err))
	}
	// We're going to rely on JSON unmarshalling logic to get data from the request to the new struct
	// Maybe one day there will be custom marshalling/unmarshalling capability here with tag parsing and method analysis for custom functions, but probably not
	// JSON tags are pretty much fit for purpose
	// We just have to make sure we have json to work with first
	var inputJSON []byte
	inputJSON, err = parseRequest(req)
	if err != nil {
		return &httpapi.Response{}, wrapErr(codes.Internal, err)
	}
	switch pattern {
	case apiMethodPatternStructStruct:
		res, err = callStructStruct(ctx, inputJSON, procType, caller)
	case apiMethodPatternStructStream:
		res, err = callStructStream(ctx, inputJSON, procType, caller)
	case apiMethodPatternStreamStruct:
		res, err = callStreamStruct(ctx, inputJSON, procType, caller)
	case apiMethodPatternStreamStream:
		res, err = callStreamStream(ctx, inputJSON, procType, caller)
	default:
		// This should be truly impossible during normal operation, we've checked for this like 20 times before this point
		res, err = &httpapi.Response{}, wrapErr(codes.Unimplemented, fmt.Errorf("nonstandard grpc signature not implemented"))
	}
	return
}

func (s *Server) findProc(httpMethod httpapi.Method, procName string) (procType reflect.Type, caller reflect.Value, pattern apiMethodPattern, err error) {
	var methodString string
	methodString, err = methodToString(httpMethod)
	if err != nil {
		return
	}
	methodMap, found := s.getAPI()[methodString]
	if !found {
		err = fmt.Errorf("no %s methods defined in api", methodString)
		return
	}
	var apiProc apiMethod
	apiProc, found = methodMap[procName]
	if !found {
		err = fmt.Errorf("no procedure %s defined for %s method in api", procName, httpMethod)
		return
	}
	procType = apiProc.reflection.Type
	pattern = apiProc.pattern
	caller = reflect.ValueOf(s.getInnerServer()).MethodByName(procName)
	return
}

// One struct in, one struct out
func callStructStruct(ctx context.Context, inputJSON []byte, procType reflect.Type, caller reflect.Value) (res *httpapi.Response, err error) {
	// Create new instance of struct argument to pass into real implementation
	builtRequest := reflect.New(procType.In(2).Elem())
	builtRequestPtr := builtRequest.Interface()
	builtRequestMessage, ok := builtRequestPtr.(proto.Message)
	if !ok {
		return &httpapi.Response{}, status.Error(codes.InvalidArgument, "httpgrpc: cannot convert json data to non-proto message using protojson")
	}
	if inputJSON != nil {
		err = protojson.Unmarshal(inputJSON, builtRequestMessage)
		if err != nil {
			return &httpapi.Response{}, status.Error(codes.InvalidArgument, fmt.Sprintf("httpgrpc: %v", err))
		}
	}
	// Actually call the inner procedure
	returnValues := caller.Call([]reflect.Value{reflect.ValueOf(ctx), builtRequest})
	var outJSON []byte
	if returnValues[0].CanInterface() {
		outMessage, ok := (returnValues[0].Interface()).(proto.Message)
		if ok {
			outJSON, err = protojson.Marshal(outMessage)
			if err != nil || outJSON != nil && (len(outJSON) == 0 || string(outJSON) == "null" || string(outJSON) == "{}") {
				outJSON = nil
			}
		}
	}
	if returnValues[1].CanInterface() {
		err, _ = returnValues[1].Interface().(error)
	}
	res = &httpapi.Response{
		Payload: outJSON,
	}
	if err == nil {
		res.StatusCode = 200 // TODO: parse status code specifically from outJSON
	} else {
		res.StatusCode = 500
	}
	return res, err
}

// One struct in, stream of structs out
func callStructStream(ctx context.Context, inputJSON []byte, procType reflect.Type, caller reflect.Value) (res *httpapi.Response, err error) {
	// Actually call the inner procedure
	// This is impossible - we need to construct a new type which satisfies the stream interface, but that's not currently possible using go reflection.
	// https://github.com/golang/go/issues/16522 is where discussion on this topic is ongoing, though it started as early as 2012
	// This problem will make all stream interfaces impossible to manage through pure reflection, though it may be possible to require
	// users to pass in their own types that can satisfy all necessary stream interfaces. This will be a really awkward solution, since
	// protobufs don't expose these interfaces by default. If the compiled protobufs are in the same package this is fine, but for separate
	// protobuf packages to the application service you'll need to manually expose those types or create new ones that have to be manually maintained.
	// Everything about this sucks.
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("httpgrpc: Struct In, Stream Out is not yet supported, please manually implement exceptions for endpoint %s", procType))
}

// Stream of structs in, one struct out
func callStreamStruct(ctx context.Context, inputJSON []byte, procType reflect.Type, caller reflect.Value) (res *httpapi.Response, err error) {
	// Actually call the inner procedure
	// This is impossible - we need to construct a new type which satisfies the stream interface, but that's not currently possible using go reflection.
	// https://github.com/golang/go/issues/16522 is where discussion on this topic is ongoing, though it started as early as 2012
	// This problem will make all stream interfaces impossible to manage through pure reflection, though it may be possible to require
	// users to pass in their own types that can satisfy all necessary stream interfaces. This will be a really awkward solution, since
	// protobufs don't expose these interfaces by default. If the compiled protobufs are in the same package this is fine, but for separate
	// protobuf packages to the application service you'll need to manually expose those types or create new ones that have to be manually maintained.
	// Everything about this sucks.
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("httpgrpc: Stream In, Struct Out is not yet supported, please manually implement exceptions for endpoint %s", procType))
}

// Stram of structs in, stream of structs out
func callStreamStream(ctx context.Context, inputJSON []byte, procType reflect.Type, caller reflect.Value) (res *httpapi.Response, err error) {
	// Actually call the inner procedure
	// This is impossible - we need to construct a new type which satisfies the stream interface, but that's not currently possible using go reflection.
	// https://github.com/golang/go/issues/16522 is where discussion on this topic is ongoing, though it started as early as 2012
	// This problem will make all stream interfaces impossible to manage through pure reflection, though it may be possible to require
	// users to pass in their own types that can satisfy all necessary stream interfaces. This will be a really awkward solution, since
	// protobufs don't expose these interfaces by default. If the compiled protobufs are in the same package this is fine, but for separate
	// protobuf packages to the application service you'll need to manually expose those types or create new ones that have to be manually maintained.
	// Everything about this sucks.
	return nil, status.Error(codes.Unimplemented, fmt.Sprintf("httpgrpc: Stream In, Stream Out is not yet supported, please manually implement exceptions for endpoint %s", procType))
}

func parseRequest(req *httpapi.Request) (finalJSON []byte, err error) {
	// First we convert query parameters to a map
	queryMap := map[string]interface{}{}
	for key, value := range req.GetParams() {
		// TODO: don't ignore/overwrite duplicate keys here
		passedValue := ""
		for _, passedValue = range value.GetValues() {
		}
		queryMap[key] = passedValue
	}
	switch req.GetMethod() {
	case httpapi.Method_CONNECT, httpapi.Method_GET, httpapi.Method_HEAD, httpapi.Method_OPTIONS, httpapi.Method_TRACE:
		// No request body, only query params are possible
		if len(queryMap) > 0 {
			finalJSON, err = json.Marshal(queryMap)
			if err != nil {
				err = fmt.Errorf("failed to marshal query parameters to JSON: %v", err)
			}
		}
	case httpapi.Method_DELETE, httpapi.Method_PATCH, httpapi.Method_POST, httpapi.Method_PUT:
		// Merge request body with query params
		bodyJSON := req.GetPayload()
		if bodyJSON != nil && len(queryMap) > 0 {
			var bodyMap map[string]interface{}
			err = json.Unmarshal(bodyJSON, &bodyMap)
			if err != nil {
				err = fmt.Errorf("failed to unmarshall request body JSON: %v", err)
				break
			}
			// Merge both maps, using request body's values on conflict
			mergedMaps := mergemap.Merge(queryMap, bodyMap)
			finalJSON, err = json.Marshal(mergedMaps)
		} else if bodyJSON != nil {
			finalJSON = bodyJSON
		} else if len(queryMap) > 0 {
			finalJSON, err = json.Marshal(queryMap)
		}
	default:
		// Invalid http method
		// It shouldn't be possible to hit this normally, we do validation before we reach this point
		err = fmt.Errorf("invalid http method")
	}
	return
}

func methodToString(in httpapi.Method) (out string, err error) {
	switch in {
	case httpapi.Method_GET:
		out = "GET"
	case httpapi.Method_HEAD:
		out = "HEAD"
	case httpapi.Method_POST:
		out = "POST"
	case httpapi.Method_PUT:
		out = "PUT"
	case httpapi.Method_DELETE:
		out = "DELETE"
	case httpapi.Method_CONNECT:
		out = "CONNECT"
	case httpapi.Method_OPTIONS:
		out = "OPTIONS"
	case httpapi.Method_TRACE:
		out = "TRACE"
	case httpapi.Method_PATCH:
		out = "PATCH"
	}
	if out == "" {
		err = fmt.Errorf("unknown HTTP method")
	}
	return
}

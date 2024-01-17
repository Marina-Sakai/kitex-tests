// Code generated by Kitex v0.8.0. DO NOT EDIT.

package combineservice

import (
	"context"
	"errors"
	"fmt"
	combine_extend "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/combine_extend"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/echo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

type CombineService interface {
	combine_extend.A
	combine_extend.B
}

var serviceMethods = map[string]kitex.MethodInfo{
	"EchoPingPongNew": kitex.NewMethodInfo(
		echoPingPongNewHandler,
		newPingPongOnlyServiceEchoPingPongNewArgs,
		newPingPongOnlyServiceEchoPingPongNewResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"EchoBidirectionalNew": kitex.NewMethodInfo(
		echoBidirectionalNewHandler,
		newStreamOnlyServiceEchoBidirectionalNewArgs,
		newStreamOnlyServiceEchoBidirectionalNewResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingBidirectional),
	),
}

var (
	combineServiceServiceInfo                = NewServiceInfo()
	combineServiceServiceInfoForClient       = NewServiceInfoForClient()
	combineServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return combineServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return combineServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return combineServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(true, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CombineService"
	handlerType := (*CombineService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "combine_extend",
	}
	extra["combine_service"] = true
	extra["combined_service_list"] = []string{"A", "B"}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func echoPingPongNewHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.PingPongOnlyServiceEchoPingPongNewArgs)
	realResult := result.(*echo.PingPongOnlyServiceEchoPingPongNewResult)
	success, err := handler.(echo.PingPongOnlyService).EchoPingPongNew(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPingPongOnlyServiceEchoPingPongNewArgs() interface{} {
	return echo.NewPingPongOnlyServiceEchoPingPongNewArgs()
}

func newPingPongOnlyServiceEchoPingPongNewResult() interface{} {
	return echo.NewPingPongOnlyServiceEchoPingPongNewResult()
}

func echoBidirectionalNewHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st, ok := arg.(*streaming.Args)
	if !ok {
		return errors.New("StreamOnlyService.EchoBidirectionalNew is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &streamOnlyServiceEchoBidirectionalNewServer{st.Stream}
	return handler.(echo.StreamOnlyService).EchoBidirectionalNew(stream)
}

type streamOnlyServiceEchoBidirectionalNewClient struct {
	streaming.Stream
}

func (x *streamOnlyServiceEchoBidirectionalNewClient) Send(m *echo.EchoRequest) error {
	return x.Stream.SendMsg(m)
}
func (x *streamOnlyServiceEchoBidirectionalNewClient) Recv() (*echo.EchoResponse, error) {
	m := new(echo.EchoResponse)
	return m, x.Stream.RecvMsg(m)
}

type streamOnlyServiceEchoBidirectionalNewServer struct {
	streaming.Stream
}

func (x *streamOnlyServiceEchoBidirectionalNewServer) Send(m *echo.EchoResponse) error {
	return x.Stream.SendMsg(m)
}

func (x *streamOnlyServiceEchoBidirectionalNewServer) Recv() (*echo.EchoRequest, error) {
	m := new(echo.EchoRequest)
	return m, x.Stream.RecvMsg(m)
}

func newStreamOnlyServiceEchoBidirectionalNewArgs() interface{} {
	return echo.NewStreamOnlyServiceEchoBidirectionalNewArgs()
}

func newStreamOnlyServiceEchoBidirectionalNewResult() interface{} {
	return echo.NewStreamOnlyServiceEchoBidirectionalNewResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoPingPongNew(ctx context.Context, req1 *echo.EchoRequest) (r *echo.EchoResponse, err error) {
	var _args echo.PingPongOnlyServiceEchoPingPongNewArgs
	_args.Req1 = req1
	var _result echo.PingPongOnlyServiceEchoPingPongNewResult
	if err = p.c.Call(ctx, "EchoPingPongNew", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoBidirectionalNew(ctx context.Context) (StreamOnlyService_EchoBidirectionalNewClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "EchoBidirectionalNew", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &streamOnlyServiceEchoBidirectionalNewClient{res.Stream}
	return stream, nil
}

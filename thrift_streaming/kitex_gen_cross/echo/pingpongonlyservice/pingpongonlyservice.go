// Code generated by Kitex v0.8.0. DO NOT EDIT.

package pingpongonlyservice

import (
	"context"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen_cross/echo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return pingPongOnlyServiceServiceInfo
}

var pingPongOnlyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PingPongOnlyService"
	handlerType := (*echo.PingPongOnlyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"EchoPingPongNew": kitex.NewMethodInfo(echoPingPongNewHandler, newPingPongOnlyServiceEchoPingPongNewArgs, newPingPongOnlyServiceEchoPingPongNewResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "echo",
		"ServiceFilePath": `idl/api.thrift`,
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

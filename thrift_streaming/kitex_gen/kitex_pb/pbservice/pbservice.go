// Code generated by Kitex v0.8.0. DO NOT EDIT.

package pbservice

import (
	"context"
	"errors"
	kitex_pb "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/kitex_pb"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var _ = streaming.KitexUnusedProtection

var serviceMethods = map[string]kitex.MethodInfo{
	"EchoPingPong": kitex.NewMethodInfo(
		echoPingPongHandler,
		newEchoPingPongArgs,
		newEchoPingPongResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	pBServiceServiceInfo                = NewServiceInfo()
	pBServiceServiceInfoForClient       = NewServiceInfoForClient()
	pBServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return pBServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return pBServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return pBServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PBService"
	handlerType := (*kitex_pb.PBService)(nil)
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
		"PackageName": "kitex_pb",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func echoPingPongHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(kitex_pb.Request)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(kitex_pb.PBService).EchoPingPong(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *EchoPingPongArgs:
		success, err := handler.(kitex_pb.PBService).EchoPingPong(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*EchoPingPongResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newEchoPingPongArgs() interface{} {
	return &EchoPingPongArgs{}
}

func newEchoPingPongResult() interface{} {
	return &EchoPingPongResult{}
}

type EchoPingPongArgs struct {
	Req *kitex_pb.Request
}

func (p *EchoPingPongArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(kitex_pb.Request)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *EchoPingPongArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *EchoPingPongArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *EchoPingPongArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *EchoPingPongArgs) Unmarshal(in []byte) error {
	msg := new(kitex_pb.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EchoPingPongArgs_Req_DEFAULT *kitex_pb.Request

func (p *EchoPingPongArgs) GetReq() *kitex_pb.Request {
	if !p.IsSetReq() {
		return EchoPingPongArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EchoPingPongArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EchoPingPongArgs) GetFirstArgument() interface{} {
	return p.Req
}

type EchoPingPongResult struct {
	Success *kitex_pb.Response
}

var EchoPingPongResult_Success_DEFAULT *kitex_pb.Response

func (p *EchoPingPongResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(kitex_pb.Response)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *EchoPingPongResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *EchoPingPongResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *EchoPingPongResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *EchoPingPongResult) Unmarshal(in []byte) error {
	msg := new(kitex_pb.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EchoPingPongResult) GetSuccess() *kitex_pb.Response {
	if !p.IsSetSuccess() {
		return EchoPingPongResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EchoPingPongResult) SetSuccess(x interface{}) {
	p.Success = x.(*kitex_pb.Response)
}

func (p *EchoPingPongResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoPingPongResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoPingPong(ctx context.Context, Req *kitex_pb.Request) (r *kitex_pb.Response, err error) {
	var _args EchoPingPongArgs
	_args.Req = Req
	var _result EchoPingPongResult
	if err = p.c.Call(ctx, "EchoPingPong", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
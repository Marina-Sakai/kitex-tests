// Code generated by Kitex v0.8.0. DO NOT EDIT.
package streamonlyservicechild

import (
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen_cross/echo"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler echo.StreamOnlyServiceChild, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

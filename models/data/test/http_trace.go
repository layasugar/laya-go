package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/dao/cal/http_test"
)

func HttpToHttpTraceTest(ctx *laya.WebContext) (*Rsp, error) {
	d, err := http_test.HttpToHttpTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Code,
	}

	return &res, nil
}

func HttpToRpcTraceTest(ctx *laya.WebContext) (*Rsp, error) {
	d, err := http_test.HttpToGrpcTraceTest(ctx)
	if err != nil {
		return nil, err
	}

	var res = Rsp{
		Code: d.Message,
	}

	return &res, nil
}

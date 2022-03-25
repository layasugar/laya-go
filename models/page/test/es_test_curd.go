package test

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/models/data/test"
	"log"
)

func esTestCurd(ctx *laya.WebContext) {
	eid, err := test.EsUserCreate(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("es数据创建成功, %s", eid)

	err = test.EsUserUpdate(ctx, eid)
	if err != nil {
		panic(err)
	}
	log.Printf("es数据更新成功")

	res, err := test.EsUserSelect(ctx, eid)
	if err != nil {
		panic(err)
	}
	log.Printf("es数据查询成功, %v", res)

	err = test.EsUserDel(ctx, eid)
	if err != nil {
		panic(err)
	}
	log.Print("es数据删除成功")
}

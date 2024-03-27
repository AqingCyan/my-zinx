package znet

import "my-zinx/ziface"

// BaseRouter 实现 router 时，先嵌入这个 BaseRouter 基类，然后根据需要对这个基类的方法进行重写就好了
type BaseRouter struct {
}

/**
 * 这里之所以 BaseRouter 的方法都为空
 * 是因为有的 Router 不希望有 PreHandle、PostHandle 这两个业务
 * 所以 Router 全部继承 BaseRouter 的好处就是，不需要实现 PreHandle、PostHandle
 */

func (br *BaseRouter) PreHandle(request ziface.IRequest) {
	//TODO implement me
	panic("implement me")
}

func (br *BaseRouter) Handle(request ziface.IRequest) {
	//TODO implement me
	panic("implement me")
}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {
	//TODO implement me
	panic("implement me")
}

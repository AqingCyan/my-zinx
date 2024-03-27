package ziface

// IRouter 路由的抽象接口，路由里的数据都是 IRequest
type IRouter interface {
	// PreHandle 在处理 conn 业务之前的 Hook
	PreHandle(request IRequest)

	// Handle 在处理 conn 业务时的钩子 Hook
	Handle(request IRequest)

	// PostHandle 在处理 conn 业务之后的钩子 Hook
	PostHandle(request IRequest)
}

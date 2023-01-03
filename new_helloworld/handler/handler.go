package handler

// 前缀解决名称冲突问题
const HelloServiceName = "handler/HelloService"

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

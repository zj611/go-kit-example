package pkg

import "fmt"

type MyServer interface {
	Hello(name string) string
	Bye(a int,b int) string
}

//用于实现上面定义的接口
type Server struct {

}

func (s Server) Hello(name string) string  {
	return fmt.Sprintf("Hello! %s",name)
}

func (s Server) Bye(a int,b int) string  {
	return fmt.Sprintf("Bye! %d",a + b)
}




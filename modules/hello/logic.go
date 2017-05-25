package hello

type Logic struct {
}

func (l Logic) helloFunction(name string) HelloRespone {
	return HelloRespone{Name: name, Message: "Hello"}
}

package router

type http struct {
	routerRegisterHttps []RouterRegisterHttp
}

func (h *http) register(registerName string, routerRegister routerRegister) error {
	h.routerRegisterHttps = append(h.routerRegisterHttps, routerRegister.routerRegisterHttp())
	return nil
}

package ginger

type RegisterHandler func(engine Ginger)

func (e *gingerEngine) RegisterHandler(handler RegisterHandler) {
	handler(e)
}

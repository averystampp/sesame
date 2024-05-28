package sesame

var middleWare map[int]Handler

func RegisterMiddleware(h Handler) {
	if middleWare == nil {
		middleWare = make(map[int]Handler)
	}
	middleWare[len(middleWare)+1] = h
}

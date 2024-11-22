package skelmi

// Block represents a logical block, e.g. session, user etc.
type Block interface {
	GetMiddlewares() []Middleware
	GetHandlers() []Handler
}

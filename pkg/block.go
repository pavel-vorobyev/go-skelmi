package skelmi

// Block represents a logical block, e.g. session, user etc.
type Block interface {
	GetPath() string
	GetMiddlewares() []Middleware
	GetHandlers() []Handler
}

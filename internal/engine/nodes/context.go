package nodes

// Context is the graph run state a node may read or write.
// engine.Context implements this; nodes must not import engine.
type Context interface {
	SetVariable(key string, value any)
	GetVariable(key string) (any, bool)
	DeleteVariable(key string)
}

package engine

import "github.com/author-spirit/lamrim/internal/engine/nodes"

var _ nodes.Context = (*Context)(nil)

// Context is shared across a graph run.
// Variables are explicit key/value state (set/delete).
// Steps hold each node's result, keyed by node ID.
type Context struct {
	Variables map[string]any    `json:"variables"`
	Steps     map[string]Result `json:"steps"`
	Last      Result            `json:"last"`
}

func NewContext() *Context {
	return &Context{
		Variables: make(map[string]any),
		Steps:     make(map[string]Result),
	}
}

func (c *Context) SetVariable(key string, value any) {
	if c.Variables == nil {
		c.Variables = make(map[string]any)
	}
	c.Variables[key] = value
}

func (c *Context) GetVariable(key string) (any, bool) {
	if c.Variables == nil {
		return nil, false
	}
	value, ok := c.Variables[key]
	return value, ok
}

func (c *Context) DeleteVariable(key string) {
	delete(c.Variables, key)
}

func (c *Context) RecordStep(nodeID string, result Result) {
	if c.Steps == nil {
		c.Steps = make(map[string]Result)
	}
	c.Steps[nodeID] = result
	c.Last = result
}

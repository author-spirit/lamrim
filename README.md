## Learning
Create a project
`go mod init github.com/author-spirit/lamrim`

Run
`go run ./cmd/lamrim -list`
`go run ./cmd/lamrim -workflow first`

Start server + web client
`go run ./cmd/lamrim -serve`
Open http://localhost:8080

Workflow projects live in `workflows/` (`config.json` + `main.js`).
Web client lives in `client/`.


## Graph Representation
```json
{
 "nodes": {
	"node_6gbf6kjrg": {
		"type": "node_type",
		"config": {
			"name": "node_name",
		}
	}
 },
 "edges": {
	"node_6gbf6kjrg": ["node_mmio6712", "node_12hd7oqv"]
 }
}
```

- Better traversal time - O(v+e)
- Find - O(1)
- Delete - O(1)


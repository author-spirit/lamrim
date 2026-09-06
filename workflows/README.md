# Workflows

Each subfolder is a workflow project. No Go files here.

## Project layout

```text
workflows/
  first/
    config.json   # metadata + graph definition
    main.js       # user-provided executable code
```

## Run

From the repo root:

```bash
go run ./cmd/lamrim -list
go run ./cmd/lamrim -workflow first
```

`main.js` runs with Node.js after the graph executes. It receives:

- `LAMRIM_WORKFLOW` — project folder name
- `LAMRIM_CONTEXT` — JSON execution context (variables, steps)

## config.json

```json
{
  "name": "first",
  "description": "Example workflow",
  "nodes": [
    {
      "id": "start",
      "name": "start",
      "type": "trigger",
      "input": { "type": "event" }
    }
  ],
  "edges": {
    "start": ["init"]
  }
}
```

Node `id` is optional; `name` is used when omitted.

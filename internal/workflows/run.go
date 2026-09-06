package workflows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/author-spirit/lamrim/internal/engine"
)

func RunScript(project Project, graph *engine.Graph) error {
	scriptPath := filepath.Join(project.Dir, ScriptFile)
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return nil
	}

	contextJSON, err := json.Marshal(graph.Context)
	if err != nil {
		return fmt.Errorf("context json: %w", err)
	}

	cmd := exec.Command("node", ScriptFile)
	cmd.Dir = project.Dir
	cmd.Env = append(os.Environ(),
		"LAMRIM_WORKFLOW="+project.Name,
		"LAMRIM_CONTEXT="+string(contextJSON),
	)

	var stderr bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return fmt.Errorf("main.js: %w: %s", err, stderr.String())
		}
		return fmt.Errorf("main.js: %w", err)
	}
	return nil
}

package workflows

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigFile = "config.json"
	ScriptFile = "main.js"
)

// Project is a runnable workflow folder under workflows/.
type Project struct {
	Name string
	Dir  string
}

// Discover returns runnable workflow projects from root (e.g. workflows/).
// A project is a subdirectory that contains config.json.
func Discover(root string) ([]Project, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	projects := make([]Project, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dir := filepath.Join(root, entry.Name())
		if _, err := os.Stat(filepath.Join(dir, ConfigFile)); err != nil {
			continue
		}

		projects = append(projects, Project{
			Name: entry.Name(),
			Dir:  dir,
		})
	}

	return projects, nil
}

// Find returns one project by folder name.
func Find(root, name string) (Project, error) {
	projects, err := Discover(root)
	if err != nil {
		return Project{}, err
	}
	for _, project := range projects {
		if project.Name == name {
			return project, nil
		}
	}
	return Project{}, fmt.Errorf("workflow %q not found", name)
}

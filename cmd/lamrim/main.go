package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/author-spirit/lamrim/internal/server"
	"github.com/author-spirit/lamrim/internal/workflows"
)

const (
	defaultWorkflowsDir = "workflows"
	defaultClientDir    = "client"
	defaultAddr         = ":8080"
)

func main() {
	serve := flag.Bool("serve", false, "start HTTP server")
	addr := flag.String("addr", defaultAddr, "server listen address")
	list := flag.Bool("list", false, "list runnable workflows")
	name := flag.String("workflow", "", "workflow project to run")
	dir := flag.String("workflows-dir", defaultWorkflowsDir, "workflows root directory")
	clientDir := flag.String("client-dir", defaultClientDir, "static client directory")
	flag.Parse()

	if *serve {
		runServer(*addr, *dir, *clientDir)
		return
	}

	runCLI(*list, *name, *dir)
}

func runServer(addr, workflowsDir, clientDir string) {
	srv := server.New(workflowsDir, clientDir)
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.Handler()); err != nil {
		log.Fatal(err)
	}
}

func runCLI(list bool, name, dir string) {
	projects, err := workflows.Discover(dir)
	if err != nil {
		log.Fatal(err)
	}
	if len(projects) == 0 {
		log.Fatalf("no runnable workflows found in %q", dir)
	}

	if list || name == "" {
		printProjects(projects)
		if name == "" && !list {
			fmt.Fprintf(os.Stderr, "use -workflow <name> to run one\n")
		}
		return
	}

	project, err := workflows.Find(dir, name)
	if err != nil {
		log.Fatal(err)
	}

	graph, err := workflows.LoadGraph(project)
	if err != nil {
		log.Fatal(err)
	}

	if err := workflows.Run(project, graph); err != nil {
		log.Fatal(err)
	}

	fmt.Println(graph.RepresentGraph())
}

func printProjects(projects []workflows.Project) {
	names := make([]string, len(projects))
	for i, project := range projects {
		names[i] = project.Name
	}
	fmt.Println(strings.Join(names, "\n"))
}

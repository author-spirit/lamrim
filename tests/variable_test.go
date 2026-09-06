package engine_test

import (
	"testing"

	"github.com/author-spirit/lamrim/internal/engine"
)

func TestNodeVariableAPI(t *testing.T) {
	g := engine.NewGraph("vars")

	init := g.AddNode("init", engine.Variable)
	if err := init.SetVariable("user", "martha"); err != nil {
		t.Fatal(err)
	}
	if err := init.SetVariable("n", 1); err != nil {
		t.Fatal(err)
	}

	one := g.AddNode("one", engine.Variable)
	if err := one.SetVariables(`{"city":"goa","last_name":"mayfield"}`); err != nil {
		t.Fatal(err)
	}

	drop := g.AddNode("drop", engine.Variable)
	if err := drop.DeleteVariable("n"); err != nil {
		t.Fatal(err)
	}

	if got, ok := init.GetVariable("user"); !ok || got != "martha" {
		t.Fatalf("configured user: got %v ok=%v", got, ok)
	}

	if err := g.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	if got, ok := g.Context.GetVariable("user"); !ok || got != "martha" {
		t.Fatalf("user: got %v ok=%v", got, ok)
	}
	if got, ok := g.Context.GetVariable("last_name"); !ok || got != "mayfield" {
		t.Fatalf("last_name in context: got %v ok=%v", got, ok)
	}
	if _, ok := g.Context.GetVariable("n"); ok {
		t.Fatal("n should have been deleted")
	}
}

func TestExecuteNodeRunsTrigger(t *testing.T) {
	g := engine.NewGraph("trig")
	node := g.AddNode("start", engine.Trigger)
	if err := node.SetTrigger("time"); err != nil {
		t.Fatal(err)
	}

	result, err := node.ExecuteNode(g.Context)
	if err != nil {
		t.Fatalf("execute trigger: %v", err)
	}
	if result.Output["type"] != "time" {
		t.Fatalf("output: %#v", result.Output)
	}
}

func TestSetCondition(t *testing.T) {
	g := engine.NewGraph("cond")
	node := g.AddNode("check", engine.Condition)
	if err := node.SetCondition("user == martha"); err != nil {
		t.Fatal(err)
	}

	result, err := node.ExecuteNode(g.Context)
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["expression"] != "user == martha" {
		t.Fatalf("output: %#v", result.Output)
	}
}

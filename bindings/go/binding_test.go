package tree_sitter_kotlin_test

import (
	"testing"

	tree_sitter_kotlin "github.com/rover-app/tree-sitter-kotlin/bindings/go"
	tree_sitter "github.com/smacker/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_kotlin.Language())
	if language == nil {
		t.Errorf("Error loading Kotlin grammar")
	}
}

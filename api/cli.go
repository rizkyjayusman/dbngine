package api

import (
	"bufio"
	"dbngin3/engine"
	"dbngin3/parser"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	lexer            *parser.Lexer
	parser           *parser.Parser
	semanticAnalyzer *parser.SelectSemanticAnalyzer
	queryOptimizer   *parser.SelectQueryOptimizer
}

func NewCLI() *CLI {
	schema := engine.NewSchemaManager()
	return &CLI{
		lexer:            &parser.Lexer{},
		parser:           &parser.Parser{},
		semanticAnalyzer: &parser.SelectSemanticAnalyzer{Schema: schema},
		queryOptimizer:   &parser.SelectQueryOptimizer{Schema: schema},
	}
}

func (cli *CLI) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple DBEngine CLI (Type 'exit' to quit)")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		query := strings.TrimSpace(scanner.Text())
		if query == "exit" {
			fmt.Println("Exiting...")
			break
		}
		result := cli.ExecuteQuery(query)
		fmt.Println(result)
	}
}

func (cli *CLI) ExecuteQuery(query string) error {
	if err := cli.lexer.SetInput(query); err != nil {
		return err
	}

	var tokens []parser.Token
	tokens, err := cli.lexer.Tokenize()
	if err != nil {
		return err
	}

	if err = cli.parser.SetToken(tokens); err != nil {
		return err
	}

	nodes, err := cli.parser.Parse()
	if err != nil {
		return err
	}

	if node, ok := nodes.(*parser.SelectStatement); ok {
		if err := cli.semanticAnalyzer.Analyze(node); err != nil {
			return err
		}

		if err := cli.queryOptimizer.Optimize(node); err != nil {
			return err
		}

		fmt.Println("Select Statement")
		fmt.Println("Table: ", node.Table)
		fmt.Println("Columns: ", node.Columns)
	} else {
		fmt.Println("Invalid Syntax")
	}
	return nil
}

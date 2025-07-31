package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/c-bata/go-prompt"
	"github.com/dradtke/go-wren"
	"github.com/fatih/color"
)

const (
	promptPrefix    = "wren> "
	multilinePrefix = "... "
	welcomeMessage  = `
╭─────────────────────────────────────────╮
│           Wren Interactive REPL         │
│                                         │
│  Type Wren code and press Enter to      │
│  execute. Use .help for commands.       │
│                                         │
│  Multi-line input: Use \ at end of      │
│  line to continue on next line.         │
╰─────────────────────────────────────────╯
`
)

type REPL struct {
	vm        *wren.VM
	multiLine bool
	buffer    strings.Builder
	history   []string
	colors    struct {
		prompt  *color.Color
		output  *color.Color
		error   *color.Color
		success *color.Color
		info    *color.Color
		keyword *color.Color
	}
}

func NewREPL() *REPL {
	r := &REPL{
		vm:      wren.NewVM(),
		history: make([]string, 0),
	}

	// Initialize colors
	r.colors.prompt = color.New(color.FgCyan, color.Bold)
	r.colors.output = color.New(color.FgWhite)
	r.colors.error = color.New(color.FgRed, color.Bold)
	r.colors.success = color.New(color.FgGreen)
	r.colors.info = color.New(color.FgYellow)
	r.colors.keyword = color.New(color.FgMagenta, color.Bold)

	return r
}

func (r *REPL) completer(d prompt.Document) []prompt.Suggest {
	keywords := []prompt.Suggest{
		{Text: "System.print", Description: "Print to console"},
		{Text: "if", Description: "Conditional statement"},
		{Text: "else", Description: "Else clause"},
		{Text: "while", Description: "While loop"},
		{Text: "for", Description: "For loop"},
		{Text: "class", Description: "Class definition"},
		{Text: "var", Description: "Variable declaration"},
		{Text: "fun", Description: "Function definition"},
		{Text: "return", Description: "Return statement"},
		{Text: "import", Description: "Import statement"},
		{Text: "break", Description: "Break statement"},
		{Text: "continue", Description: "Continue statement"},
		{Text: "true", Description: "Boolean true"},
		{Text: "false", Description: "Boolean false"},
		{Text: "null", Description: "Null value"},
		{Text: "super", Description: "Superclass reference"},
		{Text: "this", Description: "Current instance"},
		{Text: "is", Description: "Type checking"},
		{Text: "and", Description: "Logical AND"},
		{Text: "or", Description: "Logical OR"},
		{Text: "not", Description: "Logical NOT"},
		// Fiber/concurrency keywords
		{Text: "Fiber.new", Description: "Create new fiber"},
		{Text: "Fiber.yield", Description: "Yield from fiber"},
		{Text: "Fiber.abort", Description: "Abort fiber with error"},
		{Text: "fiber.call", Description: "Call/resume fiber"},
		{Text: "fiber.try", Description: "Try fiber call (catch errors)"},
		{Text: "fiber.transfer", Description: "Transfer to fiber"},
		{Text: "fiber.isDone", Description: "Check if fiber is done"},
		{Text: "fiber.error", Description: "Get fiber error message"},
	}

	// Add special REPL commands
	replCommands := []prompt.Suggest{
		{Text: "help", Description: "Show help message"},
		{Text: "exit", Description: "Exit the REPL"},
		{Text: "quit", Description: "Exit the REPL"},
		{Text: "clear", Description: "Clear the screen"},
		{Text: "history", Description: "Show command history"},
		{Text: "reset", Description: "Reset the VM state"},
	}

	text := d.GetWordBeforeCursor()
	if strings.HasPrefix(text, ".") {
		return prompt.FilterHasPrefix(replCommands, text, true)
	}

	return prompt.FilterHasPrefix(keywords, text, true)
}

func (r *REPL) executor(input string) {
	input = strings.TrimSpace(input)

	// Handle empty input
	if input == "" {
		return
	}

	// Handle REPL commands
	if strings.HasPrefix(input, ".") {
		r.handleCommand(input)
		return
	}

	// Handle multi-line continuation
	if strings.HasSuffix(input, "\\") {
		r.multiLine = true
		r.buffer.WriteString(strings.TrimSuffix(input, "\\"))
		r.buffer.WriteString("\n")
		return
	}

	// Get complete code (either single line or accumulated multi-line)
	var code string
	if r.multiLine {
		r.buffer.WriteString(input)
		code = r.buffer.String()
		r.buffer.Reset()
		r.multiLine = false
	} else {
		code = input
	}

	// Add to history
	r.history = append(r.history, code)

	// Display the code with syntax highlighting
	r.displayCode(code)

	// Execute the Wren code
	err := r.vm.Interpret(code)
	if err != nil {
		r.colors.error.Printf("Error: %v\n", err)
	}
}

func (r *REPL) displayCode(code string) {
	// Use chroma for syntax highlighting
	err := quick.Highlight(os.Stdout, code, "wren", "terminal", "monokai")
	if err != nil {
		// Fallback to plain text if highlighting fails
		fmt.Print(code)
	}
	fmt.Println()
}

func (r *REPL) handleCommand(cmd string) {
	switch cmd {
	case "help":
		r.showHelp()
	case "exit", ".quit":
		r.colors.info.Println("Goodbye!")
		os.Exit(0)
	case "clear":
		r.clearScreen()
	case "history":
		r.showHistory()
	case "reset":
		r.resetVM()
	default:
		r.colors.error.Printf("Unknown command: %s\n", cmd)
		r.colors.info.Println("Type .help for available commands")
	}
}

func (r *REPL) showHelp() {
	help := `
Available Commands:
  .help     - Show this help message
  .exit     - Exit the REPL (Ctrl+C also works)
  .quit     - Exit the REPL (Ctrl+C also works)
  .clear    - Clear the screen
  .history  - Show command history
  .reset    - Reset the VM state

Wren Language Features:
  - Variables: var name = "value"
  - Functions: fun greet(name) { System.print("Hello " + name) }
  - Classes: class Person { construct new(name) { _name = name } }
  - Loops: for (i in 1..10) { System.print(i) }
  - Conditionals: if (condition) { ... } else { ... }

Multi-line Input:
  Use \ at the end of a line to continue on the next line.
  Example:
    wren> if (true) { \
    ...   System.print("Hello") \
    ... }

Examples:
  System.print("Hello, World!")
  var x = 42
  System.print(x * 2)

  class Greeter {
    construct new(name) { _name = name }
    greet() { System.print("Hello, " + _name + "!") }
  }
  var g = Greeter.new("Wren")
  g.greet()

Fiber/Concurrency Examples:
  // Basic fiber with yielding
  var fiber = Fiber.new {
    System.print("Fiber: Starting")
    Fiber.yield("yielded-value")
    System.print("Fiber: Resumed")
  }
  var result = fiber.call()  // "yielded-value"
  fiber.call()               // completes fiber

  // Producer-consumer pattern
  var producer = Fiber.new {
    for (i in 1..3) {
      Fiber.yield(i)
    }
  }
  while (!producer.isDone) {
    System.print(producer.call())
  }

  // Error handling
  var errorFiber = Fiber.new {
    null.badMethod  // Will error
  }
  System.print(errorFiber.try())  // Catches error
`
	r.colors.info.Print(help)
}

func (r *REPL) clearScreen() {
	// ANSI escape sequence to clear screen
	fmt.Print("\033[2J\033[H")
	r.colors.success.Print(welcomeMessage)
}

func (r *REPL) showHistory() {
	if len(r.history) == 0 {
		r.colors.info.Println("No history available")
		return
	}

	r.colors.info.Println("\nCommand History:")
	for i, cmd := range r.history {
		r.colors.output.Printf("%3d: %s\n", i+1, cmd)
	}
	fmt.Println()
}

func (r *REPL) resetVM() {
	r.vm = wren.NewVM()
	r.colors.success.Println("VM state has been reset")
}

func (r *REPL) getPrompt() (string, bool) {
	if r.multiLine {
		return multilinePrefix, true
	}
	return promptPrefix, true
}

func (r *REPL) Run() {
	// Check if we're in an interactive terminal
	fileInfo, _ := os.Stdin.Stat()
	isInteractive := fileInfo.Mode()&os.ModeCharDevice != 0

	if !isInteractive {
		// Non-interactive mode - read from stdin line by line
		r.runNonInteractive()
		return
	}

	// Print welcome message
	r.colors.success.Print(welcomeMessage)
	r.colors.info.Println("Type .help for available commands or start typing Wren code!")
	fmt.Println()

	// Set up the prompt
	p := prompt.New(
		r.executor,
		r.completer,
		prompt.OptionTitle("Wren REPL"),
		prompt.OptionPrefix(promptPrefix),
		prompt.OptionLivePrefix(r.getPrompt),
		prompt.OptionMaxSuggestion(16),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionWordSeparator(" \t\n.,;:(){}[]"),
		prompt.OptionInputTextColor(prompt.Cyan),
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionDescriptionBGColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.White),
	)

	p.Run()
}

func (r *REPL) runNonInteractive() {
	// For non-interactive mode, read all input and execute as one program
	var buffer strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		// Skip comment-only lines that start with //
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && !strings.HasPrefix(trimmed, "//") {
			buffer.WriteString(line)
			buffer.WriteString("\n")
		} else if trimmed == "" {
			buffer.WriteString("\n") // Preserve empty lines for formatting
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	// Execute the complete program
	program := buffer.String()
	if strings.TrimSpace(program) != "" {
		err := r.vm.Interpret(program)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}

func checkWrenLibrary() {
	// Simple test to see if Wren library is available
	vm := wren.NewVM()
	err := vm.Interpret(`System.print("Wren library check: OK")`)
	if err != nil {
		color.Red("Error: Wren library not properly initialized: %v", err)
		color.Yellow("\nTroubleshooting:")
		color.Yellow("1. Make sure you've built the Wren library first:")
		color.Yellow("   cd wren/projects/make && make")
		color.Yellow("2. Ensure libwren.a exists in wren/lib/")
		color.Yellow("3. Try running: go build in the project root")
		os.Exit(1)
	}
}

func main() {
	// Check if Wren library is properly set up
	checkWrenLibrary()

	// Handle command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "-h":
			fmt.Println("Wren Interactive REPL")
			fmt.Println("\nUsage:")
			fmt.Println("  wren-repl          Start interactive REPL")
			fmt.Println("  wren-repl --help   Show this help")
			fmt.Println("  wren-repl --version Show version")
			return
		case "--version", "-v":
			fmt.Println("Wren REPL v1.0.0")
			fmt.Println("Built with go-wren bindings")
			return
		default:
			color.Red("Unknown option: %s", os.Args[1])
			color.Yellow("Use --help for usage information")
			os.Exit(1)
		}
	}

	// Create and run REPL
	repl := NewREPL()
	repl.Run()
}

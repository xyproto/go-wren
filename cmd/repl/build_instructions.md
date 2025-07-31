# Building and Using the Wren REPL

## Prerequisites

1. Build the Wren library first:
```bash
cd wren/projects/make && make
```

2. Verify the library files exist:
```bash
ls wren/lib/
# Should show: libwren.a libwren.so
```

## Building the REPL

```bash
go mod tidy  # Download dependencies
go build -o wren-repl ./cmd/repl
```

## Usage Examples

### Interactive Mode
```bash
./wren-repl
```

This starts the full interactive REPL with:
- Colorful syntax highlighting
- Tab completion for Wren keywords
- Command history
- Multi-line input support
- Built-in help system

### Non-Interactive Mode
```bash
# Execute a Wren file
./wren-repl < script.wren

# Execute from pipe
echo 'System.print("Hello!")' | ./wren-repl

# Execute inline
./wren-repl <<< 'var x = 42; System.print(x)'
```

### Test Examples

Run the included test:
```bash
./wren-repl < cmd/repl/wren_test.wren
```

## REPL Features Demonstrated

✅ **Colorful Interface**: Welcome screen, colored prompts, and syntax highlighting  
✅ **Tab Completion**: Intelligent completion for Wren keywords and REPL commands  
✅ **Multi-line Input**: Use `\` at line end to continue input  
✅ **Command History**: Navigate through previous commands  
✅ **Built-in Commands**: `.help`, `.exit`, `.clear`, `.history`, `.reset`  
✅ **Error Handling**: Clear error messages for syntax and runtime errors  
✅ **Non-interactive Mode**: Supports pipes and file input  
✅ **Wren Integration**: Full integration with go-wren VM

## Architecture

The REPL consists of:
- **main.go**: Main REPL implementation with interactive and non-interactive modes
- **Dependencies**: 
  - `github.com/c-bata/go-prompt`: Interactive prompt library
  - `github.com/fatih/color`: Terminal colors
  - `github.com/alecthomas/chroma`: Syntax highlighting
- **Integration**: Uses the go-wren bindings to execute Wren code

## Example Session

```
wren> System.print("Hello, World!")
Hello, World!

wren> var greeting = "Hi there"
wren> System.print(greeting + "!")
Hi there!

wren> class Person { \
...     construct new(name) { _name = name } \
...     greet() { System.print("Hello, I'm " + _name) } \
... }
wren> var p = Person.new("Alice")
wren> p.greet()
Hello, I'm Alice

wren> .help
# Shows comprehensive help with examples

wren> .exit
Goodbye!
```

The REPL successfully provides a colorful, ergonomic interface for interactive Wren development!
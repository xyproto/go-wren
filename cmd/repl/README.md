# Wren Interactive REPL

A colorful and ergonomic Read-Eval-Print Loop (REPL) for the Wren scripting language, built using the go-wren bindings.

## Features

- 🎨 **Colorful Output**: Syntax highlighting and colored prompts/messages
- 🚀 **Ergonomic Interface**: Tab completion, command history, and multi-line input support
- 📚 **Interactive Help**: Built-in help system with examples and documentation
- 🔄 **Multi-line Support**: Continue code across multiple lines using `\`
- 💾 **Command History**: Navigate through previous commands
- 🧹 **State Management**: Reset VM state or clear screen as needed
- 📝 **Non-interactive Mode**: Works with pipes and scripts

## Building

Make sure the Wren library is built first:

```bash
cd wren/projects/make && make
```

Then build the REPL:

```bash
go build -o wren-repl ./cmd/repl
```

## Usage

### Interactive Mode

Simply run the REPL to start an interactive session:

```bash
./wren-repl
```

You'll see a welcome screen and can start typing Wren code immediately:

```
╭─────────────────────────────────────────╮
│           Wren Interactive REPL         │
│                                         │
│  Type Wren code and press Enter to     │
│  execute. Use .help for commands.      │
│                                         │
│  Multi-line input: Use \ at end of     │
│  line to continue on next line.        │
╰─────────────────────────────────────────╯

wren> System.print("Hello, World!")
Hello, World!

wren> var name = "Wren"
wren> System.print("Hello, " + name)
Hello, Wren
```

### Non-interactive Mode

The REPL also works with pipes and scripts:

```bash
echo 'System.print("Hello from pipe")' | ./wren-repl
```

Or execute a file:

```bash
./wren-repl < my_script.wren
```

## REPL Commands

The REPL provides several built-in commands:

| Command | Description |
|---------|-------------|
| `.help` | Show help message with examples |
| `.exit` or `.quit` | Exit the REPL |
| `.clear` | Clear the screen |
| `.history` | Show command history |
| `.reset` | Reset the VM state |

## Multi-line Input

Use `\` at the end of a line to continue input on the next line:

```wren
wren> if (true) { \
...   System.print("This is a multi-line") \
...   System.print("conditional statement") \
... }
This is a multi-line
conditional statement
```

## Tab Completion

The REPL provides intelligent tab completion for:

- Wren language keywords (`if`, `else`, `while`, `for`, `class`, etc.)
- Built-in functions (`System.print`, etc.)
- REPL commands (`.help`, `.exit`, etc.)

Start typing and press Tab to see available completions.

## Examples

### Basic Variables and Math

```wren
wren> var x = 10
wren> var y = 20
wren> System.print(x + y)
30
```

### Functions

```wren
wren> fun greet(name) { \
...   System.print("Hello, " + name + "!") \
... }
wren> greet("Wren")
Hello, Wren!
```

### Classes

```wren
wren> class Person { \
...   construct new(name) { \
...     _name = name \
...   } \
...   greet() { \
...     System.print("Hi, I'm " + _name) \
...   } \
... }
wren> var person = Person.new("Alice")
wren> person.greet()
Hi, I'm Alice
```

### Loops and Ranges

```wren
wren> for (i in 1..5) { \
...   System.print("Count: " + i.toString) \
... }
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
```

### Fibers and Concurrency

Wren supports cooperative multitasking through fibers:

```wren
wren> var fiber = Fiber.new { \
...   System.print("Fiber: Starting") \
...   Fiber.yield("yielded-value") \
...   System.print("Fiber: Resumed") \
... }
wren> var result = fiber.call()
Fiber: Starting
wren> System.print(result)
yielded-value
wren> fiber.call()
Fiber: Resumed
```

#### Producer-Consumer Pattern

```wren
wren> var producer = Fiber.new { \
...   for (i in 1..3) { \
...     System.print("Producing " + i.toString) \
...     Fiber.yield(i) \
...   } \
... }
wren> while (!producer.isDone) { \
...   var item = producer.call() \
...   if (item != null) System.print("Consumed " + item.toString) \
... }
Producing 1
Consumed 1
Producing 2
Consumed 2
Producing 3
Consumed 3
```

#### Error Handling in Fibers

```wren
wren> var errorFiber = Fiber.new { \
...   null.badMethod \
... }
wren> System.print(errorFiber.try())
Null does not implement 'badMethod'.
wren> System.print(errorFiber.isDone)
true
```

#### Fiber Transfer (Cooperative Multitasking)

```wren
wren> var fiberA = Fiber.new { System.print("A") }
wren> var fiberB = Fiber.new { \
...   System.print("B before") \
...   fiberA.transfer() \
...   System.print("B after") \
... }
wren> fiberB.transfer()
B before
A
```

## Command Line Options

```bash
./wren-repl --help      # Show help message
./wren-repl --version   # Show version information
```

## Error Handling

The REPL provides clear error messages for syntax and runtime errors:

```wren
wren> var x = 
Error: compilation error: main:1: Error: Expected expression.

wren> nonexistent()
Error: compilation error: main:1: Error: Undefined variable 'nonexistent'.
```

## Customization

The REPL uses the following libraries for enhanced functionality:

- **github.com/c-bata/go-prompt**: Interactive prompt with completion and history
- **github.com/fatih/color**: Colorful terminal output
- **github.com/alecthomas/chroma**: Syntax highlighting

Colors and prompts are customizable by modifying the source code in `main.go`.

## Troubleshooting

If you encounter issues:

1. **"Wren library not properly initialized"**: Make sure you've built the Wren library first:
   ```bash
   cd wren/projects/make && make
   ```

2. **Build errors**: Ensure you have Go 1.13+ and run `go mod tidy` to update dependencies.

3. **Interactive mode issues**: The REPL automatically detects if you're in a terminal. For scripting, it will use non-interactive mode automatically.

## License

This REPL tool is part of the go-wren bindings project and follows the same license as the parent project.
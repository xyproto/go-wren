// Test script for Wren REPL
// This file demonstrates various Wren language features

// Basic output
System.print("=== Basic Output ===")
System.print("Hello, Wren REPL!")

// Variables
System.print("\n=== Variables ===")
var message = "Variables work!"
var number = 42
System.print(message)
System.print("Number: " + number.toString)

// Math operations
System.print("\n=== Math Operations ===")
var a = 10
var b = 5
System.print("Addition: " + (a + b).toString)
System.print("Subtraction: " + (a - b).toString)
System.print("Multiplication: " + (a * b).toString)
System.print("Division: " + (a / b).toString)

// Conditionals
System.print("\n=== Conditionals ===")
var x = 15
if (x > 10) {
  System.print("x is greater than 10")
} else {
  System.print("x is not greater than 10")
}

// Loops
System.print("\n=== Loops ===")
System.print("Counting from 1 to 5:")
for (i in 1..5) {
  System.print("  " + i.toString)
}

// Functions
System.print("\n=== Functions ===")
fun greet(name) {
  return "Hello, " + name + "!"
}

System.print(greet("World"))
System.print(greet("Wren"))

// Lists
System.print("\n=== Lists ===")
var fruits = ["apple", "banana", "cherry"]
System.print("Fruits:")
for (fruit in fruits) {
  System.print("  " + fruit)
}

// Classes
System.print("\n=== Classes ===")
class Counter {
  construct new() {
    _count = 0
  }
  
  increment() {
    _count = _count + 1
  }
  
  get count { _count }
}

var counter = Counter.new()
counter.increment()
counter.increment()
counter.increment()
System.print("Counter value: " + counter.count.toString)

System.print("\n=== Test Complete ===")
System.print("All examples executed successfully!")
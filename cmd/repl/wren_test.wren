System.print("=== Wren REPL Test ===")

// Variables
var x = 42
var name = "World"
System.print("Variable x: " + x.toString)
System.print("Variable name: " + name)

// Math
var a = 10
var b = 5
System.print("Math: " + a.toString + " + " + b.toString + " = " + (a + b).toString)

// Conditionals  
if (x > 40) {
  System.print("x is greater than 40")
}

// Loops
System.print("Counting 1 to 3:")
for (i in 1..3) {
  System.print("  " + i.toString)
}

// Lists
var fruits = ["apple", "banana", "cherry"]
System.print("Fruits:")
for (fruit in fruits) {
  System.print("  " + fruit)
}

// Simple class
class Greeter {
  construct new(greeting) {
    _greeting = greeting
  }
  
  greet(name) {
    System.print(_greeting + ", " + name + "!")
  }
}

var greeter = Greeter.new("Hello")
greeter.greet("Wren")

System.print("=== Test Complete ===")
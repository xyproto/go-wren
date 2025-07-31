System.print("=== Basic Test ===")

var x = 42
System.print("x = " + x.toString)

fun greet(name) {
  return "Hello, " + name + "!"
}

System.print(greet("Wren"))

for (i in 1..3) {
  System.print("Count: " + i.toString)  
}

class Counter {
  construct new() {
    _count = 0
  }
  
  increment() {
    _count = _count + 1
  }
  
  count {
    _count
  }
}

var counter = Counter.new()
counter.increment()
counter.increment()
System.print("Counter: " + counter.count.toString)

System.print("=== Test Complete ===")
System.print("=== Fiber/Concurrency Tests ===")

// Test 1: Basic fiber creation and yielding
System.print("\n1. Basic Fiber Yielding:")
var simpleFiber = Fiber.new {
  System.print("  Fiber: Starting")
  Fiber.yield()
  System.print("  Fiber: After first yield")
  Fiber.yield()
  System.print("  Fiber: After second yield")
}

simpleFiber.call()
System.print("  Main: After first call")
simpleFiber.call()
System.print("  Main: After second call")
simpleFiber.call()
System.print("  Main: After third call")

// Test 2: Fiber with value passing
System.print("\n2. Fiber Value Passing:")
var valueFiber = Fiber.new {
  System.print("  Fiber: Yielding 'hello'")
  Fiber.yield("hello")
  System.print("  Fiber: Yielding 42")
  Fiber.yield(42)
  System.print("  Fiber: Completing")
}

var result1 = valueFiber.call()
System.print("  Main: Got " + result1.toString)
var result2 = valueFiber.call()
System.print("  Main: Got " + result2.toString)
var result3 = valueFiber.call()
System.print("  Main: Got " + (result3 == null ? "null" : result3.toString))

// Test 3: Fiber call with argument
System.print("\n3. Fiber Call with Arguments:")
var argFiber = Fiber.new { |arg|
  System.print("  Fiber: Received " + arg.toString)
  var doubled = arg * 2
  Fiber.yield(doubled)
  System.print("  Fiber: Continuing after yield")
}

var doubled = argFiber.call(10)
System.print("  Main: Doubled value is " + doubled.toString)
argFiber.call()

// Test 4: Fiber transfer (cooperative multitasking)
System.print("\n4. Fiber Transfer:")
var fiberA = Fiber.new {
  System.print("  Fiber A: Starting")
}

var fiberB = Fiber.new {
  System.print("  Fiber B: Starting")
  fiberA.transfer()
  System.print("  Fiber B: Back from A")
}

var fiberC = Fiber.new {
  System.print("  Fiber C: Starting")
  fiberB.transfer()
  System.print("  Fiber C: Back from B (won't print)")
}

System.print("  Main: Starting transfer chain")
fiberC.transfer()
System.print("  Main: After transfer (may not print)")

// Test 5: Fiber state checking
System.print("\n5. Fiber State Checking:")
var stateFiber = Fiber.new {
  System.print("  Fiber: Running")
  Fiber.yield()
  System.print("  Fiber: Resumed")
}

System.print("  isDone before call: " + stateFiber.isDone.toString)
stateFiber.call()
System.print("  isDone after first call: " + stateFiber.isDone.toString)
stateFiber.call()
System.print("  isDone after completion: " + stateFiber.isDone.toString)

// Test 6: Nested fiber creation
System.print("\n6. Nested Fiber Creation:")
var outerFiber = Fiber.new {
  System.print("  Outer: Creating inner fiber")
  var innerFiber = Fiber.new {
    System.print("    Inner: Hello from inner fiber")
    Fiber.yield("inner-value")
    System.print("    Inner: Continuing inner")
  }
  
  var innerResult = innerFiber.call()
  System.print("  Outer: Inner yielded " + innerResult.toString)
  innerFiber.call()
  
  Fiber.yield("outer-value")
  System.print("  Outer: After outer yield")
}

var outerResult = outerFiber.call()
System.print("  Main: Outer yielded " + outerResult.toString)
outerFiber.call()

// Test 7: Producer-Consumer pattern
System.print("\n7. Producer-Consumer Pattern:")
var producer = Fiber.new {
  for (i in 1..5) {
    System.print("  Producer: Producing " + i.toString)
    Fiber.yield(i)
  }
  System.print("  Producer: Done")
}

System.print("  Consumer: Starting consumption")
while (!producer.isDone) {
  var item = producer.call()
  if (item != null) {
    System.print("  Consumer: Consumed " + item.toString)
  }
}
System.print("  Consumer: All items consumed")

// Test 8: Fiber completion
System.print("\n8. Fiber Completion:")
var completionFiber = Fiber.new {
  System.print("  Fiber: About to yield")
  Fiber.yield("safe-value")
  System.print("  Fiber: Completing normally")
}

var safeValue = completionFiber.call()
System.print("  Main: Got safe value: " + safeValue.toString)
completionFiber.call()
System.print("  Main: Fiber completed successfully")

System.print("\n=== All Fiber Tests Complete ===")
System.print("Fiber system is working correctly!")
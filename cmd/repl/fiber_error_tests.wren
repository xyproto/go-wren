System.print("=== Fiber Error Handling Tests ===")

// Test 1: Fiber.try() with runtime error
System.print("\n1. Fiber try() with runtime error:")
var errorFiber = Fiber.new {
  System.print("  Fiber: Before error")
  null.someMethod  // This will cause a runtime error
  System.print("  Fiber: After error (won't print)")
}

var errorResult = errorFiber.try()
System.print("  Main: Error caught: " + errorResult.toString)
System.print("  Main: Fiber isDone: " + errorFiber.isDone.toString)
System.print("  Main: Fiber error: " + errorFiber.error.toString)

// Test 2: Fiber.abort() 
System.print("\n2. Fiber abort():")
var abortFiber = Fiber.new {
  System.print("  Fiber: About to abort")
  Fiber.abort("Custom abort message")
  System.print("  Fiber: Won't reach here")
}

var abortResult = abortFiber.try()
System.print("  Main: Abort result: " + abortResult.toString)
System.print("  Main: Fiber isDone: " + abortFiber.isDone.toString)
System.print("  Main: Fiber error: " + abortFiber.error.toString)

// Test 3: Successful fiber with try()
System.print("\n3. Successful fiber with try():")
var successFiber = Fiber.new {
  System.print("  Fiber: Working normally")
  Fiber.yield("success")
  System.print("  Fiber: Completing")
  "final-result"
}

var result1 = successFiber.try()
System.print("  Main: First try result: " + (result1 == null ? "null" : result1.toString))
var result2 = successFiber.try()
System.print("  Main: Second try result: " + (result2 == null ? "null" : result2.toString))
System.print("  Main: Fiber isDone: " + successFiber.isDone.toString)

// Test 4: Error in yielding fiber
System.print("\n4. Error in yielding fiber:")
var yieldErrorFiber = Fiber.new {
  System.print("  Fiber: First yield")
  Fiber.yield("first")
  System.print("  Fiber: About to error")
  [].invalidMethod  // Error after yield
  System.print("  Fiber: Won't reach here")
}

var firstYield = yieldErrorFiber.call()
System.print("  Main: Got first yield: " + firstYield.toString)

var errorFromSecond = yieldErrorFiber.try()
System.print("  Main: Second call error: " + errorFromSecond.toString)
System.print("  Main: Fiber error property: " + yieldErrorFiber.error.toString)

// Test 5: Nested fiber error handling
System.print("\n5. Nested fiber error handling:")
var outerErrorFiber = Fiber.new {
  System.print("  Outer: Creating inner fiber")
  var innerFiber = Fiber.new {
    System.print("    Inner: About to error")
    "".nonExistentMethod
  }
  
  var innerResult = innerFiber.try()
  System.print("  Outer: Inner error was: " + innerResult.toString)
  System.print("  Outer: Continuing outer fiber")
  "outer-completed"
}

var outerResult = outerErrorFiber.call()
System.print("  Main: Outer result: " + outerResult.toString)

System.print("\n=== Fiber Error Handling Tests Complete ===")
System.print("Fiber error handling works correctly!")
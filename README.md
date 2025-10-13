# Go Structured Panic Handler: Try-Catch-Finally

This project implements a **structured exception handling** pattern for Go that mimics the familiar `try-catch-finally` construct found in languages like Java, C++, and Python. It leverages Go's built-in `defer` and `recover` functions to wrap protected code blocks, ensuring clean recovery from panics and guaranteeing execution of cleanup logic.

## Why Use This?

Go's standard panic handling is done with a combination of `defer` and `recover()`, which can become verbose and repetitive, especially when a panic needs to be caught, logged, and then followed by cleanup code. This implementation provides:

1.  **Readability:** Encapsulates the logic into a clear `Try`, `Catch`, and `Finally` structure.
2.  **Detailed Context:** Automatically captures the **stack trace** and packages the recovered value into a comprehensive `Panic` struct for easy logging and inspection.
3.  **Guaranteed Cleanup:** The `Finally` block is guaranteed to execute, regardless of whether the `Try` block finished successfully or panicked.

## Code Structure

### 1. The `Panic` Struct

This struct wraps the details of a recovered panic, making the data easy to work with in the `Catch` block.

| Field | Type | Description |
| :--- | :--- | :--- |
| `Recover` | `any` | The raw value passed to the `panic()` call. |
| `Message` | `string` | A formatted string describing the panic. |
| `Stack` | `string` | The full **stack trace** captured via `runtime/debug.Stack()`. |

It also implements the built-in `error` interface.

### 2. The `PanicHandler` Struct

This is the central control mechanism, defining the blocks of code to be executed.

| Field | Type | Description |
| :--- | :--- | :--- |
| `Try` | `func()` | **(Required)** The function containing the code that may panic. |
| `Catch` | `func(Panic)` | **(Optional)** Executed if a panic occurs in `Try`. Receives the detailed `Panic` struct. |
| `Finally`| `func()` | **(Optional)** **Always** executed after `Try`, regardless of success or panic. |

### 3. The `Handle()` Method

This method sets up the deferred functions in the correct order:

1.  **`defer p.Finally()`**: Ensures the cleanup code runs first during unwinding.
2.  **`defer func(){ if r := recover() ... }()`**: Sets the `recover()` trap. If a panic is active, it runs the `Catch` logic and stops the panic from continuing.
3.  **`p.Try()`**: Executes the main protected code.

## Usage Example

The following example demonstrates a panic being thrown, caught, and followed by the guaranteed `Finally` block.

```go
func main() {
  PanicHandler{
    Try: func(){
      log.Println("Try block executed.")
      // A deliberate panic is thrown here
      panic(fmt.Errorf("Thrown Error: Database Connection Failed")) 
    },
    Catch: func(p Panic){
      log.Printf("--- Catch Block ---\n")
      log.Printf("Message: %s\n", p.Message)
      log.Printf("Recovered Type: %T\n", p.Recover)
      // log.Printf("Stack Trace:\n%s", p.Stack) // Uncomment to see full trace
    },
    Finally: func(){
      log.Println("--- Finally Block ---\nCleanup logic runs regardless of panic.")
    },
  }.Handle()
  
  log.Println("Program continued execution normally after recovery.")
}
```

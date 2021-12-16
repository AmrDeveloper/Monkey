# Monkey
Monkey programming language 🐒 project from "Writing An Interpreter In Go" and "Writing A Compiler In Go" Books

![Fabonacci](/media/fabonacci.png)

![CountDown](/media/countdown.png)

### Language Features:
- Integers, booleans, strings, arrays, hash maps
- A REPL
- Arithmetic expressions
- Let statements
- First-class and higher-order functions
- Built-in functions
- Recursion
- Closures

### Run REPL using Compiler and Virtual Machine

```
go run . -engine=vm
```

### Run REPL using Tree-walking interpreter

```
go run . -engine=eval
```

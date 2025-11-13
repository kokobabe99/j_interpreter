# J Language Interpreter

A Interpreter project (tokenizer,parser,Interpreter) for the custom **J Language**.

---

## ✅ Features

- Recognizes **all language keywords** (`pkg`, `imp`, `def`, `var`, `if`, `fr`, `select`, `panic`, `recover`, etc.)
- Supports:
    - Identifiers (Unicode allowed)
    - Integer & float literals with underscore rules
    - String literals (`"..."`) and raw strings (`` `...` ``)
    - Character literals (`'a'`, `'\n'`, `'\x41'`)
    - Type indicators (`i32`, `f64`, `bool`, `string`, ...)
    - Operators & delimiters (`==`, `<=`, `:=`, `<-`, etc.)
    - Line comments (`//`)
    - **Nested block comments** (`/* ... /* ... */ ... */`)
- **Lexical error detection** with **line and column number**
- Outputs **JSON** containing all tokens and errors
- Writes result to:
    - **stdout**, and
    - an output file based on source file name
        - e.g. `main.jl → main_jl_output.txt`

---

## ✅ How to Run

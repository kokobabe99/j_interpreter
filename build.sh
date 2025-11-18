#!/bin/bash
set -e

echo "Generating ANTLR parser..."
/opt/homebrew/opt/openjdk/bin/java -jar /opt/homebrew/opt/antlr/antlr-4.13.2-complete.jar \
  -Dlanguage=Go -visitor -package parser_src -o parser_src JLang.g4

echo "Waiting for file system to settle..."
sleep 5

echo "Tidying modules..."
go mod tidy

echo "Building..."
go build -o bin/parser cmd/parser/main.go
go build -o bin/tokenizer cmd/tokenizer/main.go
go build -o bin/interpreter cmd/interpreter/main.go

echo "Running samples #1................................................"
go run cmd/parser/main.go samples/sample1.jl

echo "Running samples #2................................................"
go run cmd/parser/main.go samples/sample2.jl

echo "Running samples #3................................................"
go run cmd/parser/main.go samples/sample3.jl

echo "Running samples #4................................................"
go run cmd/parser/main.go samples/sample4.jl

echo "Running interpreter tests #1........................................."
go run cmd/interpreter/main.go samples/sample1.jl || true

echo "Running interpreter tests #2........................................."
go run cmd/interpreter/main.go samples/sample2.jl || true

echo "Running interpreter tests #3........................................."
go run cmd/interpreter/main.go samples/sample3.jl || true

echo "Running interpreter tests #4........................................."
go run cmd/interpreter/main.go samples/sample4.jl || true
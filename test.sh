#!/bin/bash
set -e

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


echo "Running interpreter tests #1........................................."
go run cmd/interpreter/main.go samples/sample1.jl || true

echo "Running interpreter tests #2........................................."
go run cmd/interpreter/main.go samples/sample2.jl || true

echo "Running interpreter tests #3........................................."
go run cmd/interpreter/main.go samples/sample3.jl || true

echo "Running interpreter tests #if........................................."
go run cmd/interpreter/main.go samples/sample_if.jl || true

echo "Running interpreter tests #for........................................."
go run cmd/interpreter/main.go samples/sample_for.jl || true

# echo "Running interpreter tests #joto........................................."
# go run cmd/interpreter/main.go samples/sample_joto.jl || true
echo "Running interpreter tests #case........................................."
go run cmd/interpreter/main.go samples/sample_case_break.jl || true
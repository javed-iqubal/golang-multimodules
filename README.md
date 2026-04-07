# Golang Multimodules

## In root project "golang-multimodules"
- golang-multimodules$ go work init ./hello ./greeting

## In hello module
- golang-multimodules/hello$ go mod edit -replace github.com/demo/greetings=../greetings
- golang-multimodules/hello$ go mod tidy
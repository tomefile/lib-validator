# Tomefile Validator library

Library to validate Tomefile Parser output

<!-- vim-markdown-toc GFM -->

* [Example Usage](#example-usage)
    * [Using it directly](#using-it-directly)
    * [Using parser's post-processing (recommended)](#using-parsers-post-processing-recommended)

<!-- vim-markdown-toc -->

## Example Usage

### Using it directly

```go
node, err := libvalidator.Validate(...)
if err != nil {
    err.BeautyPrint(os.Stderr)
    os.Exit(1)
}
```

### Using parser's post-processing (recommended)

```go
// This has the benefit of having fields expanded:
// - libparser.ExecNode{}.Binary -> changes to full file path to binary
// - libparser.StringNode{}.Segments -> stores format segments

parser := libparser.New(
    "example.tome",
    bufio.NewReader(file),
    libparser.Validate,
)
tree, err := parser.Parse()
if err != nil {
    err.BeautyPrint(os.Stderr)
    os.Exit(1)
}
```

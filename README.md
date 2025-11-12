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
node, derr := libvalidator.Validate(...)
if derr != nil {
    derr.Print(os.Stderr)
    os.Exit(1)
}
```

### Using parser's post-processing (recommended)

```go
// This has the benefit of having fields expanded:
// - libparser.ExecNode{}.Binary -> changes to full file path to binary
// - libparser.StringNode{}.Segments -> stores format segments

parser := libparser.New(file).With(libparser.Validate)

tree, derr := parser.Parse()
if derr != nil {
    derr.Print(os.Stderr)
    os.Exit(1)
}
```

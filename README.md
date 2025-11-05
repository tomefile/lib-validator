# Tomefile Validator library

Library to validate Tomefile Parser output

## Example Usage

```go
node, err := libvalidator.Validate(&libparser.Node{...})
if err != nil {
    log.Fatalf("%#v: %s", node, err.Error())
}
```

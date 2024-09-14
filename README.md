# Keycheck
 Keycheck is a Go library for validating data against conditions.

## Installation

`go get github.com/brianxor/keycheck`

## Supported Modes

- Or
- And

## Default Types

> [!NOTE]
> You can define your own custom type.

- None
- Failure
- Success
- Custom
- Retry
- Error

## Supported Conditions

- Contains
- Equal

## Usage

```go

var (
	source = "success: true"
	statusCode = 200
)

orModeKeychain := keycheck.NewKeychain().
	SetMode(keycheck.OrMode).
	SetType(keycheck.Success).
	SetKey(source, "true", keycheck.ContainsCondition).
	SetKey(statusCode, 200, keycheck.EqualCondition)

andModeKeychain := keycheck.NewKeychain().
	SetMode(keycheck.AndMode).
	SetType(keycheck.Failure).
	SetKey(source, "false", keycheck.ContainsCondition).
	SetKey(statusCode, 403, keycheck.EqualCondition)
	

newKeycheck := keycheck.NewKeycheck().
	AddKeychains(orModeKeychain, andModeKeychain)

keycheckResult, ok := newKeycheck.Validate()

if !ok {
	fmt.Println("Keycheck validation failed.")
	return
}

fmt.Println(keycheckResult)
```
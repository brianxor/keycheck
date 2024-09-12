# Keycheck
 Keycheck is a Go library for validating data against conditions.

## Installation

`go get github.com/brianxor/keycheck`

## Supported Conditions

- Contains
- Equal

## Usage

Usage is straightforward:

### Contains
```go
input := "success: true"

// Build a keychain
//
// If input contains "true", then the keychain is valid.
newKeychain := keycheck.NewKeychain().
	SetType("SUCCESS").
	SetKey(input, "true", keycheck.ContainsCondition)

// Build a keycheck
newKeycheck := keycheck.NewKeycheck().
	AddKeychains(newKeychain)

// Validate keycheck
keycheckResult, ok := newKeycheck.Validate()

if !ok {
	fmt.Println("Keycheck validation failed.")
	return
}

fmt.Println(keycheckResult)
```

### Equal
```go
input := "success"

// Build a keychain
//
// If input is equal to "success", then the keychain is valid.
newKeychain := keycheck.NewKeychain().
	SetType("SUCCESS").
	SetKey(input, "success", keycheck.EqualCondition)

// Build a keycheck
newKeycheck := keycheck.NewKeycheck().
	AddKeychains(newKeychain)

// Validate keycheck
keycheckResult, ok := newKeycheck.Validate()

if !ok {
	fmt.Println("Keycheck validation failed.")
	return
}

fmt.Println(keycheckResult)
```

You can also combine the modes:

```go
input := "success: true"
statusCode := 204

// Build a keychain
newKeychain := keycheck.NewKeychain().
	SetType("SUCCESS").
	SetKey(input, "true", keycheck.ContainsCondition).
	SetKey(statusCode, 204, keycheck.EqualCondition) 

// Build a keycheck
newKeycheck := keycheck.NewKeycheck().
	AddKeychains(newKeychain)

// Validate keycheck
keycheckResult, ok := newKeycheck.Validate()

if !ok {
	fmt.Println("Keycheck validation failed.")
	return
}

fmt.Println(keycheckResult)
```

You can also have multiple keychains:

```go
input := "success: true"
statusCode := 204

// Build keychains

successKeychain := keycheck.NewKeychain().
	SetType("SUCCESS").
	SetKey(input, "true", keycheck.ContainsCondition).
	SetKey(statusCode, 204, keycheck.EqualCondition) 

errorKeychain := keycheck.NewKeychain().
	SetType("ERROR").
	SetKey(input, "false", keycheck.ContainsCondition).
	SetKey(statusCode, 403, keycheck.EqualCondition)

// Build a keycheck
newKeycheck := keycheck.NewKeycheck().
	AddKeychains(newKeychain, errorKeychain)

// Validate keycheck
keycheckResult, ok := newKeycheck.Validate()

if !ok {
	fmt.Println("Keycheck validation failed.")
	return
}

fmt.Println(keycheckResult)
```
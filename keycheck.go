package keycheck

import (
	"strings"
)

type Keychain struct {
	Type string
	Keys []*Key
}

type Condition int

const (
	ContainsCondition Condition = iota
	EqualCondition
)

type Key struct {
	Input     interface{}
	Expected  interface{}
	Condition Condition
}

func NewKeychain() *Keychain {
	return &Keychain{}
}

func (keychain *Keychain) SetType(keychainType string) *Keychain {
	keychain.Type = keychainType
	return keychain
}

func (keychain *Keychain) SetKey(input interface{}, expected interface{}, condition Condition) *Keychain {
	keychain.Keys = append(keychain.Keys, &Key{
		Input:     input,
		Expected:  expected,
		Condition: condition,
	})

	return keychain
}

type Keycheck struct {
	Keychains []*Keychain
}

func NewKeycheck() *Keycheck {
	return &Keycheck{
		Keychains: []*Keychain{},
	}
}

func (keycheck *Keycheck) AddKeychains(keychains ...*Keychain) *Keycheck {
	keycheck.Keychains = append(keycheck.Keychains, keychains...)
	return keycheck
}

func (keycheck *Keycheck) Validate() (string, bool) {
	for _, keychain := range keycheck.Keychains {
		for _, key := range keychain.Keys {
			if checkCondition(key) {
				return keychain.Type, true
			}
		}
	}

	return "", false
}

func checkCondition(key *Key) bool {
	switch key.Condition {
	case EqualCondition:
		return checkEqualCondition(key)
	case ContainsCondition:
		return checkContainsCondition(key)
	default:
		return false
	}
}

func checkEqualCondition(key *Key) bool {
	switch input := key.Input.(type) {
	case string:
		if expected, ok := key.Expected.(string); ok {
			return input == expected
		}
	case int:
		if expected, ok := key.Expected.(int); ok {
			return input == expected
		}
	}
	return false
}

func checkContainsCondition(key *Key) bool {
	if inputStr, ok := key.Input.(string); ok {
		if expectedStr, ok := key.Expected.(string); ok {
			return strings.Contains(inputStr, expectedStr)
		}
	}
	return false
}

package helpers

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

const characterAlfabetUpper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const characterAlfabetLower string = "abcdefghijklmnopqrstuvwxyz"
const characterNumber string = "0123456789"

var (
	Generator = struct {
		Generate                 func(length int, newChar ...string) string
		GenerateCustom           func(length int, char string) string
		GenerateNumber           func(length int) (int, error)
		GenerateNumberString     func(length int) string
		GenerateAlphabet         func(length int, newChar ...string) string
		GenerateAlphabetLower    func(length int, newChar ...string) string
		GenerateAlphabetUpper    func(length int, newChar ...string) string
		GenerateAlphabetNumUpper func(length int, newChar ...string) string
	}{
		Generate: func(length int, newChar ...string) string {
			characters := characterAlfabetUpper + characterAlfabetLower + characterNumber
			result := ""
			if len(newChar) > 0 {
				result += newChar[0]
			}
			charactersLength := len(characters)
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(charactersLength)
				result += string(characters[randomIndex])
			}
			return result
		},
		GenerateCustom: func(length int, char string) string {
			characters := char
			result := ""
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(len(characters))
				result += string(characters[randomIndex])
			}
			return result
		},

		GenerateNumber: func(length int) (int, error) {
			if length >= 20 {
				return 0, errors.New("length should not be greater than 20")
			}

			rand.New(rand.NewSource(time.Now().UnixNano()))
			min := int64(math.Pow10(length - 1))
			max := int64(math.Pow10(length) - 1)
			randomNumber := rand.Intn(int(max-min+1)) + int(min)
			return randomNumber, nil
		},

		GenerateAlphabet: func(length int, newChar ...string) string {
			characters := characterAlfabetUpper + characterAlfabetLower
			result := ""
			if len(newChar) > 0 {
				result += newChar[0]
			}
			charactersLength := len(characters)
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(charactersLength)
				result += string(characters[randomIndex])
			}
			return result
		},

		GenerateNumberString: func(length int) string {
			characters := characterNumber
			result := ""
			charactersLength := len(characters)
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(charactersLength)
				result += string(characters[randomIndex])
			}
			return result
		},

		GenerateAlphabetLower: func(length int, newChar ...string) string {
			characters := characterAlfabetLower
			result := ""
			if len(newChar) > 0 {
				result += newChar[0]
			}
			charactersLength := len(characters)
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(charactersLength)
				result += string(characters[randomIndex])
			}
			return result
		},

		GenerateAlphabetUpper: func(length int, newChar ...string) string {
			characters := characterAlfabetUpper
			result := ""
			if len(newChar) > 0 {
				result += newChar[0]
			}
			charactersLength := len(characters)
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(charactersLength)
				result += string(characters[randomIndex])
			}
			return result
		},

		GenerateAlphabetNumUpper: func(length int, newChar ...string) string {
			characters := characterAlfabetUpper + characterNumber
			result := ""
			if len(newChar) > 0 {
				result += newChar[0]
			}
			charactersLength := len(characters)
			rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i < length; i++ {
				randomIndex := rand.Intn(charactersLength)
				result += string(characters[randomIndex])
			}
			return result
		},
	}
)

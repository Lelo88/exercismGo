package dndcharacter

import (
	"crypto/rand"
	"math"
	"math/big"
)

type Character struct {
    Strength     int
    Dexterity    int
    Constitution int
    Intelligence int
    Wisdom       int
    Charisma     int
    Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    for {
        n, _ := rand.Int(rand.Reader, big.NewInt(16))
        ability := 3 + int(n.Int64())
        if ability >= 3 && ability <= 18 {
            return ability
        }
    }
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    constitution := Ability()
    return Character{
        Strength:     Ability(),
        Dexterity:    Ability(),
        Constitution: constitution,
        Intelligence: Ability(),
        Wisdom:       Ability(),
        Charisma:     Ability(),
        Hitpoints:    10 + Modifier(constitution),
    }
}
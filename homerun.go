package main

import "fmt"

type Batter struct {
	Height  int
	Weight  int
	Fitness int
	Age     int
	Intel   int
	Contact int
	Eyes    int
	Power   int
}

func (b *Batter) CalcAttrib() {
	weightRatio := int((float32(b.Weight) / float32(b.Height)) * 10)
	contactRatio := 50 - weightRatio
	fitnessPower := b.Fitness * 10
	fitnessContact := 30 - fitnessPower
	ageBonus := 40 - b.Age
	intelBonus := b.Intel * 10
	b.Power = weightRatio + fitnessPower + ageBonus
	b.Contact = fitnessContact + ageBonus + contactRatio
	eyeBonus := 130 - b.Power - b.Contact
	b.Eyes = intelBonus + ageBonus + eyeBonus
	switch {
	case b.Power > 99:
		b.Power = 99
	case b.Contact > 99:
		b.Contact = 99
	case b.Eyes > 99:
		b.Eyes = 99
	}
}

func main() {
	b := new(Batter)
	b.Height = 75
	b.Weight = 175
	b.Fitness = 1
	b.Intel = 3
	b.Age = 20
	b.CalcAttrib()
	fmt.Printf("Contact: %v / Power: %v / Eyes: %v", b.Contact, b.Power, b.Eyes)
}

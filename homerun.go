package main

import (
	"fmt"
	"strconv"
)

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

func (b *Batter) Define() {
	var h, w, f, a, i string
	println("How tall is your batter in inches? (60 - 80)\n#")
	fmt.Scan(&h)
	println("\n\nHow much does your batter weigh?\n#")
	fmt.Scan(&w)
	println("\n\nHow old is your batter? (20 to 40)\n#")
	fmt.Scan(&a)
	println("\n\nWhen your batter exercises, does he:\n1) Focus on cardio\n2) Balance between cardio and weights\n3) Mainly lift weights")
	fmt.Scan(&f)
	println("\n\nHow devoted to game preparation is your batter?:\n1) Would rather work out\n2) Pretty devoted\n3) Stays up all night watching tape")
	fmt.Scan(&i)
	var err error
	b.Height, err = strconv.Atoi(h)
	b.Weight, err = strconv.Atoi(w)
	b.Age, err = strconv.Atoi(a)
	b.Fitness, err = strconv.Atoi(f)
	b.Intel, err = strconv.Atoi(i)
	if err != nil {
		println("An error.")
	}
}

func main() {
	b := new(Batter)
	b.Define()
	b.CalcAttrib()
	fmt.Printf("Contact: %v / Power: %v / Eyes: %v\n\n", b.Contact, b.Power, b.Eyes)
}

package main

import (
	"fmt"
	"github.com/troyharris/newrand"
	"strconv"
)

var homeruns int = 0
var pitches int = 0
var outs int = 0
var powSwing int = 3

type FieldLoc struct {
	Depth int
	Side  int
}

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

type Ball struct {
	Speed    int
	Acc      int
	Movement int
	HitLoc   FieldLoc
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

func (b *Ball) Pitch() {
	b.Speed = newrand.Intr(70, 99)
	b.Acc = newrand.Intr(55, 100)
	b.Movement = newrand.Intr(40, 100) - (b.Speed / 2)
}

func (b Batter) Swing(ball *Ball) (string, *FieldLoc) {
	var result string
	location := new(FieldLoc)
	checkSwingPer := b.Eyes - ball.Acc + 70
	if checkSwingPer > 99 {
		checkSwingPer = 99
	}
	if ball.Acc < 70 && newrand.Hit(checkSwingPer) {
		result = "ball"
		return result, location
	}
	contactPer := (b.Contact + b.Eyes + newrand.Intr(80, 100)) - ((100 - ball.Acc) + ball.Movement + ball.Speed)
	if contactPer > 99 {
		contactPer = 99
	}
	if newrand.Hit(contactPer) {
		result = "hit"
		var d int
		powerLvl := ((float32(b.Power)) + (float32(contactPer)) + float32(newrand.Intr(1, 50))) / 2
		switch {
		case powerLvl < 35:
			d = 0
		case powerLvl < 65:
			d = 1
		case powerLvl < 80:
			d = 2
		default:
			d = 3
		}
		dirStart := 0
		dirEnd := 4
		if contactPer > 75 {
			dirStart = 1
		}
		if newrand.Hit(b.Eyes) {
			dirEnd = 3
		}
		direction := newrand.Intr(dirStart, dirEnd)
		location.Depth = d
		location.Side = direction
		return result, location
	} else {
		result = "strike"
		outs++
		return result, location
	}
	return result, location
}

func (b *Ball) Hit(l *FieldLoc) string {
	b.HitLoc.Depth = l.Depth
	b.HitLoc.Side = l.Side
	var dir string
	var loc string
	var result string
	switch b.HitLoc.Side {
	case 0:
		dir = "foul"
	case 1:
		dir = "left field"
	case 2:
		dir = "center field"
	case 3:
		dir = "right field"
	case 4:
		dir = "foul"
	}
	switch b.HitLoc.Depth {
	case 0:
		loc = "shallow"
	case 1:
		loc = "middle"
	case 2:
		loc = "deep"
	case 3:
		loc = "homerun"
	}
	if dir != "ball" && loc != "homerun" {
		outs++
	}
	if dir == "foul" {
		result = "Ball hit foul"
		return result
	}
	if loc == "homerun" {
		result = "It's a home run!"
		homeruns++
		return result
	}
	result = fmt.Sprintf("Ball hit to %s %s.", loc, dir)
	return result
}

func play(bt *Batter, bl *Ball) string {
	var msg string
	bl.Pitch()
	result, location := bt.Swing(bl)
	switch result {
	case "ball":
		msg = "Batter lays off that pitch. Ball"
	case "strike":
		msg = "A big swing and a miss"
	case "hit":
		msg = bl.Hit(location)
	}
	return msg
}

func atBat(bt *Batter, bl *Ball) string {
	var s string
	fmt.Printf("\nThe pitcher winds up. Here comes the pitch. You: 1) take a fluid swing 2) swing for the fences (%v left)", powSwing)
	fmt.Scan(&s)
	switch s {
	case "1":
		return play(bt, bl)
	case "2":
		return play(bt, bl)
	default:
		atBat(bt, bl)
	}
	return ""
}

func main() {
	b := new(Batter)
	b.Define()
	b.CalcAttrib()
	fmt.Printf("Contact: %v / Power: %v / Eyes: %v\n\n", b.Contact, b.Power, b.Eyes)
	ball := new(Ball)
	for outs = 0; outs < 10; pitches++ {
		fmt.Printf("\nThere are %v outs.\n", outs)
		fmt.Println("")
		fmt.Println(atBat(b, ball))
	}
	fmt.Printf("Home runs: %v", homeruns)
}

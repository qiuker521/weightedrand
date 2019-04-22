package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/go-ego/murmur"
	"github.com/qiuker521/weightedrand"
)

func main() {
	randpick()
	hashpick()
}
func randpick() {
	c := weightedrand.Chooser{}
	var choices = []weightedrand.Choice{}
	for i := 1; i <= 5; i++ {
		//every choice has a random weight
		choices = append(choices, weightedrand.Choice{strconv.Itoa(i), rand.Float64()})
	}
	c.NewChooser(choices...)
	//This is always a random pick
	log.Println(c.Pick())
}

func hashpick() {
	c := weightedrand.Chooser{}

	var choices = []weightedrand.Choice{}
	for i := 1; i <= 5; i++ {
		//every choice has the same weight to test hash pick
		choices = append(choices, weightedrand.Choice{strconv.Itoa(i), 1})
	}
	c.NewChooser(choices...)

	//This is a hash pick which will always return choice "4".
	log.Println(c.PickByHash(float64(murmur.Sum32("This is a string, such as a user id"))))
}

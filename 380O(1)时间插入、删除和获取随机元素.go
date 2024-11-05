package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	vals     []int
	valToIdx map[int]int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		vals:     []int{},
		valToIdx: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.valToIdx[val]; exists {
		return false
	}
	this.valToIdx[val] = len(this.vals)
	this.vals = append(this.vals, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.valToIdx[val]
	if !exists {
		return false
	}
	lastVal := this.vals[len(this.vals)-1]
	this.vals[idx] = lastVal
	this.valToIdx[lastVal] = idx
	this.vals = this.vals[:len(this.vals)-1]
	delete(this.valToIdx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

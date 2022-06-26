package main

import (
	"fmt"
	"strings"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	items []Checkable
}

func (c *Checker) Add(checkable Checkable) {
	c.items = append(c.items, checkable)
}

func (c *Checker) Check() {
	for _, v := range c.items {
		if !v.Health() {
			fmt.Println(v.GetID() + " not working")
		}
	}
}

func (c Checker) String() string {
	var ids []string
	for _, v := range c.items {
		ids = append(ids, v.GetID())
	}
	return strings.Join(ids, ", ")
}

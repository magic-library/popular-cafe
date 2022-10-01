package main

import "time"

type Table struct {
	coffeeCount map[int]int
	capacity    int
}

func NewTable(capacity int) *Table {
	return &Table{
		coffeeCount: map[int]int{},
		capacity:    capacity,
	}
}

func (table *Table) 커피를_올리다() {

}

func (table *Table) 커피를_차감하다() {

}

type Coffee struct {
	Id                int
	Name              string
	ManufacturingTime time.Duration
	WarmTime          time.Duration
}

var coffeeList = []Coffee{
	{
		Id:                1,
		Name:              "Americano",
		ManufacturingTime: time.Second * 10,
		WarmTime:          time.Second * 3,
	},
	{
		Id:                2,
		Name:              "Cappuccino",
		ManufacturingTime: time.Second * 20,
		WarmTime:          time.Second * 5,
	},
	{
		Id:                3,
		Name:              "Macchiato",
		ManufacturingTime: time.Second * 25,
		WarmTime:          time.Second * 8,
	},
}

package main

import "fmt"

type Cart struct {
	Items map[string]*int
}

func NewCart() *Cart {
	return &Cart{
		make(map[string]*int),
	}
}

func (c *Cart) tambahProduk(kodeProduk string, kuantitas int) {
	var countVal int = 0
	data := c.Items[kodeProduk]
	if data != nil {
		countVal = *data
	}
	countVal += kuantitas
	c.Items[kodeProduk] = &countVal
}

func (c *Cart) hapusProduk(kodeProduk string) {
	delete(c.Items, kodeProduk)
}

func (c *Cart) tampilkanCart() {
	for key, val := range c.Items {
		fmt.Printf("%v (%v)\n", key, *val)
	}
}

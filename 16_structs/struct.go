package main

import (
	"fmt"
	"time"
)

// order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
// func (o *order) changeStatus(status string) {
//	o.status = status
//}

func main() {

	myOrder := newOrder("1", 30.50, "received")
	fmt.Println(myOrder)
	/* myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	} */

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)

	/* myOrder.createdAt = time.Now()

	fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder.status = "paid"

	fmt.Println("Order struct", myOrder2)
	fmt.Println("Order struct", myOrder) */
}

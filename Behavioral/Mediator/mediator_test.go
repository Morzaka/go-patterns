package Mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {

	farmer := new(Farmer)
	cannery := new(Cannery)
	shop := new(Shop)

	farmer.AddMoney(7500.00)
	cannery.AddMoney(15000.00)
	shop.AddMoney(30000.00)

	ConnectColleagues(farmer, cannery, shop)

	// Фермер виротив помідорів на 1000 кг ваги і інформує посередника про завершення своєї роботи.
	// Далі посередник відправляє помідори на консервний завод.
	// Після того, як консервний завод виробляє 1000 упаковок кетчупу,
	// він повідомляє посереднику про доставку в магазин.

	farmer.GrowTomato(1000)

	expect := float64(54750)
	result := shop.GetMoney()

	if result != expect {
		t.Errorf("Expect result to equal %f, but %f.\n", expect, result)
	}
}

package main

import "fmt"

type driver struct {
}

func (c *driver) installFivePinWheelOnCar(automobile car) {
	fmt.Println("driver install five pin wheel on car.")
	automobile.installFivePinWheel()
}

type car interface {
	installFivePinWheel()
}
type toyota struct {
}

func (t *toyota) installFivePinWheel() {
	fmt.Println("Five pin wheel is installed into toyota.")
}

type lada struct{}

func (l *lada) installFourPinWheel() {
	fmt.Println("Four pin wheel is installed into lada.")
}

type ladaAdapter struct {
	ladaCar *lada
}

func (automobile *ladaAdapter) installFivePinWheel() {
	fmt.Println("Adapter converts five pin to four pin.")
	automobile.ladaCar.installFourPinWheel() //выполняем функцию доступную для структуры лада
}
func main() {
	driver := &driver{}
	toyota := &toyota{}

	driver.installFivePinWheelOnCar(toyota)

	ladatoyota := &lada{}
	ladatoyotaAdapter := &ladaAdapter{
		ladaCar: ladatoyota,
	}

	driver.installFivePinWheelOnCar(ladatoyotaAdapter)
}

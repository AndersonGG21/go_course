package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarV        = "CAR"
	MotorCycleV = "MOTORCYCLE"
	TruckV      = "TRUCK"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarV:
		return &Car{Time: time}, nil
	case MotorCycleV:
		return &MotorCycle{Time: time}, nil
	case TruckV:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("vehicle %s not exists", v)
	}
}

type Car struct {
	Time int
}

type MotorCycle struct {
	Time int
}

type Truck struct {
	Time int
}

func (c Car) Distance() float64 {
	return 100 * float64(c.Time) / 60
}

func (m MotorCycle) Distance() float64 {
	return 120 * float64(m.Time) / 60
}

func (t Truck) Distance() float64 {
	return 70 * float64(t.Time) / 60
}

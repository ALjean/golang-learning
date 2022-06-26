package main

import "fmt"

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceID int
	Status    string // TODO Enums ?
}

func GenerateCheck() []HealthCheck {
	healthCheckList := make([]HealthCheck, 6)

	for i := range healthCheckList {
		var status string
		if i%2 == 0 {
			status = PassStatus
		} else {
			status = FailStatus
		}

		healthCheckList[i] = HealthCheck{i, status}
	}

	return healthCheckList
}

func main() {
	healthCheckList := GenerateCheck()
	for _, v := range healthCheckList {
		if v.Status == PassStatus {
			fmt.Println("There will be id:")
			fmt.Println(v.ServiceID)
		}
	}
}

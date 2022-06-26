package main

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

// TODO complete lesson
func main() {

}

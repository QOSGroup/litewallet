package health

type HealthCheck interface {
	HealthCheck() error
	IsHardDependency() bool
	Description() string
}

func NewHealthCheck(check func() error, isHardDependency bool, description string) HealthCheck {
	return &healthChecker{
		check:            check,
		isHardDependency: isHardDependency,
		description:      description,
	}
}

type healthChecker struct {
	check            func() error
	isHardDependency bool
	description      string
}

func (hc healthChecker) HealthCheck() error {
	return hc.check()
}
func (hc healthChecker) IsHardDependency() bool {
	return hc.isHardDependency
}
func (hc healthChecker) Description() string {
	return hc.description
}

package cmd

type connectChecker struct {
}

func (dc *connectChecker) Desc() string {
	return "Connect checker"
}

func (dc *connectChecker) Check() error {
	return nil
}

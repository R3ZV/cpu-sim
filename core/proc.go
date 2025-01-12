package core

type Proc struct {
	Pid       int
	Burst     int
	InitBurst int
	Arrive    int
	Priority  int

	// -1 for no period
	Period int
}

func NewProc(pid, burst, arrive, priority, period int) *Proc {
	return &Proc{
		pid,
		burst,
		burst,
		arrive,
		priority,
		period,
	}
}

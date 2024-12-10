package core

type Proc struct {
	Pid            int
	Burst          int
	RemainingBurst int

	// time until the process should execute again
	Period   int
	Priority int
}

func NewProc(pid int, burst int, period int) *Proc {
	p := Proc{
		Pid:            pid,
		Burst:          burst,
		RemainingBurst: burst,
		Period:         period,
		Priority:       1,
	}

	return &p
}

package core

type Proc struct {
	Pid    int
	Burst  int
	Arrive int
}

func NewProc(pid, burst, arrive int) *Proc {
	return &Proc{
		pid,
		burst,
		arrive,
	}
}

package core

type Proc struct {
	Pid    int
	Burst  int
	InitBurst int
	Arrive int
	Priority int
}

func NewProc(pid, burst, arrive, priority int) *Proc {
	return &Proc{
		pid,
		burst,
		burst, 
		arrive,
		priority,
	}
}

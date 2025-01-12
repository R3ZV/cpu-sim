package core

type Proc struct {
	Pid          int
	Burst        int
	InitBurst    int
	Arrive       int
	Priority     int
	ParentPeriod int //for the real-time algorithms, we'll need to know how often it's made
}
type ProcGenerator struct {
	Id     int
	Burst  int
	Period int
}

func NewProc(pid, burst, arrive, priority, pperiod int) *Proc {
	return &Proc{
		pid,
		burst,
		burst,
		arrive,
		priority,
		pperiod,
	}
}
func NewGenerator(id, burst, period int) *ProcGenerator {
	return &ProcGenerator{
		id,
		burst,
		period,
	}
}

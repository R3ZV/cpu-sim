package core

type Proc struct {
	Pid    int
	Burst  int
	Arrive int
}

type metaProc struct {
	Pid    int
	Burst  int
	Period int
}

func NewProc(pid, burst, arrive int) *Proc {
	return &Proc{
		pid,
		burst,
		arrive,
	}
}

func newMetaProc(pid, burst, period int) *metaProc {
	return &metaProc{
		pid,
		burst,
		period,
	}
}

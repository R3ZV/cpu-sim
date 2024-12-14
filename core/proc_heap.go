package core

type ProcHeap struct {
	procs []Proc
	cmp   func(a, b Proc) bool
}

func (self ProcHeap) Len() int      { return len(self.procs) }
func (self ProcHeap) Swap(i, j int) { self.procs[i], self.procs[j] = self.procs[j], self.procs[i] }

func (self ProcHeap) Less(i, j int) bool { return self.cmp(self.procs[i], self.procs[j]) }
func (self *ProcHeap) Push(proc any)     { (*self).procs = append((*self).procs, proc.(Proc)) }

func (self *ProcHeap) Pop() any {
	old := (*self).procs
	n := len(old)

	item := old[n-1]

	(*self).procs = old[0 : n-1]
	return &item
}

func NewProcHeap(cmp func(a, b Proc) bool) *ProcHeap {
	return &ProcHeap{
		procs: []Proc{},
		cmp:   cmp,
	}
}

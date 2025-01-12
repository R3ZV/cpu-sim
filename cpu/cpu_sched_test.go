package cpu

import (
	"sim/core"
	"sim/log"
	"sim/sched"

	"fmt"
	"testing"
)

func TestOptimalFCFS(t *testing.T) {
	algo := sched.NewFCFS("FCFS")
	jobs := []core.Proc{
		*core.NewProc(0, 10, 0, 0, -1),
		*core.NewProc(1, 10, 0, 0, -1),
		*core.NewProc(2, 10, 0, 0, -1),
	}

	cpu := NewCPU(algo)
	cpu.AddJobs(jobs)
	for !cpu.IsDone() {
		cpu.Tick()
	}

	log.Assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

	got := fmt.Sprintf("%.2f%%", cpu.Usage())
	want := "100.00%"
	if got != want {
		t.Fatalf(`[CPU Usage]: Expected = %s, Got = %s`, want, got)
	}

	got = fmt.Sprintf("%.2f", cpu.TurnaroundTime())
	want = "20.00"
	if got != want {
		t.Fatalf(`[Turnaround Time]: Expected = %s, Got = %s`, want, got)
	}

	// TODO: add waiting time
	// got = fmt.Sprintf("%.2f", cpu.WaitTime())
	// want = "100.00"
	// if !got.MatchString(want) {
	//     t.Fatalf(`[Waiting Time]: Expected = %s, Got = %s`, want, got)
	// }
}

func TestWorstFCFS(t *testing.T) {
	algo := sched.NewFCFS("FCFS")
	jobs := []core.Proc{
		*core.NewProc(0, 100, 0, 0, -1),
		*core.NewProc(1, 10, 0, 0, -1),
		*core.NewProc(2, 10, 0, 0, -1),
	}

	cpu := NewCPU(algo)
	cpu.AddJobs(jobs)
	for !cpu.IsDone() {
		cpu.Tick()
	}

	log.Assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

	got := fmt.Sprintf("%.2f%%", cpu.Usage())
	want := "100.00%"
	if got != want {
		t.Fatalf(`[CPU Usage]: Expected = %s, Got = %s`, want, got)
	}

	got = fmt.Sprintf("%.2f", cpu.TurnaroundTime())
	want = "110.00"
	if got != want {
		t.Fatalf(`[Turnaround Time]: Expected = %s, Got = %s`, want, got)
	}

	// TODO: add waiting time
	// got = fmt.Sprintf("%.2f", cpu.WaitTime())
	// want = "100.00"
	// if !got.MatchString(want) {
	//     t.Fatalf(`[Waiting Time]: Expected = %s, Got = %s`, want, got)
	// }
}

func TestOptimalSJF(t *testing.T) {
	algo := sched.NewSJF("SJF")
	jobs := []core.Proc{
		*core.NewProc(1, 10, 0, 0, -1),
		*core.NewProc(0, 100, 0, 0, -1),
		*core.NewProc(3, 20, 0, 0, -1),
		*core.NewProc(2, 30, 0, 0, -1),
	}

	cpu := NewCPU(algo)
	cpu.AddJobs(jobs)
	for !cpu.IsDone() {
		cpu.Tick()
	}

	log.Assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

	got := fmt.Sprintf("%.2f%%", cpu.Usage())
	want := "100.00%"
	if got != want {
		t.Fatalf(`[CPU Usage]: Expected = %s, Got = %s`, want, got)
	}

	got = fmt.Sprintf("%.2f", cpu.TurnaroundTime())
	want = "65.00"
	if got != want {
		t.Fatalf(`[Turnaround Time]: Expected = %s, Got = %s`, want, got)
	}

	// TODO: add waiting time
	// got = fmt.Sprintf("%.2f", cpu.WaitTime())
	// want = "100.00"
	// if !got.MatchString(want) {
	//     t.Fatalf(`[Waiting Time]: Expected = %s, Got = %s`, want, got)
	// }
}

func TestWorstSJF(t *testing.T) {
	algo := sched.NewSJF("SJF")
	jobs := []core.Proc{
		*core.NewProc(0, 10, 10, 0, -1),
		*core.NewProc(1, 10, 10, 0, -1),
		*core.NewProc(2, 100, 0, 0, -1),
	}

	cpu := NewCPU(algo)
	cpu.AddJobs(jobs)
	for !cpu.IsDone() {
		cpu.Tick()
	}

	log.Assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

	got := fmt.Sprintf("%.2f%%", cpu.Usage())
	want := "100.00%"
	if got != want {
		t.Fatalf(`[CPU Usage]: Expected = %s, Got = %s`, want, got)
	}

	got = fmt.Sprintf("%.2f", cpu.TurnaroundTime())
	want = "103.33"
	if got != want {
		t.Fatalf(`[Turnaround Time]: Expected = %s, Got = %s`, want, got)
	}

	// TODO: add waiting time
	// got = fmt.Sprintf("%.2f", cpu.WaitTime())
	// want = "100.00"
	// if !got.MatchString(want) {
	//     t.Fatalf(`[Waiting Time]: Expected = %s, Got = %s`, want, got)
	// }
}

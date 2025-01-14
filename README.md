# Scheduling Simulator

This project simulates various real time algorithms and benchmarks their
performance in different scenarios.

Implemented algorithms:

 - [x] First Come, First Served (FCFS)
 - [x] Shortest Job First (SJF)
 - [x] Preemptive Shortest Job First (PSJF)
 - [x] Priority based
 - [x] Rate monotonic (RM)
 - [x] Earliest Deadline First (EDF)
 - [ ] Round Robin (RR)

# Metrics

In order to compare scheduling algorithms we will evaluate them based on the
following metrics:

- [x] Turnaround time (T_completion - T_arrival)
- [x] Waiting time (T_turnaround - T_burst)
- [x] Response time (T_first_on_cpu - T_arrival)

# Results

![Workload 1 Usage.jpg](plots/Workload_1_Usage.jpg)
![Workload 1 Turnaround.jpg](plots/Workload_1_Turnaround.jpg)
![Workload 1 Waiting.jpg](plots/Workload_1_Waiting.jpg)
![Workload 1 Responding.jpg](plots/Workload_1_Responding.jpg)
![Workload 2 Usage.jpg](plots/Workload_2_Usage.jpg)
![Workload 2 Turnaround.jpg](plots/Workload_2_Turnaround.jpg)
![Workload 2 Waiting.jpg](plots/Workload_2_Waiting.jpg)
![Workload 2 Responding.jpg](plots/Workload_2_Responding.jpg)
![Workload 3 Usage.jpg](plots/Workload_3_Usage.jpg)
![Workload 3 Turnaround.jpg](plots/Workload_3_Turnaround.jpg)
![Workload 3 Waiting.jpg](plots/Workload_3_Waiting.jpg)
![Workload 3 Responding.jpg](plots/Workload_3_Responding.jpg)
![Workload 4 Usage.jpg](plots/Workload_4_Usage.jpg)
![Workload 4 Turnaround.jpg](plots/Workload_4_Turnaround.jpg)
![Workload 4 Waiting.jpg](plots/Workload_4_Waiting.jpg)
![Workload 4 Responding.jpg](plots/Workload_4_Responding.jpg)
![Workload 5 Usage.jpg](plots/Workload_5_Usage.jpg)
![Workload 5 Turnaround.jpg](plots/Workload_5_Turnaround.jpg)
![Workload 5 Waiting.jpg](plots/Workload_5_Waiting.jpg)
![Workload 5 Responding.jpg](plots/Workload_5_Responding.jpg)

# References
- https://pages.cs.wisc.edu/~remzi/OSTEP/cpu-sched.pdf
- Operating System Concepts, 10th Edition
by Abraham Silberschatz, Peter B. Galvin, Greg Gagne
- Liu, C. L.; Layland, J. (1973), "Scheduling algorithms for multiprogramming in a hard real-time environment", Journal of the ACM, 20 (1), pp. 46–61

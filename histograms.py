import matplotlib.pyplot as plt
import os

output_dir = "plots"
if not os.path.exists(output_dir):
    os.makedirs(output_dir)


# these can be all sorts of metrics
def plot_workload(name, metric, listAlgs, listTimes, hist_color):
    save_name_original = name + " " + metric + ".jpg"
    save_name = save_name_original.replace(" ", "_")
    save_path = os.path.join(output_dir, save_name)
    print(f"![{save_name_original}]({save_path})")
    plt.bar(listAlgs, listTimes, color=hist_color)
    plt.xlabel(name)
    plt.title(metric)
    plt.savefig(save_path, format="jpg", dpi=300)
    plt.clf()

    
class Workload:
    algorithms = []
    usages = []
    tas = []
    waits = []
    resps = []

f = open("data", "r")
content = f.readlines()
count_loads = 0
for line in content:
    words = line.split()
    if len(words) > 0 and words[0] == "Workload":
        count_loads = max(int(words[1][0:-1]), count_loads)

# workloads in the test data are 1-indexed, arrays are 0-indexed,
loads = [Workload() for i in range(count_loads + 1)]
for line in content:
    words = line.split()
    if len(words) > 0:
        if words[0] == "Workload":
            current_load = int(words[1][0:-1])
        elif words[0] == "Testing":
            for load in loads:
                load.algorithms.append(words[1])
        elif words[0] == "Usage:":
            loads[current_load].usages.append(float(words[1][0:-1]))
        elif words[0] == "Turnaround":
            loads[current_load].tas.append(float(words[2]))
        elif words[0] == "Waiting":
            loads[current_load].waits.append(float(words[2]))
        elif words[0] == "Response":
            loads[current_load].waits.append(float(words[2]))

for i in range(1, count_loads + 1):
    plot_workload(f"Workload {i}", "Usage", loads[i].algorithms, loads[i].usages, "green")
    plot_workload(f"Workload {i}", "Turnaround", loads[i].algorithms, loads[i].tas, "blue")
    plot_workload(f"Workload {i}", "Waiting", loads[i].algorithms, loads[i].waits, "red")
    plot_workload(f"Workload {i}", "Responding", loads[i].algorithms, loads[i].waits, "darkviolet")
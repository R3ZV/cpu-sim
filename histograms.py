import matplotlib.pyplot as plt
import os

output_dir = "plots"
if not os.path.exists(output_dir):
    os.makedirs(output_dir)

def plot_workload(name, metric, listAlgs, listTimes): # these can be all sots of times
    save_name=name+'_'+metric+'.jpg'
    save_path=os.path.join(output_dir, save_name)
    plt.bar(listAlgs, listTimes, color='red')
    plt.xlabel(name)
    plt.title(metric)
    plt.savefig(save_path, format='jpg', dpi=300)

class Workload: 
    algorithms=[]
    usages=[]
    tas=[]
    waits=[]

f=open("data", "r")
content=f.readlines()
loads=[Workload() for i in range(count_loads+1)] #workloads in the test data are 1-indexed, arrays are 0-indexed, 
count_loads=0
for line in content:
   words=line.split()
   if(len(words)>0 and words[0]=='Workload'):
      count_loads=max(int(words[1][0:-1]), count_loads)
      if(words[0]=='Workload'):
         current_load=int(words[1][0:-1])
      elif words[0]=='Testing':
         for load in loads:
            load.algorithms.append(words[1])
      elif words[0]=='Usage:':
         loads[current_load].usages.append(float(words[1][0:-1]))
      elif words[0]=='Turnaround':
         loads[current_load].tas.append(float(words[2]))
      elif words[0]=='Waiting':
         loads[current_load].waits.append(float(words[2]))

for i in range(1, count_loads+1):
    plot_workload(f"Workload {i}", "Usage", loads[i].algorithms, loads[i].usages)
    plot_workload(f"Workload {i}", "Turnaround", loads[i].algorithms, loads[i].tas)
    plot_workload(f"Workload {i}", "Waiting", loads[i].algorithms, loads[i].waits)

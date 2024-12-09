package main

import ("fmt"
         "container/heap")

type proces struct{
    pid int
    burst int
    remainingBurst int
    period int //simulam un CPU embedded, asa ca vom sti toate procesele, si ele se repeta
    priority int   
}
func newProces(pid int, burst int, period int) *proces{
    p := proces{pid:pid , burst:burst, remainingBurst:burst, period:period, priority:1} 
    return &p //si schema de milioane e ca el NU IESE DIN SCOPE
}
func compareProces(a,b proces) bool{
   return a.burst<b.burst
}

//implementam structura care tine procesele, cu metodele aferente, poate ca asta ar trebui sa fie in alt fisier
type procesHeap []proces
func (h procesHeap) Len() int {
 return len(h)
}
func (h procesHeap) Less(i, j int) bool {
    return compareProces(h[i], h[j])
}
func (h procesHeap) Swap(i, j int) {
   h[i], h[j]= h[j], h[i]
} 
///restul nu mai au surprize
func (h *procesHeap) Push(x interface{}) {
	*h = append(*h, x.(proces))
}

func (h *procesHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func rateMonotonic(procese []proces){
   var time=0
   var busyTime=0
   const MAX_TIME=200 //TODO fa-l pe asta LCM, 200 e asa de test
   q := &procesHeap{}
   heap.Init(q)
   for i:=0; i<len(procese); i++{
      heap.Push(q, procese[i]) 
   }
   for time<MAX_TIME{
      if(q.Len()>0){//daca e ceva ce se poate executa, se executa
          busyTime++
          var currProces proces
          currProces=heap.Pop(q).(proces)
          currProces.remainingBurst--///calculam cat burst mai are 
          if(currProces.remainingBurst!=0){
             heap.Push(q, currProces)
          }
      }
      for i:=0; i<len(procese); i++{
         if(time % procese[i].period == 0){///mai baga unul la rand, daca i-a venit sorocul
             heap.Push(q, procese[i]);
         }
      }
      time++//TODO vezi daca nu se poate simula mai eficient, poate sa deducem cand va fi eveniment cu alt heap
   }
}
func main(){
    var s=newProces(1, 2, 3) ///TODO scrie teste pentru rate-monotonic scheduling
    fmt.Println(s.burst)
}
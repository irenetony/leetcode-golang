package main

import "fmt"

const (
	Ready   = 0
	Cooling = 1
)

type TasksStatus struct {
	Ptr            int
	Status         int
	CoolingCount   int
	RemainingCount int
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	res := 0
	heapSortByte(tasks)

	tasksStatusMap := make(map[byte]*TasksStatus)
	coolingTasksMap := make(map[byte]bool)
	readyTasksMap := make(map[byte]bool)
	var lastTasks byte = 0
	for i := 0; i < len(tasks); i++ {
		if tasks[i] != lastTasks {
			lastTasks = tasks[i]
			tasksStatusMap[tasks[i]] = &TasksStatus{Ptr: i, Status: Ready, RemainingCount: 1}
			readyTasksMap[tasks[i]] = true
		} else {
			tasksStatusMap[tasks[i]].RemainingCount++
		}
	}

	for len(tasksStatusMap) > 0 {
		res++
		if len(readyTasksMap) > 0 { // run one of longest ready tasks
			maxRemainTask := 0
			var task byte
			for t := range readyTasksMap {
				if tasksStatusMap[t].RemainingCount > maxRemainTask {
					maxRemainTask = tasksStatusMap[t].RemainingCount
					task = t
				}
			}
			status := tasksStatusMap[task]
			status.Status = Cooling
			status.CoolingCount = n
			status.Ptr++
			status.RemainingCount--
			delete(readyTasksMap, task)
			coolingTasksMap[task] = true
			if status.Ptr == len(tasks) || tasks[status.Ptr] != task { // task done
				delete(coolingTasksMap, task)
				delete(tasksStatusMap, task)
			}
			for coolingTask := range coolingTasksMap {
				if coolingTask != task {
					tasksStatusMap[coolingTask].CoolingCount--
					if tasksStatusMap[coolingTask].CoolingCount == 0 {
						tasksStatusMap[coolingTask].Status = Ready
						readyTasksMap[coolingTask] = true
						delete(coolingTasksMap, coolingTask)
					}
				}
			}
		} else { // idle
			for coolingTask := range coolingTasksMap {
				tasksStatusMap[coolingTask].CoolingCount--
				if tasksStatusMap[coolingTask].CoolingCount == 0 {
					tasksStatusMap[coolingTask].Status = Ready
					readyTasksMap[coolingTask] = true
					delete(coolingTasksMap, coolingTask)
				}
			}
		}
	}

	return res
}

func heapSortByte(arr []byte) {
	initHeapByte(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr, 0, i)
	}
}

func initHeapByte(arr []byte) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		shiftDown(arr, i, len(arr))
	}
}

func shiftDown(arr []byte, idx, limit int) {
	for idx <= limit/2-1 {
		if 2*idx+2 < limit { // has right child
			if arr[2*idx+1] >= arr[2*idx+2] && arr[idx] <= arr[2*idx+1] {
				arr[idx], arr[2*idx+1] = arr[2*idx+1], arr[idx]
				idx = 2*idx + 1
			} else if arr[2*idx+1] <= arr[2*idx+2] && arr[idx] <= arr[2*idx+2] {
				arr[idx], arr[2*idx+2] = arr[2*idx+2], arr[idx]
				idx = 2*idx + 2
			} else {
				break
			}
		} else if 2*idx+1 < limit { // has only left child
			if arr[2*idx+1] >= arr[idx] {
				arr[2*idx+1], arr[idx] = arr[idx], arr[2*idx+1]
				idx = 2*idx + 1
			} else {
				break
			}
		} else {
			break
		}
	}
}

func main() {
	arr := []byte{'D', 'D', 'A', 'C', 'C', 'B', 'A', 'B'}
	//arr := []byte{'A', 'B', 'C', 'D'}
	//arr := []byte{'D', 'A', 'B', 'C'}
	heapSortByte(arr)
	fmt.Printf("%+v\n", arr)
	fmt.Printf("res=%+v", leastInterval([]byte{'A', 'B', 'C', 'A', 'B'}, 2))
}

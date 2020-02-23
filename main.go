package main

import (
	"codility/stack/solution"
	"fmt"
)

func main() {
	sol := solution.Solution{Stack: make([]int, 0)}
	sol.Push(4)
	sol.Begin() // start transaction 1
	sol.Push(7) // stack: 7, 4
	sol.Begin() // start transaction 2
	sol.Push(2) // stack: 2, 7, 4
	fmt.Printf("%t\n", sol.Rollback()) // rollback transaction 2, will be true, stack: 7, 4
	fmt.Printf("%d\n", sol.Top()) // will be 7
	sol.Begin() // start transaction 3
	sol.Push(10) // stack: 10, 7, 4
	sol.Commit() // commit transaction 3, stack: 10, 7, 4
	fmt.Printf("%d\n", sol.Top()) // will be 10
	sol.Rollback() // rollback transaction 1, will be true, stack: 4
	fmt.Printf("%d\n", sol.Top()) // will be 4
	fmt.Printf("%t", sol.Commit()) // will be false, no open transaction
}
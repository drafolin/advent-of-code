package day10

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aclements/go-z3/z3"
)

type node[E any] struct {
	state    []E
	distance int
}

type machine struct {
	requiredState      []bool
	buttons            [][]int
	joltageRequirement []int
}

func machineFromString(s string) machine {
	parts := strings.Split(s, " ")
	requiredState := make([]bool, len(parts[0])-2)
	for i, char := range parts[0][1 : len(parts[0])-1] {
		requiredState[i] = char == '#'
	}

	buttons := make([][]int, len(parts[1:len(parts)-1]))
	for i, button := range parts[1 : len(parts)-1] {
		buttonActions := strings.Split(button[1:len(button)-1], ",")
		buttons[i] = make([]int, len(buttonActions))
		for j, action := range buttonActions {
			buttons[i][j], _ = strconv.Atoi(action)
		}
	}

	joltageRequirementString := parts[len(parts)-1][1 : len(parts[len(parts)-1])-1]
	joltageRequirementsParts := strings.Split(joltageRequirementString, ",")
	joltageRequirements := make([]int, len(joltageRequirementsParts))
	for i, requirement := range joltageRequirementsParts {
		joltageRequirements[i], _ = strconv.Atoi(string(requirement))
	}

	return machine{
		requiredState:      requiredState,
		buttons:            buttons,
		joltageRequirement: joltageRequirements,
	}
}

func Main() {
	f, err := os.ReadFile("day10/input")
	if err != nil {
		panic(err)
	}

	s := string(f)
	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")

	machines := make([]machine, 0)
	for _, line := range lines {
		machines = append(machines, machineFromString(line))
	}

	timeStart := time.Now()
	res := firstPart(machines)
	timeEnd := time.Now()
	fmt.Println("First part took", timeEnd.Sub(timeStart))
	fmt.Println("First part result: ", res)

	timeStart = time.Now()
	res = secondPart(machines)
	timeEnd = time.Now()
	fmt.Println("Second part took", timeEnd.Sub(timeStart))
	fmt.Println("Second part result: ", res)
}

func firstPart(machines []machine) int {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, machine := range machines {
		wg.Go(func() {
			startNode := node[bool]{
				state:    make([]bool, len(machine.requiredState)),
				distance: 0,
			}
			nodes := []node[bool]{startNode}

			for len(nodes) > 0 {
				currentNode := slices.MinFunc(nodes, func(a, b node[bool]) int {
					return a.distance - b.distance
				})

				if slices.Equal(currentNode.state, machine.requiredState) {
					ch <- currentNode.distance
					return
				}

				for _, button := range machine.buttons {
					newState := make([]bool, len(currentNode.state))
					copy(newState, currentNode.state)
					for _, action := range button {
						newState[action] = !newState[action]
					}

					if neighborIndex := slices.IndexFunc(nodes, func(n node[bool]) bool {
						return slices.Equal(n.state, newState)
					}); neighborIndex != -1 {
						nodes[neighborIndex].distance = min(nodes[neighborIndex].distance, currentNode.distance+1)
					} else {
						nodes = append(nodes, node[bool]{
							state:    newState,
							distance: currentNode.distance + 1,
						})
					}
				}

				index := slices.IndexFunc(nodes, func(n node[bool]) bool {
					return slices.Equal(n.state, currentNode.state)
				})

				nodes = append(nodes[:index], nodes[index+1:]...)
			}
		})
	}
	total := 0
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		total += i
	}
	return total
}

func secondPart(machines []machine) int {
	total := 0
	ctx := z3.NewContext(nil)

	for _, machine := range machines {
		solver := z3.NewSolver(ctx)

		array := []z3.Int{}
		for i := range machine.buttons {
			variable := ctx.Const(fmt.Sprintf("x%d", i), ctx.IntSort()).(z3.Int)
			solver.Assert(variable.GE(ctx.FromInt(0, ctx.IntSort()).(z3.Int)))
			array = append(array, variable)
		}

		for i, joltage := range machine.joltageRequirement {
			ok := []z3.Int{}
			for j, variable := range array {
				if slices.Contains(machine.buttons[j], i) {
					ok = append(ok, variable)
				}
			}

			solver.Assert(ok[0].Add(ok[1:]...).Eq(ctx.FromInt(int64(joltage), ctx.IntSort()).(z3.Int)))
		}

		sum := array[0]
		for _, variable := range array[1:] {
			sum = sum.Add(variable)
		}

		if v, _ := solver.Check(); !v {
			panic("no solution found")
		}

		var res z3.Int
		for {
			// Once no more solutions can be found, break
			if v, err := solver.Check(); !v || err != nil {
				break
			}

			model := solver.Model()
			res = model.Eval(sum, true).(z3.Int)

			// Try to find a smaller solution
			solver.Assert(sum.LT(res))
		}

		resInt, _, _ := res.AsInt64()
		total += int(resInt)
	}
	return total
}

func bfsSecondPart(machines []machine) int {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, machine := range machines {
		wg.Go(func() {
			state := make([]int, len(machine.joltageRequirement))
			copy(state, machine.joltageRequirement)
			startNode := &node[int]{
				state:    state,
				distance: 0,
			}
			queue := []*node[int]{startNode}
			explored := [][]int{startNode.state}

			for len(queue) > 0 {
				currentNode := queue[0]
				queue = queue[1:]

				if slices.Max(currentNode.state) == 0 {
					ch <- currentNode.distance
					return
				}

			buttonLoop:
				for _, button := range machine.buttons {
					newState := make([]int, len(currentNode.state))
					copy(newState, currentNode.state)
					for _, action := range button {
						newValue := newState[action] - 1
						if newValue < 0 {
							// This path can't ever get true
							continue buttonLoop
						}
						newState[action] = newValue
					}

					if slices.ContainsFunc(explored, func(s []int) bool {
						return slices.Equal(s, newState)
					}) {
						continue
					}

					explored = append(explored, newState)

					queue = append(queue, &node[int]{
						state:    newState,
						distance: currentNode.distance + 1,
					})
				}
			}
		})
	}

	total := 0
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		total += i
	}
	return total
}

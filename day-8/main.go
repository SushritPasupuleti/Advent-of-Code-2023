package main

import (
	// "bufio"
	"fmt"
	"os"
	"slices"
	_ "slices"
	"strings"
)

type TreeNode struct {
	Left     *TreeNode
	LeftVal  string
	Right    *TreeNode
	RightVal string
	Value    string
}

func (t TreeNode) String() string {
	return fmt.Sprintf("Node: %s Left: %v Right: %v LeftV: %s RightV: %s", t.Value, t.Left, t.Right, t.LeftVal, t.RightVal)
}

func main() {
	//Format: LRLRLRLRRRLLRRRR
	//
	//		  AAA = (BBB, CCC)
	//		  BBB = (DDD, EEE)

	// input, _ := os.ReadFile("input_test.txt")
	// input, _ := os.ReadFile("input_test2.txt")
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("input_test_p2.txt")

	tree, directions := parseTree(string(input))

	// fmt.Println("Tree: ", tree)
	fmt.Println("Directions: ", directions)
	// treeX := findNodeEndingWithX(tree, "A", "1")
	// treeX = findNodeEndingWithX(treeX, "Z", "2")
	// fmt.Println("TreeX: ", len(treeX))

	steps := traverseFromNodeToNode(tree, "AAA", "ZZZ", directions)

	fmt.Println("Part 1: Steps: ", steps)

	wA := findNodesEndingWithX(tree, "A", "1")
	wZ := findNodesEndingWithX(tree, "Z", "2")
	fmt.Println("wA: ", wA)
	fmt.Println("wZ: ", wZ)

	steps = traverseFromNodeToNodeList(tree, wA, wZ, directions)
	fmt.Println("Part 2: Steps: ", steps)

	// steps := traverseFromNodeToNodePt2(tree, "AAA", "ZZZ", directions)
	// fmt.Println("Part 2: Steps: ", steps)
}

func parseTree(input string) (map[string]*TreeNode, []string) {

	nodes := make(map[string]*TreeNode)
	directions := make([]string, 0)

	lines := strings.Split(input, "\n")

	for i, line := range lines {

		if line == "" {
			continue
		}

		if i == 0 {
			directions = parseDirections(line)
		}

		if i != 0 {
			node := parseNode(line)
			nodes[node.Value] = &node
		}
	}

	// for _, node := range nodes {
		// fmt.Println("Node: ", node)
	// }

	return nodes, directions
}

// Returns an array of directions
// Input Format: LRLRLRLRRRLLRRRR
func parseDirections(directions string) []string {

	directionsArr := strings.Split(directions, "")

	return directionsArr
}

// Returns a node
// Input Format: AAA = (BBB, CCC)
func parseNode(node string) TreeNode {

	valSplit := strings.Split(node, " = ")

	valStr := valSplit[0]
	rHS := valSplit[1]

	lrSplit := strings.Split(rHS, ", ")

	//(BBB, CCC)
	lrStr := (lrSplit[0])[1:]                //Trim the first (
	rStr := (lrSplit[1])[:len(lrSplit[1])-1] //Trim the last )

	// fmt.Println("LR: ", lrStr)
	// fmt.Println("R: ", rStr)

	return TreeNode{
		Left:     nil,
		LeftVal:  lrStr,
		Right:    nil,
		RightVal: rStr,
		Value:    valStr,
	}
}

func traverseFromNodeToNode(nodes map[string]*TreeNode, start string, end string, directions []string) int {

	steps := 0
	found := false
	for found == false {
		for _, direction := range directions {
			if direction == "L" {
				start = nodes[start].LeftVal
				// fmt.Println("Start(L): ", start)
				steps++
			} else if direction == "R" {
				start = nodes[start].RightVal
				// fmt.Println("Start(R): ", start)
				steps++
			}

			if start == end {
				found = true
				// fmt.Println("Found: ", start)
				break
			}
		}
	}

	fmt.Println("Count: ", steps)

	return steps
}

func findNodesEndingWithX(nodes map[string]*TreeNode, end string, rep string) []string {

	var foundNodes []string
	// fmt.Println("End: ", end)
	for _, node := range nodes {
		if strings.HasSuffix(node.Value, end) {
			// fmt.Println("Node: ", node)
			foundNodes = append(foundNodes, node.Value)
		} else {
		}
	}

	return foundNodes
}

func traverseFromNodeToNodeList(nodes map[string]*TreeNode, start []string, end []string, directions []string) int {

	// steps := 0
	stepResults := make(map[string]int)

	for _, startNode := range start {
		steps := 0
		found := false
		for found == false {
			startNode = startNode
			for _, direction := range directions {
				if direction == "L" {
					startNode = nodes[startNode].LeftVal
					// fmt.Println("Start(L): ", startNode)
					steps++
				} else if direction == "R" {
					startNode = nodes[startNode].RightVal
					// fmt.Println("Start(R): ", startNode)
					steps++
				}

				if slices.Contains(end, startNode) {
					found = true
					// fmt.Println("Found: ", startNode)
					stepResults[startNode] = steps
					break
				}
			}
		}
	}

	// fmt.Println("Count: ", steps)
	steps := make([]int, 0)
	for _, v := range stepResults {
		// fmt.Println("Key: ", k, " Value: ", v)
		steps = append(steps, v)
	}

	lcmResult := lcmArray(steps)

	// fmt.Println("LCM: ", lcmResult)

	return lcmResult
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
		// fmt.Println("a: ", a, " b: ", b)
	}
	return a
}

func lcmArray(arr []int) int {
	lcm := 1
	for _, v := range arr {
		lcm = lcm * v / gcd(lcm, v)
	}
	return lcm
}

package main

import "fmt"

// Nodes a generic type for slice of any type.

//type Nodes[T toWork] []T

// Sort sorts does not mind about the fields or values of T.
func Sort(input []toWork) {
	length := len(input)

	var i1, i2 int
	//rightGap = float64(length - 1)
	for gap := length - 1; gap > 0; gap -= shrinkFactorReducer(gap) {
		for i1, i2 = 0, gap; i2 < length; i1++ {
			if input[i1].id > input[i2].id {
				//log.Println("changing", input[i1], input[i2])
				input[i1], input[i2] = input[i2], input[i1]
			}
			i2++
		}
	}
	//log.Println(input)
}

type Node struct {
	leftNode  *Node
	rightNode *Node
	id        int
}

type treeOfNodes struct {
	*Node
	input []toWork
}

func SortByTree(input []toWork) {
	var tree treeOfNodes
	for _, value := range input {
		if tree.Node != nil {
			tree.addNode(value.id)
		} else {
			tree.Node = &Node{id: value.id}
		}
	}
	tree.Print(input, 0)
	//log.Println(input)
}

func (n *Node) Print(input []toWork, idx int) {
	if n != nil {
		n.leftNode.Print(input, idx)
		input[idx].id = n.id
		idx++
		n.rightNode.Print(input, idx)
	}
}

func (n *Node) addNode(id int) {
	if id > n.id {
		if n.rightNode == nil {
			n.rightNode = &Node{id: id}
		} else {
			n.rightNode.addNode(id)
		}
		return
	}

	if n.leftNode == nil {
		n.leftNode = &Node{id: id}
	} else {
		n.leftNode.addNode(id)
	}
}

var rightGap float64

func shrinkFactorReducer(gap int) int {
	//rightGap = float64(int((rightGap / 1.3) + .5))
	if gap < 9 {
		return 1
	}
	return int(float64(gap) / 1.3)
}

type toWork struct {
	id int
}

type toWorkSLice []toWork

func (t toWorkSLice) Less(i int, j int) bool {
	return t[i].id < t[j].id
}

func (t toWorkSLice) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t toWorkSLice) Len() int {
	return len(t)
}

var testSlice = toWorkSLice{
	{id: 12},
	{id: 10},
	{id: 4},
	{id: 2},
	{id: 3},
	{id: 20},
	{id: 22},
	{id: 25},
	{id: 5},
	{id: 9},
	{id: 15},
	{id: 222},
	{id: 1},
}

var rightSlice = toWorkSLice{
	{id: 1},
	{id: 2},
	{id: 3},
	{id: 4},
	{id: 5},
	{id: 9},
	{id: 10},
	{id: 12},
	{id: 15},
	{id: 20},
	{id: 22},
	{id: 25},
	{id: 222},
}

func main() {

	Sort(testSlice)

	//var testSlice2 = make([]toWork, 5000)
	//for key := range testSlice2 {
	//	testSlice2[key].id = rand.Int()
	//}
	//Sort(testSlice2)
	//for key := range testSlice2 {
	//	if key != 0 {
	//		if testSlice2[key].id < testSlice2[key-1].id {
	//			fmt.Println("algorithm does not work")
	//		}
	//	}
	//}

	SortByTree(testSlice)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || haystack == needle {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	var lastHaystackIndex = len(haystack) - len(needle) + 1
	fmt.Println("needle:", needle)
	fmt.Println("----")
	for key := 0; key < lastHaystackIndex; key++ {
		fmt.Println("sample:", haystack[key:key+len(needle)])
		if haystack[key:key+len(needle)] == needle {
			return key
		}
	}
	return -1
}

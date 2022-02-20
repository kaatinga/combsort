package main

// Nodes a generic type for slice of any type.

//type Nodes[T toWork] []T

// Sort sorts does not mind about the fields or values of T.
func Sort(input []toWork) {
	length := len(input)

	var i1, i2 int
	rightGap = float64(length - 1)
	for gap := length - 1; gap > 0; gap -= reducer(gap) {
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

func reducer(gap int) int {
	rightGap = float64(int((rightGap / 1.3) + .5))
	if gap < 9 {
		return 1
	}
	return gap >> 2
}

type toWork struct {
	id int
}

var testSlice = []toWork{
	{id: 4},
	{id: 10},
	{id: 12},
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

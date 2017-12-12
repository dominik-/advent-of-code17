package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day7() {
	scanner := GetScanner("inputs/input7")
	//currentRoots := make([]*Node, 0)
	allNodes := make(map[string]*Node, 0)
	childRefsMap := make(map[string]int)
	for {
		if scanner.Scan() {
			node := parseLine(scanner.Text())
			allNodes[node.Name] = node
			for _, childRef := range node.ChildRefs {
				childRefsMap[childRef]++
			}
			//fmt.Printf("Node %v with children: %+v\n", node, node.ChildRefs)
		} else {
			break
		}
	}
	fmt.Printf("We have %d nodes!\n", len(allNodes))
	fmt.Printf("ChildMap size: %d\n", len(childRefsMap))
	//we find the root node by getting all nodeRefs and finding the node thats not in there, but still has children
	var root *Node
	for _, n := range allNodes {
		//node has children
		if len(n.ChildRefs) > 0 {
			//node itself is not a child-node
			_, found := childRefsMap[n.Name]
			if !found {
				root = n
				fmt.Printf("Found root: %v\n", n)
				break
			}
		}
	}
	root.buildChildren(allNodes)
	difference, differrerrererer := root.isImbalancedBy(0, nil)
	fmt.Printf("Program imbalanced by %d, node to subtract difference from: %v!\n", difference, differrerrererer)
}

func (n *Node) buildChildren(allNodes map[string]*Node) {
	for _, childRef := range n.ChildRefs {
		n.Children = append(n.Children, allNodes[childRef])
	}
	for _, child := range n.Children {
		child.buildChildren(allNodes)
	}
}

func (n *Node) getWeight() (totalWeight int) {
	for _, child := range n.Children {
		totalWeight += child.getWeight()
	}
	return totalWeight + n.Weight
}

func (n *Node) isImbalancedBy(previousDifference int, previousNode *Node) (difference int, differentNode *Node) {
	difference = previousDifference
	differentNode = previousNode
	weightsCountMap := make(map[int][]*Node)
	//get weights of all children
	for _, child := range n.Children {
		_, exists := weightsCountMap[child.getWeight()]
		if !exists {
			weightsCountMap[child.getWeight()] = make([]*Node, 0)
		}
		weightsCountMap[child.getWeight()] = append(weightsCountMap[child.getWeight()], child)
	}
	//fmt.Printf("Weights in map: %+v\n", weightsCountMap)
	//check whether there is more than one weight, if so, go further down in tree
	if len(weightsCountMap) > 1 {
		differenceIndex := 0
		weights := make([]int, 0)
		for weight, nodeset := range weightsCountMap {
			weights = append(weights, weight)
			if len(nodeset) == 1 {
				differentNode = nodeset[0]
				differenceIndex = len(weights) - 1
			}
		}
		if differenceIndex == 0 {
			difference = weights[0] - weights[1]
		} else {
			difference = weights[1] - weights[0]
		}
		fmt.Printf("Weights: %+d\n", weights)
		// if there are only two subtrees, we need to check both
		if len(n.Children) == 2 {
			fmt.Println("Tow subtrees!")
			imbalance1, _ := n.Children[0].isImbalancedBy(difference, differentNode)
			if imbalance1 != difference {
				return n.Children[0].isImbalancedBy(difference, differentNode)
			}
			return n.Children[1].isImbalancedBy(difference, differentNode)
		}
		//otherwise (i.e. more than 2), we find the one with length of nodes == 1
		fmt.Println("n subtrees!")
		fmt.Printf("%+v\n", weightsCountMap)
		for _, nodes := range weightsCountMap {
			if len(nodes) == 1 {
				return nodes[0].isImbalancedBy(difference, differentNode)
			}
		}
	}
	return difference, differentNode
}

func parseLine(line string) *Node {
	words := strings.Split(line, " ")

	weight, _ := strconv.Atoi(strings.Trim(words[1], "()"))

	newNode := &Node{
		Weight:    weight,
		Name:      words[0],
		ChildRefs: make([]string, 0),
		Children:  make([]*Node, 0),
	}

	if len(words) > 2 {
		//cut off the first 3 strings
		words = words[3:]
		//remove trailing commas from node references
		for _, w := range words {
			newNode.ChildRefs = append(newNode.ChildRefs, strings.TrimRight(w, ","))
		}
	}
	return newNode
}

type Node struct {
	Parent    *Node
	Children  []*Node
	ChildRefs []string
	Name      string
	Weight    int
}

func (n *Node) String() string {
	return fmt.Sprintf("Name: %s, Weight: %d", n.Name, n.Weight)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BinaryTree struct {
	Value int         `json:"value"`
	Left  *BinaryTree `json:"left"`
	Right *BinaryTree `json:"right"`
}

type TreeInput struct {
	Nodes []BinaryTree `json:"nodes"`
	Root  string       `json:"root"`
}

type MaxPathSumResponse struct {
	MaxPathSum int `json:"maxPathSum"`
}

func maxPathSum(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	leftSum := max(0, maxPathSum(root.Left))
	rightSum := max(0, maxPathSum(root.Right))

	maxSum := leftSum + rightSum + root.Value

	return max(maxSum, max(leftSum, rightSum)+root.Value)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func handleMaxPathSum(w http.ResponseWriter, r *http.Request) {
	var treeInput TreeInput
	err := json.NewDecoder(r.Body).Decode(&treeInput)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	fmt.Println(treeInput.Nodes)
	root := buildTree(treeInput.Nodes, treeInput.Root)

	result := MaxPathSumResponse{MaxPathSum: maxPathSum(root)}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func buildTree(nodes []BinaryTree, rootID string) *BinaryTree {
	nodeMap := make(map[string]*BinaryTree)

	for _, node := range nodes {
		nodeMap[fmt.Sprintf("%v", node.Value)] = &node
	}

	for _, node := range nodes {
		if node.Left != nil {
			node.Left = nodeMap[fmt.Sprintf("%v", node.Left.Value)]
		}
		if node.Right != nil {
			node.Right = nodeMap[fmt.Sprintf("%v", node.Right.Value)]
		}
	}

	return nodeMap[rootID]
}

func main() {
	http.HandleFunc("/maxPathSum", handleMaxPathSum)

	fmt.Printf("Server running on :%d\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}

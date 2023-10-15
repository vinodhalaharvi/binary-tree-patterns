package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func DotToPng(s string, s2 string) {
	err2 := exec.Command("dot", "-Tpng", s, "-o", s2).Run()
	if err2 != nil {
		os.Exit(1)
	}
}

func PrintDotFileToFile(root *TreeNode, s string) {
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "digraph G {\n")
	printDotFileToFile(root, w)
	fmt.Fprintf(w, "}\n")
	w.Flush()
}

func printDotFileToFile(root *TreeNode, w *bufio.Writer) {
	if root == nil {
		return
	}
	if root.Left != nil {
		fmt.Fprintf(w, "\t%d -> %d;\n", root.Val, root.Left.Val)
	}
	if root.Right != nil {
		fmt.Fprintf(w, "\t%d -> %d;\n", root.Val, root.Right.Val)
	}
	printDotFileToFile(root.Left, w)
	printDotFileToFile(root.Right, w)
}

func printDotFile(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		fmt.Printf("\t%d -> %d;\n", root.Val, root.Left.Val)
	}
	if root.Right != nil {
		fmt.Printf("\t%d -> %d;\n", root.Val, root.Right.Val)
	}
	printDotFile(root.Left)
	printDotFile(root.Right)
}

func DrawInputTree(root *TreeNode) {
	dotFileName := "spiralorder_traversal.dot"
	pngFileName := "spiralorder_traversal.png"
	PrintDotFileToFile(root, dotFileName)
	DotToPng(dotFileName, pngFileName)
}

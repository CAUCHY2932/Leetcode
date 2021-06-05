package main

import "fmt"

type Tree struct{
	Left *Tree
	Right *Tree
	Val int
}

func BuildTree(arr []int, index int)*Tree{
	if index >= len(arr){
		return nil
	}
	if arr[index] == -100 { // 理论上给的都是节点，不是值
		return nil
	}
	dummy := newNode(arr[index])
	dummy.Left = BuildTree(arr, index << 1 + 1)
	dummy.Right = BuildTree(arr, index << 1 + 2)
	return dummy
}

func newNode(val int)*Tree{
	return &Tree{
		Left: nil,
		Right: nil,
		Val: val,
	}
}


func InOrderPrint(root *Tree){
	if root == nil {
		return
	}
	InOrderPrint(root.Left)
	fmt.Printf("[cur] print item: %v\n", root.Val)
	InOrderPrint(root.Right)
}

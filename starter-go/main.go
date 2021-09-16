package main // main must be package main

import "fmt" // imports must be grouped together

func main() { // main function must be named main

	/* GO LANG 101 */
	/*USING IF ELSE CONDITIONALS*/
	// x := 42           // x is an int and can be declared with x :=  42 or var x int = 42
	// y := "James Bond" // y is a string and can be declared with y := "James Bond" or var y string = "James Bond"
	// z := false        // z is a bool and can be declared with z := true or var z bool = true

	// fmt.Println(x) // prints 42
	// if x == 42 {   // if x is equal to 42 then print y
	// 	fmt.Println(y)
	// } else if x > 30 {
	// 	fmt.Println(y)
	// }
	// if z != true {
	// 	fmt.Println("the answer is: ", z)
	// }

	/* MUTATIONS*/
	// mutations are O(1) time complexity because we are only changing one element and we know the index

	// a[1] = 100 // change the value of a[1] to 100
	// fmt.Println(a)

	// b[0] = 45 // b[0] is equal to 45 can by assigning the index number 2 of b to 45
	// fmt.Println(b)

	// c[2] = 55 // c[2] is equal to 55 can by assigning the index number 2 of c to 55
	// fmt.Println(c)

	/* ALGORITHMS */
	/* COPYING AN ARRAY using FOR loop	*/
	// moves a slice of integers from one array to another
	a := []int{2, 45, 43, 23, 67}              // a is an array of 5 ints expect [2 45 43 23 67]
	b := [11]int{}                             // b is an empty array of 11 ints expect [0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println("a: ", a, "b initialized:", b) // prints [2 45 43 23 67] [0 0 0 0 0 0 0 0 0 0 0]

	for i := 0; i < len(a); i++ { // for loop to copy array a to array c
		b[i] = a[i] // c[i] is equal to a[i]
		a[i] = 0    // a[i] is equal to 0 clearing a
	}
	fmt.Println("a emptied:", a, "b receiving a:", b) // prints [0 0 0 0 0] [2 45 43 23 67 0 0 0 0 0 0]

	/* ADDING ADDITIONAL NUMBERS TO EMPTY SLOTS IN AN ARRAY USING WHILE LOOP*/
	c := [6]int{14, 56, 67, 87, 90, 120} // c is an array of 6 ints expect [14 56 67 87 90 120]
	// adding all values of c to the next empty slot in b
	for i := 0; i < len(c); i++ { // for loop to set to 0 all elements of c time complexity O(n)
		// add d[i] to the next empty slot in c without writing over the initial values of c
		for j := 0; j < len(b); j++ { // if j less than len of c then add d[i] to c[j] time complexity is O(n)
			if b[j] == 0 { // if the slot is empty then add b[i] to the slot time complexity is O(1)
				b[j] = c[i] // b[j] = b[i]  time complexity is O(1)
				c[i] = 0    // clearing c[i] time complexity is O(1)
				break       // break out of the for loop time complexity is O(1)
			}
		}
	}
	// overall time complexity is O(n^2) space complexity is O(n)
	fmt.Println("b receiving c in next available slot:", b, "C emptied:", c) // prints [2 45 43 23 67 14 56 67 87 90 120] [0 0 0 0 0]
	fmt.Println("Binary Tree Creation and printing by breadth first traversal")
	/* CREATE A BINARY TREE AND TRAVERSE BREADTH FIRST	*/
	// create Treenode struct
	type Treenode struct {
		value int
		left  *Treenode
		right *Treenode
	}
	// populate Treenode struct
	root := &Treenode{value: 10}
	root.left = &Treenode{value: 5}
	root.right = &Treenode{value: 15}
	root.left.left = &Treenode{value: 1}
	root.left.right = &Treenode{value: 7}
	root.right.left = &Treenode{value: 12}

	// tree traversal printing values as it goes breadth first
	// prints 10 5 15 1 7 12
	var queue []*Treenode       // create a queue to hold the nodes
	queue = append(queue, root) // add the root node to the queue

	for len(queue) > 0 { // while the queue is not empty
		// pop the first node from the queue
		node := queue[0]              // node is equal to the first node in the queue
		queue = queue[1:]             // remove the first node from the queue
		fmt.Println(node.value)       // print the value of the node
		fmt.Println("queue: ", queue) // prints the memory address of the nodes
		// in the queue if no values [] is printed

		if node.left != nil { // if the node has a left node
			queue = append(queue, node.left) // add the left node to the queue
		}
		if node.right != nil { // if the node has a right node
			queue = append(queue, node.right) // add the right node to the queue
		}
	}

}

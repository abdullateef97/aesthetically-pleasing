package main

import "fmt"

/**
loop through the trees in the array A, at each index i, group in three trees [Ai, Ai+1, Ai+2]
check if the group of three trees are aesthetically pleasing (A.P)
i.e Ai > Ai+1 < A1+2 or Ai < Ai+1 > Ai+2
if not aesthetically pleasing, find ways to cut exactly one tree, from the group of three trees
can be one of the following possibilities
		cut Ai such that Ai+1 < Ai+2 > Ai+3 or Ai+1 > Ai+2 < Ai+3
		cut Ai+1 such that Ai < Ai+2 > Ai+3 or Ai > Ai+2 > Ai+3
		cut Ai+2 such that Ai < Ai+1 > Ai+3 or Ai > Ai+1 < Ai+3
then cut the tree from the list (delete from the array)
set A as the new array of trees

*/

// groupTreesIntoThrees returns three trees starting from index i
// returns an empty slice if i + 3 > len(slice)
func groupTreesIntoThrees(treeArray []int, index int) []int {
	if  len(treeArray) - (index + 3) < -1 {
		return []int{}
	}

	return treeArray[index: index + 3]
}

// checkIfAGroupOfTreesAreAestheticallyPleasing checks if a group of three trees are A.P
// returns true if 
// case 1:
// 			A0 > A1 < A2
// case 2: 
//			A0 < A1 > A2
// returns false if otherwise
func checkIfAGroupOfTreesAreAestheticallyPleasing(groupOfThreeTrees []int) bool {
	firstCase := groupOfThreeTrees[0] > groupOfThreeTrees[1] && groupOfThreeTrees[1] < groupOfThreeTrees[2]
	secondCase := groupOfThreeTrees[0] < groupOfThreeTrees[1] && groupOfThreeTrees[1] > groupOfThreeTrees[2]

	return firstCase || secondCase
}

func possibleWaysToCutTree(groupOfThreeTrees []int, treeAtIndexIPlus3 int) (int, int) {
	var ways int
	var indexOfTreeToCut int

	deleteTreeAndIncludeTreeAtIndexIPlus3 := func (slice []int, indexToDelete, itemToAppend int) []int {
		newGroup := make([]int, 0)
		newGroup = append(newGroup, slice[:indexToDelete]...)
		newGroup = append(newGroup, slice[indexToDelete+1:]...)
		newGroup = append(newGroup, itemToAppend)
		return newGroup
	}
	for index := range groupOfThreeTrees {
		modifiedGroupOfTrees := deleteTreeAndIncludeTreeAtIndexIPlus3(groupOfThreeTrees, index, treeAtIndexIPlus3)
		if checkIfAGroupOfTreesAreAestheticallyPleasing(modifiedGroupOfTrees) {
			ways = ways + 1
			indexOfTreeToCut = index
		}
	}
	return ways, indexOfTreeToCut
}

func AestheticallyPleasing (treeArray []int) int {
	var waysToCut int
	cantBeAesthetic := -1

	for i := 0; i < len(treeArray) - 2; i++ {
		groupOfThreeTrees := groupTreesIntoThrees(treeArray, i)
		fmt.Println("aaa", groupOfThreeTrees, treeArray)
		if len(groupOfThreeTrees) == 3 {
			if checkIfAGroupOfTreesAreAestheticallyPleasing(groupOfThreeTrees) {
				continue
			}
			// this means a tree in the treeArray has already been cut
			// since only one tree can be legally cut, return -1
			// as this would mean you would need to cut multiple trees
			if waysToCut > 0 {
				fmt.Println("fff", groupOfThreeTrees, treeArray)
				return cantBeAesthetic
			}

			possibleWaysToCut, indexOfTreeToCut := possibleWaysToCutTree(groupOfThreeTrees, treeArray[i + 3])
			fmt.Println("pp", possibleWaysToCut, i, groupOfThreeTrees)
			if possibleWaysToCut == 0 {
				return cantBeAesthetic
			}
			waysToCut += possibleWaysToCut
			treeArray = append(treeArray[:indexOfTreeToCut], treeArray[indexOfTreeToCut + 1:]...)
		}
	}
	return waysToCut
}


func main() {
	noOfWays := AestheticallyPleasing([]int{1,2,3,2,4}) // 2
	fmt.Println(noOfWays)
}
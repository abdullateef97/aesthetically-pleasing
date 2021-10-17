package main

import (
	"reflect"
	"testing"
)

func TestGroupTreesIntoThrees(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		treeArray []int
		index int
		expected []int
	}{
		{
			name: "when len(testArray) > index + 3",
			treeArray: []int{1,2,3,4,5},
			index: 1,
			expected: []int{2,3,4},
		},
		{
			name: "when (len(testArray) - index + 3) > -1",
			treeArray: []int{9,2,4,5,7},
			index: 4,
			expected: []int{},
		},
		{
			name: "when len(testArray) = index + 3",
			treeArray: []int{1,2,3,4,5,6},
			index: 3,
			expected: []int{4,5,6},
		},
	}

	for _, tt :=  range tests {
		result := groupTreesIntoThrees(tt.treeArray, tt.index)
		if !reflect.DeepEqual(result, tt.expected){
			t.Errorf("[ Test Case %s ], Expected: '%v' | Result '%v' ", tt.name, tt.expected, result)
		}
	}
}

func TestCheckIfAGroupOfTreesAreAestheticallyPleasing(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		group []int
		expected bool
	} {
		{
			name: "up-down-up",
			group: []int{3,2,4},
			expected: true,
		},
		{
			name: "down-up-down",
			group: []int{1,2,1},
			expected: true,
		},
		{
			name: "up-down-down",
			group: []int{7,3,2},
			expected: false,
		},
		{
			name: "down-up-up",
			group: []int{3,5,7},
			expected: false,
		},
		{
			name: "level-up",
			group: []int{3,3,5},
			expected: false,
		},
	}

	for _, tt := range tests {
		result := checkIfAGroupOfTreesAreAestheticallyPleasing(tt.group)
		if result != tt.expected {
			t.Errorf("[ Test Case %s ], Expected: '%v' | Result '%v' ", tt.name, tt.expected, result)
		}
	}
}

func TestPossibleWaysToCutTree(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		group []int
		treeAtIndexPlus3 int
		expectedWays int
		expectedIndex int
	} {
		{
			name: "down-up-down-up",
			group: []int{2,3,1},
			treeAtIndexPlus3: 4,
			expectedWays: 2,
			expectedIndex: 1,
		},
		{
			name: "down-up-down-up with three ways",
			group: []int{2,4,1},
			treeAtIndexPlus3: 2,
			expectedWays: 3,
			expectedIndex: 2,
		},
		{
			name: "down-up-level-down",
			group: []int{2,3,3},
			treeAtIndexPlus3: 1,
			expectedWays: 2,
			expectedIndex: 2,
		},
		{
			name: "down-up-down-level",
			group: []int{1,2,3},
			treeAtIndexPlus3: 3,
			expectedWays: 0,
			expectedIndex: 0,
		},
	}

	for _, tt := range tests {
		resultWays, resultIndex := possibleWaysToCutTree(tt.group, tt.treeAtIndexPlus3)
		if resultWays != tt.expectedWays || resultIndex != tt.expectedIndex {
			t.Errorf("[ Test Case %s ], Expected Ways: '%v' | Result Ways '%v' || Expected Index: '%v' | Result Index '%v'", tt.name, tt.expectedWays, resultWays, tt.expectedIndex, resultIndex)
		}
	}
}


func TestAestheticallyPleasing(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		treeArray []int
		expected int
	} {
		{
			name: "3 ways to cut a tree",
			treeArray: []int{3,4,5,3,7},
			expected: 3,
		},
		{
			name:  "no way to cut just a tree",
			treeArray: []int{1,2,3,4},
			expected: -1,
		},
		{
			name: "array of trees is already aesthetic",
			treeArray: []int{1,3,1,2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		result := AestheticallyPleasing(tt.treeArray)
		if result != tt.expected {
			t.Errorf("[ Test Case %s ], Expected : '%v' | Result : '%v'", tt.name, tt.expected, result)
		}
	}
}
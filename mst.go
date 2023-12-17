package main

import (
	"fmt"
	"math"
)

type MinimumSpanningTree struct {
	MSTDataset [][]float64
	Data       []MSTData
}

type MSTData struct {
	Iteration int
	Dataset   [][]float64
}

func RunMinimumSpanningTree() {

	var dataset [][]float64
	dataset = [][]float64{
		{1.1, 60},
		{8.2, 20},
		{4.2, 35},
		{1.5, 21},
		{7.6, 15},
		{2.0, 55},
		{3.9, 39},
	}

	var listIteration []MSTData
	listIteration = make([]MSTData, 0)

	mst := MinimumSpanningTree{
		MSTDataset: dataset,
		Data:       listIteration,
	}

	mst.AverageDistanceClustering()
}

type MatrixDistance struct {
	Name     []string
	Diameter []float64
}

func (m *MinimumSpanningTree) AverageDistanceClustering() {

	var centroids []MatrixDistance

	for i, a := range m.MSTDataset {
		// initiate matrix distance
		var matrixdistance MatrixDistance
		matrixdistance.Name = []string{fmt.Sprintf("%#v", i+1)}

		for j, b := range m.MSTDataset {
			if j == 0 {
				matrixdistance.Diameter = append(matrixdistance.Diameter, 0)
			} else if j > i {
				diameter := math.Sqrt(((a[0] - b[0]) * (a[0] - b[0])) + ((a[1] - b[1]) * (a[1] - b[1])))
				matrixdistance.Diameter = append(matrixdistance.Diameter, float64(int(diameter*100))/100)
			}
		}

		centroids = append(centroids, matrixdistance)
	}

	for i, value := range centroids {
		fmt.Printf("C")
		for _, name := range value.Name {
			fmt.Printf("%s", name)
		}
		fmt.Printf(" |")

		for c := 0; c < i; c++ {
			fmt.Printf("      |")
		}
		for _, d := range value.Diameter {
			fmt.Printf("%5.2f |", d)
		}
		fmt.Printf("\n")
	}

	var smallest float64
	smallest = 100
	var centroidname []string

	// get smallest centroid value
	for i, value := range centroids {
		for j, d := range value.Diameter {
			if d <= smallest && d != 0 {
				smallest = d
				centroidname = nil
				centroidname = append(centroidname, fmt.Sprintf("%#v", i+1))
				centroidname = append(centroidname, fmt.Sprintf("%#v", j+1+i))
			}
		}
	}

	fmt.Printf("C%#v%#v. %#v\n", centroidname[0], centroidname[1], smallest)

	// m.CalculateNewTable(centroidname[0], centroidname[1])
	m.CalculateNewTable(2, 6)

}

func (m *MinimumSpanningTree) CalculateNewTable(c1index int, c2index int) {

	c1 := m.MSTDataset[c1index]
	c2 := m.MSTDataset[c2index]

	newc1 := (c1[0] + c1[1]) / 2
	newc2 := (c2[0] + c2[1]) / 2

	newc := []float64{newc1, newc2}

	// Indices to remove
	indicesToRemove := []int{2}

	// Remove elements at specified indices
	for _, index := range indicesToRemove {
		m.MSTDataset = removeAtIndex(m.MSTDataset, index)
	}

	// Insert new data at index 3
	m.MSTDataset = insertAtIndex(m.MSTDataset, 3, newc)

	var centroids []MatrixDistance

	for i, a := range m.MSTDataset {
		// initiate matrix distance
		var matrixdistance MatrixDistance
		matrixdistance.Name = []string{fmt.Sprintf("%#v", i+1)}

		for j, b := range m.MSTDataset {
			if j == 0 {
				matrixdistance.Diameter = append(matrixdistance.Diameter, 0)
			} else if j > i {
				diameter := math.Sqrt(((a[0] - b[0]) * (a[0] - b[0])) + ((a[1] - b[1]) * (a[1] - b[1])))
				matrixdistance.Diameter = append(matrixdistance.Diameter, float64(int(diameter*100))/100)
			}
		}

		centroids = append(centroids, matrixdistance)
	}

	for i, value := range centroids {
		fmt.Printf("C")
		for _, name := range value.Name {
			fmt.Printf("%s", name)
		}
		fmt.Printf(" |")

		for c := 0; c < i; c++ {
			fmt.Printf("      |")
		}
		for _, d := range value.Diameter {
			fmt.Printf("%5.2f |", d)
		}
		fmt.Printf("\n")
	}

}

// removeAtIndex removes the element at the specified index from a 2D slice.
func removeAtIndex(slice [][]float64, index int) [][]float64 {
	return append(slice[:index], slice[index+1:]...)
}

// insertAtIndex inserts new data into a 2D slice at the specified index.
func insertAtIndex(slice [][]float64, index int, newData []float64) [][]float64 {
	slice = append(slice[:index], append([][]float64{newData}, slice[index:]...)...)
	return slice
}

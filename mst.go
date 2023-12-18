package main

import (
	"fmt"
	"math"
)

type MinimumSpanningTree struct {
	MSTDataset          [][]float64
	MatrixCentroids     []MatrixDistance
	Iteration           int
	NewCentroid         []float64
	NewNearestCentroid  []float64
	NewFarthestCentroid []float64
	NewMSTDataset       [][]float64
}

func RunMinimumSpanningTree() {

	var (
		mstlist []MinimumSpanningTree
		// dataset [][]float64
		mst MinimumSpanningTree

		matrixDataset []MatrixDistance
	)
	mstlist = make([]MinimumSpanningTree, 0)

	/*
		// iterasi pertama
		dataset = [][]float64{
			{1.1, 60},
			{8.2, 20},
			{4.2, 35},
			{1.5, 21},
			{7.6, 15},
			{2.0, 55},
			{3.9, 39},
		}
		mst = MinimumSpanningTree{
			MSTDataset: dataset,
		}
		mstlist = append(mstlist, mst)

		// iterasi kedua
		dataset = [][]float64{
			{1.1, 60},
			{8.2, 20},
			{4.2, 35},
			{1.5, 21},
			{7.6, 15},
			{2, 55},
		}

		mst = MinimumSpanningTree{
			MSTDataset: dataset,
		}
		mstlist = append(mstlist, mst)

		dataset = [][]float64{
			{1.1, 60},
			{8.2, 20},
			{3.9, 39},
			{1.5, 21},
			{7.6, 15},
			{2, 55},
		}

		mst = MinimumSpanningTree{
			MSTDataset: dataset,
		}
		mstlist = append(mstlist, mst)

		for _, mstdata := range mstlist {
			// mstdata.AverageDistanceFromMatrix()
			mstdata.AverageDistanceClustering()
		}

		return
	*/
	mstlist = nil

	matrixDataset = make([]MatrixDistance, 0)
	matrixDataset = append(matrixDataset, MatrixDistance{Name: []string{"1"}, Diameter: []float64{0, 11, 8, 4, 5, 10}})
	matrixDataset = append(matrixDataset, MatrixDistance{Name: []string{"2"}, Diameter: []float64{0, 6, 14, 7, 3}})
	matrixDataset = append(matrixDataset, MatrixDistance{Name: []string{"3"}, Diameter: []float64{0, 9, 2, 12}})
	matrixDataset = append(matrixDataset, MatrixDistance{Name: []string{"4"}, Diameter: []float64{0, 15, 1}})
	matrixDataset = append(matrixDataset, MatrixDistance{Name: []string{"5"}, Diameter: []float64{0, 6}})
	matrixDataset = append(matrixDataset, MatrixDistance{Name: []string{"6"}, Diameter: []float64{0}})

	mst = MinimumSpanningTree{
		MatrixCentroids: matrixDataset,
	}
	mstlist = append(mstlist, mst)

	for _, mstdata := range mstlist {
		mstdata.AverageDistanceFromMatrix()
		// mstdata.AverageDistanceClustering()
	}
}

type MatrixDistance struct {
	Name     []string
	Diameter []float64
}

func (m *MinimumSpanningTree) AverageDistanceFromMatrix() {

	for i, value := range m.MatrixCentroids {
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

	var (
		smallest         float64
		centroidname     []string
		centroidselected []int
	)

	smallest = 100

	// get smallest centroid value
	for i, value := range m.MatrixCentroids {
		// fmt.Printf("%#v. length: %#v\n", i, len(value.Diameter))
		for j, d := range value.Diameter {
			if d <= smallest && d != 0 {
				smallest = d
				centroidname = nil
				// fmt.Printf("smallest selected: %#v | %#v | %#v \n", smallest, i+1, j+1+i)
				centroidname = append(centroidname, fmt.Sprintf("%#v", i+1))
				centroidname = append(centroidname, fmt.Sprintf("%#v", j+1+i))

				centroidselected = nil
				centroidselected = append(centroidselected, i+1)
				centroidselected = append(centroidselected, j+1+i)
			}
		}
	}

	var (
		centroida []float64
		centroidb []float64

		newcentroida float64
		newcentroidb float64

		newnearestcentroid  []float64
		newfarthestcentroid []float64
	)

	// centroida = m.MatrixCentroids[centroidselected[0]]

	centroidb = m.MSTDataset[centroidselected[1]-1]

	newcentroida = (centroida[0] + centroidb[0]) / 2
	newcentroidb = (centroida[1] + centroidb[1]) / 2

	if newcentroida < newcentroidb {
		newnearestcentroid = []float64{centroida[0], centroida[1]}
	} else {
		newnearestcentroid = []float64{centroidb[0], centroidb[1]}
	}

	if newcentroida > newcentroidb {
		newfarthestcentroid = []float64{centroida[0], centroida[1]}
	} else {
		newfarthestcentroid = []float64{centroidb[0], centroidb[1]}
	}

	m.NewCentroid = []float64{newcentroida, newcentroidb}
	m.NewNearestCentroid = newnearestcentroid
	m.NewFarthestCentroid = newfarthestcentroid

	m.Iteration++

	if m.Iteration < len(m.MSTDataset) {

		/* make new dataset table average distance */
		var (
			smallest_centroid_index int
			newMSTDataset           [][]float64
			newMSTNearestDataset    [][]float64
			newMSTFarthestDataset   [][]float64
		)
		if centroidselected[0] < centroidselected[1] {
			smallest_centroid_index = centroidselected[0] - 1
		} else {
			smallest_centroid_index = centroidselected[1] - 1
		}

		for i, ds := range m.MSTDataset {
			if i == smallest_centroid_index {
				// append new
				newMSTDataset = append(newMSTDataset, m.NewCentroid)
				newMSTNearestDataset = append(newMSTNearestDataset, m.NewNearestCentroid)
				newMSTFarthestDataset = append(newMSTFarthestDataset, m.NewFarthestCentroid)
			} else if i != centroidselected[0]-1 && i != centroidselected[1]-1 {
				newMSTDataset = append(newMSTDataset, ds)
				newMSTNearestDataset = append(newMSTNearestDataset, ds)
				newMSTFarthestDataset = append(newMSTFarthestDataset, ds)
			}
		}

		fmt.Println("dataset = [][]float64{")
		for _, ds := range newMSTDataset {
			fmt.Printf("\t{%#v, %#v},\n", ds[0], ds[1])
		}
		fmt.Println("}")

		fmt.Println("dataset = [][]float64{")
		for _, ds := range newMSTNearestDataset {
			fmt.Printf("\t{%#v, %#v},\n", ds[0], ds[1])
		}
		fmt.Println("}")

		fmt.Println("dataset = [][]float64{")
		for _, ds := range newMSTFarthestDataset {
			fmt.Printf("\t{%#v, %#v},\n", ds[0], ds[1])
		}
		fmt.Println("}")
	}

}
func (m *MinimumSpanningTree) AverageDistanceClustering() {

	var (
		centroids []MatrixDistance
	)

	for i, a := range m.MSTDataset {
		// initiate matrix distance
		var matrixdistance MatrixDistance
		matrixdistance.Name = []string{fmt.Sprintf("%#v", i+1)}
		fmt.Println(matrixdistance.Name)

		for j, b := range m.MSTDataset {
			if j == 0 {
				matrixdistance.Diameter = append(matrixdistance.Diameter, 0)
			} else if j > i {
				// euclidean distance
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

	var (
		smallest         float64
		centroidname     []string
		centroidselected []int
	)

	smallest = 100

	// get smallest centroid value
	for i, value := range centroids {
		// fmt.Printf("%#v. length: %#v\n", i, len(value.Diameter))
		for j, d := range value.Diameter {
			if d <= smallest && d != 0 {
				smallest = d
				centroidname = nil
				// fmt.Printf("smallest selected: %#v | %#v | %#v \n", smallest, i+1, j+1+i)
				centroidname = append(centroidname, fmt.Sprintf("%#v", i+1))
				centroidname = append(centroidname, fmt.Sprintf("%#v", j+1+i))

				centroidselected = nil
				centroidselected = append(centroidselected, i+1)
				centroidselected = append(centroidselected, j+1+i)
			}
		}
	}

	fmt.Printf("selected centroid : %#v | %#v\n", centroidselected[0], centroidselected[1])

	var (
		centroida []float64
		centroidb []float64

		newcentroida float64
		newcentroidb float64

		newnearestcentroid  []float64
		newfarthestcentroid []float64
	)

	centroida = m.MSTDataset[centroidselected[0]-1]
	centroidb = m.MSTDataset[centroidselected[1]-1]

	newcentroida = (centroida[0] + centroidb[0]) / 2
	newcentroidb = (centroida[1] + centroidb[1]) / 2

	if newcentroida < newcentroidb {
		newnearestcentroid = []float64{centroida[0], centroida[1]}
	} else {
		newnearestcentroid = []float64{centroidb[0], centroidb[1]}
	}

	if newcentroida > newcentroidb {
		newfarthestcentroid = []float64{centroida[0], centroida[1]}
	} else {
		newfarthestcentroid = []float64{centroidb[0], centroidb[1]}
	}

	m.NewCentroid = []float64{newcentroida, newcentroidb}
	m.NewNearestCentroid = newnearestcentroid
	m.NewFarthestCentroid = newfarthestcentroid

	m.Iteration++

	if m.Iteration < len(m.MSTDataset) {

		/* make new dataset table average distance */
		var (
			smallest_centroid_index int
			newMSTDataset           [][]float64
			newMSTNearestDataset    [][]float64
			newMSTFarthestDataset   [][]float64
		)
		if centroidselected[0] < centroidselected[1] {
			smallest_centroid_index = centroidselected[0] - 1
		} else {
			smallest_centroid_index = centroidselected[1] - 1
		}

		for i, ds := range m.MSTDataset {
			if i == smallest_centroid_index {
				// append new
				newMSTDataset = append(newMSTDataset, m.NewCentroid)
				newMSTNearestDataset = append(newMSTNearestDataset, m.NewNearestCentroid)
				newMSTFarthestDataset = append(newMSTFarthestDataset, m.NewFarthestCentroid)
			} else if i != centroidselected[0]-1 && i != centroidselected[1]-1 {
				newMSTDataset = append(newMSTDataset, ds)
				newMSTNearestDataset = append(newMSTNearestDataset, ds)
				newMSTFarthestDataset = append(newMSTFarthestDataset, ds)
			}
		}

		fmt.Println("dataset = [][]float64{")
		for _, ds := range newMSTDataset {
			fmt.Printf("\t{%#v, %#v},\n", ds[0], ds[1])
		}
		fmt.Println("}")

		fmt.Println("dataset = [][]float64{")
		for _, ds := range newMSTNearestDataset {
			fmt.Printf("\t{%#v, %#v},\n", ds[0], ds[1])
		}
		fmt.Println("}")

		fmt.Println("dataset = [][]float64{")
		for _, ds := range newMSTFarthestDataset {
			fmt.Printf("\t{%#v, %#v},\n", ds[0], ds[1])
		}
		fmt.Println("}")
	}

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

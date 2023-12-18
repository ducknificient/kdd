package main

import (
	"fmt"
	"math"
)

type KNN struct {
	Diameter float64
	selected []float64
}

func RunKNearestNeighbor() {

	var knnLIst [][]float64
	var smallest float64
	smallest = 500

	ip := []float64{105, 18.9}

	knn := KNN{
		selected: []float64{},
		Diameter: smallest,
	}

	knnLIst = populateExampleDataset()

	for _, ds := range knnLIst {

		diameter := math.Sqrt((ds[0]-ip[0])*(ds[0]-ip[0]) + ((ds[1] - ip[1]) * (ds[1] - ip[1])))
		fmt.Printf("knn: %#v diameter: %#v\n", knn.Diameter, diameter)
		if diameter < knn.Diameter {

			// check terdekat sebesar N
			if diameter != 10 {
				fmt.Printf("diameter: %#v\n", diameter)
				knn.Diameter = diameter
				knn.selected = []float64{ds[0], ds[1]}
			}
		}
	}

	fmt.Printf("smallest : %#v | %#v\n", knn.Diameter, knn.selected)

}

func populateExampleDataset() (list [][]float64) {

	list = append(list, []float64{60, 18.4})
	list = append(list, []float64{85.5, 16.8})
	list = append(list, []float64{64.8, 21.6})
	list = append(list, []float64{61.5, 20.8})
	list = append(list, []float64{87, 23.6})

	list = append(list, []float64{110.1, 19.2})
	list = append(list, []float64{108, 17.6})
	list = append(list, []float64{82.8, 22.4})
	list = append(list, []float64{69, 20})
	list = append(list, []float64{93, 20.8})

	list = append(list, []float64{51, 22})
	list = append(list, []float64{81, 20})
	list = append(list, []float64{75, 19.6})
	list = append(list, []float64{52.8, 20.8})
	list = append(list, []float64{64.8, 27.2})

	list = append(list, []float64{43.2, 20.4})
	list = append(list, []float64{84, 17.6})
	list = append(list, []float64{49.2, 17.6})
	list = append(list, []float64{59.4, 16})
	list = append(list, []float64{66, 18.4})

	list = append(list, []float64{47.4, 16.4})
	list = append(list, []float64{33, 18.8})
	list = append(list, []float64{51, 14})
	list = append(list, []float64{663, 14.8})

	return list
}

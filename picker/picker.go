package picker

import (
	"math"
	"strconv"
	"strings"
)

type Picker struct {
}

type Answer struct {
	Can    bool
	Count  int
	Nums   []int
	Status string
}

func NewPicker() *Picker {
	return &Picker{}
}

func (p *Picker) Calculate(set []int, needCnt int) Answer {

	sum := needCnt
	n := len(set)

	YSize := sum + 1
	XSize := n + 1

	can := make([][]bool, YSize)
	cany := make([]bool, XSize*YSize)
	for i := range can {
		can[i], cany = cany[:XSize], cany[XSize:]
	}

	count := make([][]int, YSize)
	county := make([]int, XSize*YSize)
	for i := range can {
		count[i], county = county[:XSize], county[XSize:]
	}

	num := make([][]string, YSize)
	numy := make([]string, XSize*YSize)
	for i := range can {
		num[i], numy = numy[:XSize], numy[XSize:]
	}

	for i := 0; i <= n; i++ {
		can[0][i] = true
		count[0][i] = 0
		num[0][i] = ""
	}

	for i := 1; i <= sum; i++ {
		can[i][0] = false
		count[i][0] = -1
		num[i][0] = ""
	}

	for i := 1; i <= sum; i++ {
		for j := 1; j <= n; j++ {
			can[i][j] = can[i][j-1]
			count[i][j] = count[i][j-1]
			num[i][j] = num[i][j-1]

			if i >= set[j-1] {
				can[i][j] = can[i][j] || can[i-set[j-1]][j-1]

				if can[i][j] {
					count[i][j] = int(math.Max(float64(count[i][j-1]), float64(count[i-set[j-1]][j-1]+1)))

					if count[i][j] > count[i][j-1] {
						num[i][j] = num[i-set[j-1]][j-1] + "," + strconv.Itoa(j-1)
					}
				}
			}
		}
	}

	for index := sum; 0 <= index; index-- {

		if can[index][n] {

			str := strings.Split(num[index][n], ",")
			nums := []int{}
			for n := 0; n < len(str); n++ {
				if str[n] == "" {
					continue
				}
				cn, _ := strconv.Atoi(str[n])
				nums = append(nums, cn)
			}
			return Answer{
				Can:    can[index][n] && count[index][n] > 0,
				Count:  count[index][n],
				Nums:   nums,
				Status: "OK",
			}

		}
	}
	return Answer{Can: false}
}

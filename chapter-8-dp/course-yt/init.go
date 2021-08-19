package course_yt

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	shape [2]int

	matrix [][]int

	redRoomsCount int
	redRoomsPos   [][2]int
)


const RedRoomColor = "\033[31;1;4m%d\033[0m\t"

func init() {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	fmt.Fscanf(f, "%d %d", &shape[0], &shape[1])

	matrix = make([][]int, shape[1])

	var s string
	for i := 0; i < shape[1]; i++ {
		for j := 0; j < shape[0]; j++ {
			fmt.Fscan(f, &s)

			for _, el := range strings.Split(s, " ") {
				num, _ := strconv.Atoi(el)
				matrix[i] = append(matrix[i], num)
			}
		}
	}

	fmt.Fscanf(f, "%d", &redRoomsCount)

	var x, y int
	for i := 0; i < redRoomsCount; i++ {
		fmt.Fscanf(f, "%d %d", &x, &y)
		redRoomsPos = append(redRoomsPos, [2]int{x, y})
	}

	fmt.Println("matrix shape:", shape)
	fmt.Println("red rooms count:", redRoomsCount)
	fmt.Println("red rooms coordinates:", redRoomsPos)

	fmt.Println("matrix:")

	for i := range matrix {
		for j := range matrix[i] {
			red := false
			for k := range redRoomsPos {
				if redRoomsPos[k][0] == i && redRoomsPos[k][1] == j {
					red = true
					break
				}
			}

			if red {
				fmt.Printf(RedRoomColor, matrix[i][j])
				continue
			}

			fmt.Printf("%d\t", matrix[i][j])
		}
		println()
	}

	println()
}
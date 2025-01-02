package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int]

type Robot struct {
	startPos Vec2
	velocity Vec2
}

func ParseRobot(line string) Robot {
	fields := strings.Fields(line)
	// strip off the initial "p=" or "v=" from the two fields, then convert to vec2
	return Robot{util.NewVec2Int(fields[0][2:]), util.NewVec2Int(fields[1][2:])}
}

func (r Robot) FinalQuadrant(seconds int) ([2]bool, error) {
	fp := r.FinalPos(seconds)
	x, y := fp.Parts()
	if x == width/2 || y == height/2 {
		return [2]bool{}, fmt.Errorf("final position was in center row or column")
	}
	return [2]bool{x > width/2, y > height/2}, nil
}

func (r Robot) FinalPos(seconds int) Vec2 {
	unwrappedPos := r.startPos.Add(r.velocity.Mul(seconds))
	finalPos := Vec2{mod(unwrappedPos[0], width), mod(unwrappedPos[1], height)}
	//fmt.Printf("final pos for robot with startPos %v, vel %v = %v\n", r.startPos, r.velocity, finalPos)
	return finalPos
}

func mod(a int, b int) int {
	return ((a % b) + b) % b
}

type PuzzleInput struct {
	robots []Robot
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		problem.robots = append(problem.robots, ParseRobot(line))
	}
	return problem
}

type PuzzleImage struct {
	p          PuzzleInput
	seconds    int
	imageCache map[int][][]bool
}

func (pi *PuzzleImage) ColorModel() color.Model {
	return color.GrayModel
}

func (pi *PuzzleImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, width, height)
}

func (pi *PuzzleImage) At(x, y int) color.Color {
	pi.cacheImage()
	if pi.imageCache[pi.seconds][x][y] {
		return color.Gray{255}
	} else {
		return color.Gray{0}
	}
}

func (pi *PuzzleImage) WriteImageToFile() {
	filename := fmt.Sprintf("out/image_%04d.png", pi.seconds)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln("failed to open image for writing", err)
	}

	if err := png.Encode(f, pi); err != nil {
		f.Close()
		log.Fatalln("failed to encode image", err)
	}

	if err := f.Close(); err != nil {
		log.Fatalln("failed to close image", err)
	}

	fmt.Println("wrote image to", filename)
}

func (pi *PuzzleImage) cacheImage() {
	if _, ok := pi.imageCache[pi.seconds]; ok {
		return
	}

	pi.imageCache[pi.seconds] = make([][]bool, width)
	for r := range width {
		pi.imageCache[pi.seconds][r] = make([]bool, height)
	}

	for _, r := range pi.p.robots {
		x, y := r.FinalPos(pi.seconds).Parts()
		pi.imageCache[pi.seconds][x][y] = true
	}
}

func solve(p PuzzleInput) int64 {
	quadrantCounts := make(map[[2]bool]int)
	for _, r := range p.robots {
		quad, err := r.FinalQuadrant(100)
		if err == nil {
			quadrantCounts[quad]++
		}
	}

	ans := 1
	for _, count := range quadrantCounts {
		ans *= count
	}
	return int64(ans)
}

func writeAllImages(p PuzzleInput) {
	image := PuzzleImage{
		p:          p,
		imageCache: make(map[int][][]bool),
	}
	for i := range max(width, height) {
		image.seconds = i
		image.WriteImageToFile()
	}

	image.seconds = t
	image.WriteImageToFile()
}

//go:embed input
var input string

// const width, height = 11, 7
const width, height = 101, 103
const t = 7861

// results from investigating part 2 images:
// image_0033 has dots clustered in rows
// image_0084 has dots clustered in columns
// using insight from reddit, find the desired time t such that
// t = 33 mod 103
// t = 84 mod 101
// CRT calculator shows t = 7861

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input)
		if isPart2 {
			writeAllImages(problem)
			return 0
		} else {
			return solve(problem)
		}
	})
}

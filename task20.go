package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
p=<-3787,-3683,3352>, v=<41,-25,-124>, a=<5,9,1>
p=<6815,2269,3786>, v=<-93,23,38>, a=<-8,-6,-10>
p=<-1586,-5016,3166>, v=<2,66,-70>, a=<3,6,-2>
p=<-2392,-397,3538>, v=<76,-19,-114>, a=<0,2,0>


*/
type Vector3D struct {
	x, y, z int
}

type GPU_Particle struct {
	Position     *Vector3D
	Velocity     *Vector3D
	Acceleration *Vector3D
}

func (p *GPU_Particle) Move() {
	//New Position
	p.Position.x += p.Velocity.x
	p.Position.y += p.Velocity.y
	p.Position.z += p.Velocity.z
	//New Velocity
	p.Velocity.x += p.Acceleration.x
	p.Velocity.y += p.Acceleration.y
	p.Velocity.z += p.Acceleration.z
}

func NewGPU_Particle(input string) *GPU_Particle {
	p := &GPU_Particle{}

	trim := strings.Trim(input, "pva=<>")
	replaced := strings.Replace(trim, ",", " ", -1)
	splits := helpers.SplitBySpace(replaced)

	pos := &Vector3D{}
	pos.x, _ = strconv.Atoi(splits[0])
	pos.y, _ = strconv.Atoi(splits[1])
	pos.z, _ = strconv.Atoi(splits[2])

	vel := &Vector3D{}
	vel.x, _ = strconv.Atoi(splits[3])
	vel.y, _ = strconv.Atoi(splits[4])
	vel.z, _ = strconv.Atoi(splits[5])

	acc := &Vector3D{}
	acc.x, _ = strconv.Atoi(splits[6])
	acc.y, _ = strconv.Atoi(splits[7])
	acc.z, _ = strconv.Atoi(splits[8])

	p.Acceleration = acc
	p.Position = pos
	p.Velocity = vel

	return p
}

//Task20 Solution
func Task20() {
	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input20.txt")
	fmt.Printf("[Part 1 answer] - %v\n", task20PartOne(file))
	// fmt.Printf("[Part 2 answer] - %v\n", task18PartTwo(instructions))
}

// One with lowest acceleration will be closest to the centre.
func task20PartOne(input []string) int {
	closestAcc := 1000000.0
	closestID := 0

	for id, line := range input {
		a := helpers.SplitBySpace(line)
		b := strings.Trim(a[2], "a=<>")
		c := strings.Split(b, ",")
		sum := 0
		for _, x := range c {
			v, _ := strconv.Atoi(x)
			sum += v * v
		}

		root := math.Sqrt(float64(sum))
		fmt.Println(root)
		if root < closestAcc {
			closestAcc = root
			closestID = id
		}
	}
	return closestID
}

/*
Very simply, if you define the minimum distance between two moving bounding volumes as a function of time, D(t), then the roots (a.k.a. zeros) of that function are obviously the set of times, {t1, t2, ... tn}, when the distance between the volumes is exactly zero, D(t1)=0, D(t2)=0, ... D(tn)=0. If the volumes collide at some time in past or future, then one or more of those roots will be real. The smallest real root that is greater than or equal to the current time can usually be described as the time of the initial collision between the two volumes. If the volumes never collide in the past or future, then D(t) will have no real roots, because it is never equal to zero anywhere at any time, in which case it will only have imaginary roots. So, simply by examining the presence or lack of real roots in the distance function, you can determine if a collision occurs or not.

Polynomial functions up to degree 3 (i.e. "cubic"; ax3+bx2+cx+d) have well-defined roots that can be derived algebraically very easily. Polynomials of degree 4 (a.k.a. quartic) also have a relatively straight forward solution for their roots which was provided by Lodovico Ferrari, so it is sometimes called Ferrari's Method. For functions that are more complex than that, you can use Newton's method to systematically find roots between some tmin and tmax. Newton's method actually is somewhat ideal for collision detection, since we are usually only concerned with finding collisions between some known time interval. There's almost never any need to hunt down all of the roots for this application. Furthermore, the precision that we require for collision events often isn't very high. The precision of the root being computed can be tuned in Newton's method by simply increasing or decreasing the number of iterations spent in its loop.

https://en.wikipedia.org/wiki/Collision_detection
*/

func task20PartTwo(input []string) int {
	// for _, v := range input {
	// 	//p := NewGPU_Particle(v)
	// }

	return 0
}

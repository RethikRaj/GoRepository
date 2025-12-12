package constants

import (
	"fmt"
	"math"
)

const PackageLevelConstant = "constant"

func Constant1() {
	fmt.Println(PackageLevelConstant)

	const n  = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
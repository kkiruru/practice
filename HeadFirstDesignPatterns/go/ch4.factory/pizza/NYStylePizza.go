package pizza

import (
	"fmt"
)

func init() {
	fmt.Println("pizza package의 NYUStyleCheesePizza 파일의 init 함수")
}


type NYStyleCheesePizza struct {
	parent
}

func (n NYStyleCheesePizza) Prepare() {
	fmt.Println("Prepare " + n.Name)
}


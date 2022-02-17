package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	hello := emoji.Sprint("Hello world :world_map:")
	fmt.Println(hello)
}

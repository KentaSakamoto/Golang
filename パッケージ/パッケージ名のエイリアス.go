import (
	"sync"

	mysync "github.com/tenntenn/sync"
	greeting "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do())
}
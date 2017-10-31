package main
import "testing"
import "time"
import "fmt"

func TestSomething(t *testing.T) {
	fmt.Println(logo)
        fmt.Println("sleep 21 seconds")
	for i := 0; i < 21; i++ {
		time.Sleep(time.Second)
		fmt.Printf("slept %v seconds\n", i)
	}
}

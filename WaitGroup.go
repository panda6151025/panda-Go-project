package main
import (
	"fmt"
	"sync"
	"time"
)

func woker(id int){
	fmt.Printf("Woker %d starting\n",id)
	time.Sleep(time.Second)
	fmt.Printf("Woker %d done\n",id)
}
func main(){
	var wg sync.WaitGroup

	for i:=1;i<=5;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			woker(i)
		}()
	}
	wg.Wait()
}2=/
//=/=
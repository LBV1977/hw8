package main

import (
	//"context"
	"context"
	"fmt"
	"sync"
	"time"
	//"time"
)

/************************************************************************************************/
func rut_RxCh(ch chan int, wg *sync.WaitGroup) {	
	defer wg.Done()
	defer fmt.Println("stop rut_RxCh")

	fmt.Println("start rut_RxCh")
	//wg.Add(1)	
	for v :=range ch {
		fmt.Println(v)
	}	
}

func rut_TxCh(ctx context.Context, ch chan int,  wg *sync.WaitGroup) {
	defer close(ch) 
		defer wg.Done()	
			defer fmt.Println("stop rut_TxCh")

	fmt.Println("start rut_TxCh")
	//wg.Add(1)

	 for i := 0; i < 25500; i++ {
		select{
		case <-ctx.Done():
			return;
		default:
			ch <- i
		}		
	 }
}
/************************************************************************************************/
func rut_RxChString(ch chan string, wg *sync.WaitGroup) {	
	defer wg.Done()
	defer fmt.Println("stop rut_RxChString")

	fmt.Println("start rut_RxChString")
	for s :=range ch {
		fmt.Println(s)
	}	
}

func rut_TxChString(ctx context.Context, ch chan string,  wg *sync.WaitGroup) {
	defer close(ch) 
		defer wg.Done()	
			defer fmt.Println("stop rut_TxChString")

 lst := []string{
	"Sddffbhhshf1",
	"Sddffbhhshf2",
	"Sddffbhhshf3",
	"Sddffbhhshf4",
	"Sddffbhhshf5",
	"Sddffbhhshf6",
	"Sddffbhhshf7",
	"Sddffbhhshf8",
	}


	fmt.Println("start rut_TxChString")

	 for idx:=range lst {
		select{
		case <-ctx.Done():
			return;
		default:
			ch <-lst[idx]
		}		
	 }
}
/************************************************************************************************/

func main() {
	ctx, _ :=context.WithTimeout(context.Background(), time.Second*1)
	fmt.Println("start")
	var wg sync.WaitGroup
	wg.Add(2)

	ch:=make(chan int,5)
	go rut_RxCh(ch,&wg)
	go rut_TxCh(ctx, ch,&wg)

	// ch:=make(chan string,5)
	// go rut_RxChString(ch,&wg)
	// go rut_TxChString(ctx, ch,&wg)

	
	wg.Wait()
	fmt.Println("stop")

}
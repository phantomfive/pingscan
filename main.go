package main
import ("fmt";"sync/atomic";"time";"os/exec")


var runningThreads int32 = 0
var MAX_THREADS int32 = 300

func main() {
	var a, b, c, d int =     10, 10,   0,   0
	var a2, b2, c2, d2 int = 10, 10, 255, 255

	for ; a <= a2; a++ {
		for ; b <= b2; b++ {
			for ; c <= c2; c++ {
				for ; d <= d2; d++ {
					time.Sleep(10 * time.Millisecond)
					for ;runningThreads>MAX_THREADS; {
						time.Sleep(100 * time.Millisecond)
					}
					atomic.AddInt32(&runningThreads, 1)
					go probePort(a, b, c, d)
				}
				d = 0
			}
			c = 0
		}
		b = 0
	}

	for ; runningThreads > 0 ; {
	}
		
}

func probePort(a int, b int, c int, d int) {

	arg := fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
	fmt.Print(".")
	cmd :=exec.Command("ping", "-c", "1", "-W", "45",  arg)
	_, e := cmd.Output()
	if e==nil {
		fmt.Printf("\nFound ip: %d.%d.%d.%d\n", a, b, c, d)
	}/*else {
		fmt.Printf("err: %s    %s", out, e)
	}*/
	
	atomic.AddInt32(&runningThreads, -1)
}







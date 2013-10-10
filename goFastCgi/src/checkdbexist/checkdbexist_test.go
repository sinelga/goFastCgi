package checkdbexist

import (
	"log"
	"testing"
	"time"
)

func Test_Checkdb(t *testing.T) {

	log.Println("Start test")

	t0 := time.Unix(1381292732, 0)
//
	t1 := time.Now()
	
	log.Println(t1.Sub(t0))
	log.Println(int(time.Since(t0).Hours()))

}

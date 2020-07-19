package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"sync"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

var less string = "abcdefghijklmnopqrstuvwxyz"
var total int = 26

func main() {
	wg := new(sync.WaitGroup)
	p := mpb.New(mpb.WithWaitGroup(wg))
	wg.Add(total)
	for i := 0; i < total; i++ {
		name := fmt.Sprintf("%c :", less[i])
		bar := p.AddBar(int64(11881376),
			mpb.PrependDecorators(decor.Name(name), decor.CountersNoUnit("%d/%d", decor.WCSyncSpace)),
			mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5})))
		go task(wg, bar, i)
	}
	p.Wait()
}

func task(wg *sync.WaitGroup, bar *mpb.Bar, startIndex int) {
	defer wg.Done()
	passwd := make([]byte, 6)
	passwd[0] = less[startIndex]
	for i := 0; i < total; i++ {
		passwd[1] = less[i]
		for i := 0; i < total; i++ {
			passwd[2] = less[i]
			for i := 0; i < total; i++ {
				passwd[3] = less[i]
				for i := 0; i < total; i++ {
					passwd[4] = less[i]
					for i := 0; i < total; i++ {
						passwd[5] = less[i]
						count := 0
						indexSum := 0
						md5Value := md5.Sum(passwd)
						md5String := fmt.Sprintf("%x", md5Value)
						for index := 0; index < 32; index++ {
							if md5String[index] == '0' {
								count++
								indexSum += index
							}
						}
						if count*10+indexSum == 403 {
							fmt.Printf("%s", passwd)
							os.Exit(0)
						}
						bar.Increment()
					}
				}
			}
		}
	}
}

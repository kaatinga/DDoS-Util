package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"workerpool/worker"
	//"./worker"
)

const numOfWorkers = 4

func main() {

	var err error

	// меняем вывод для log
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	// специально для ошибок будет Ahtung
	Ahtung := log.New(os.Stderr, "ERROR: ", log.LstdFlags)

	var dDoSURLs worker.URLs

	err = dDoSURLs.AddURL("yandex", "http://yandex.ru", ``)
	if err != nil {
		Ahtung.Println(err)
	}

	err = dDoSURLs.AddURL("google", "http://google.ru", ``)
	if err != nil {
		Ahtung.Println(err)
	}

	jobCh := make(chan worker.URL)

	go func() {
		for {
			jobCh <- dDoSURLs.GetURL()
				time.Sleep(time.Millisecond * 50)
		}
	}()

	for i := 0; i < numOfWorkers; i++ {
		w := worker.NewWorker(i, jobCh)
		go w.HandleJobs()
	}

	log.Println("The DDoS util was launched!")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Println("The DDoS util was terminated!")
}

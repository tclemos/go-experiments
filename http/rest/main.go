package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type payload struct {
	a int
	b int
}

func main() {

	configureLog()
	serverAddress := serve()

	cycles := 10
	n := 100000
	log.Info().Msgf("Starting %d cycles of %d request", cycles, n)
	intervals := []float64{}
	totalStart := time.Now()
	for x := 0; x <= cycles; x++ {
		log.Info().Msgf("Starting cycle %d", x)
		start := time.Now()
		max := make(chan bool, 15)
		wg := new(sync.WaitGroup)
		for i := 0; i < n; i++ {
			max <- true
			wg.Add(1)
			go func(serverAddress string, max chan bool, wg *sync.WaitGroup, x, i int) {
				defer func(max chan bool) { <-max }(max)
				defer wg.Done()
				req, _ := http.NewRequest("GET", serverAddress+"/sum?a=1&b=1", nil)
				req.Close = true
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Fatal().Err(err).Msgf("Failed to get sum %d - %d", x, i)
				}
				res.Body.Close()

				sc := res.StatusCode
				if sc != http.StatusOK {
					log.Fatal().Int("statuscode", sc).Msgf("Failed to get sum %d - %d", x, i)
				}

				// log.Info().Int("statuscode", sc).Msgf("Everything is OK %d - %d", x, i)
			}(serverAddress, max, wg, x, i)
		}
		wg.Wait()
		interval := time.Now().Sub(start).Seconds()
		intervals = append(intervals, interval)
		log.Info().Msgf("Cycle %d finished and took %f seconds", x, interval)
	}
	totalInterval := time.Now().Sub(totalStart).Seconds()

	intervalsSum := float64(0)
	intervalsMin := float64(math.MaxFloat64)
	intervalsMax := float64(0)
	for _, v := range intervals {
		intervalsSum += v
		if intervalsMin > v {
			intervalsMin = v
		}
		if intervalsMax < v {
			intervalsMax = v
		}
	}

	intervalsAvg := intervalsSum / float64((len(intervals)))

	log.Info().Msgf("Executed %d cycles of %d requests in %f seconds", cycles, n, totalInterval)
	log.Info().Msgf("MIN: %f seconds", intervalsMin)
	log.Info().Msgf("AVG: %f seconds", intervalsAvg)
	log.Info().Msgf("MAX: %f seconds", intervalsMax)
}

func configureLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func serve() string {

	host := ""
	port := 8082
	serverAddress := fmt.Sprintf("%s:%d", host, port)
	http.HandleFunc("/sum", sum)

	srv := &http.Server{
		Addr: serverAddress,
	}

	go srv.ListenAndServe()
	return fmt.Sprintf("http://localhost:%d", port)
}

func sum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")
	r.Close = true
	defer r.Body.Close()

	a := r.URL.Query().Get("a")
	av, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid a"))
		return
	}

	b := r.URL.Query().Get("b")
	bv, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid b"))
		return
	}

	res := map[string]int{
		"sum": av + bv,
	}

	resB, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(resB)
}

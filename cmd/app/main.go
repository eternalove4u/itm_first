package main

import (
	"flag"
	"fmt"
	"itm_first/internal/port"
	"itm_first/internal/usecase/presenter"
	"itm_first/internal/usecase/producer"
	"os"
)

const (
	defaultOutpuFile = "internal/repository/data/output.txt"
	shortHand        = " (shorthand)"
)

var (
	inputPathPtr  string
	outputPathPtr string
)

func init() {
	flag.StringVar(&inputPathPtr, "input", "", "path to input file")
	flag.StringVar(&inputPathPtr, "i", "", "path to input file"+shortHand)

	flag.StringVar(&outputPathPtr, "output", defaultOutpuFile, "path to output file")
	flag.StringVar(&outputPathPtr, "o", defaultOutpuFile, "path to output file"+shortHand)
}

func main() {
	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Println("use: -i <input path> -o <output path>")
		os.Exit(1)
	}
	flag.Parse()

	prod, err := producer.NewFileProducer(inputPathPtr)
	if err != nil {
		panic(fmt.Errorf("create new producer: %w", err))
	}
	pres, err := presenter.NewFilePresenter(outputPathPtr)
	if err != nil {
		panic(fmt.Errorf("create new presenter: %w", err))
	}
	s, err := port.NewService(prod, pres)
	if err != nil {
		panic(fmt.Errorf("create new service: %w", err))
	}

	if err := s.Run(); err != nil {
		panic(fmt.Errorf("service run: %w", err))
	}
}

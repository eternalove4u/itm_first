package main

import (
	"flag"
	"fmt"
	"itm_first/internal/app/service"
	"itm_first/internal/pkg/presenter"
	"itm_first/internal/pkg/producer"
	"os"
)

var (
	inputPathPtr = flag.String("input",
		"",
		"path to input file",
	)
	outputPathPtr = flag.String("output",
		"cmd/data/output.txt",
		"path to output file",
	)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected -input=\"...\"")
		os.Exit(1)
	}
	flag.Parse()

	prod := producer.NewFileProducer(*inputPathPtr)
	pres := presenter.NewFilePresenter(*outputPathPtr)
	s := service.NewService(prod, pres)

	if err := s.Run(); err != nil {
		fmt.Errorf("%v\n", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"itm_first/internal/port"
	"itm_first/internal/usecase/presenter"
	"itm_first/internal/usecase/producer"
	"os"
)

var (
	inputPathPtr = flag.String("input",
		"",
		"path to input file",
	)
	outputPathPtr = flag.String("output",
		"internal/repository/data/output.txt",
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
	s := port.NewService(prod, pres)

	if err := s.Run(); err != nil {
		fmt.Errorf("%v\n", err)
	}
}

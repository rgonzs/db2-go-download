package main

import (
	"bufio"
	"db2-download/download/config"
	"db2-download/download/core"
	"flag"
	"log"
	"os"
	"sync"
)

func main() {

	ruc := flag.String("ruc", "", "Ruc/Ruc con indicador usado para crear la carpeta")
	filename := flag.String("file", "", "Ruta del archivo con los ids")
	db := flag.String("db", "", "DB a usar con los ids")
	flag.Parse()
	if *ruc == "" && *filename == "" && *db == "" {
		flag.Usage()
		os.Exit(0)
	}

	logfile, err := os.OpenFile(*ruc+".log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(*ruc, *filename, *db)

	wg := &sync.WaitGroup{}

	const threads int = 20

	ids := make(chan string, threads)

	connstr := config.Db_selector(*db)
	db2 := config.Db_connect(connstr)

	for thread := 1; thread <= threads; thread++ {
		go core.ProcessDocument(thread, *ruc, connstr.Table_name, db2, ids, wg)
	}

	// LEEMOS EL ARCHIVO
	file, err := os.OpenFile(*filename, os.O_RDONLY, 0400)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			wg.Add(1)
			ids <- line
		}
	}
	close(ids)

	wg.Wait()
	log.Println("fin")

}

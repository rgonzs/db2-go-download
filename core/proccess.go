package core

import (
	"db2-download/download/repository"
	"log"
	"os"
	"sync"

	"github.com/ibmdb/go_ibm_db"
)

func ProcessDocument(idthread int, ruc string, table_name string, pool *go_ibm_db.DBP, jobs <-chan string, wg *sync.WaitGroup) {
	cdr_path := "downloaded/" + ruc + "/CDR/"
	ubl_path := "downloaded/" + ruc + "/XML/"
	xmldata_path := "downloaded/" + ruc + "/XMLDATA/"
	for id := range jobs {
		data, err := repository.GetDataFromId(id, table_name, pool)
		if err != nil {
			log.Println("Error: ProcessDocument - No existe el id", id)
			log.Println(err)
			log.Println(data)
			wg.Done()
			continue
		}
		if data.ARCHIVORESPUESTA != nil {
			writeBytesToFile(cdr_path+"R-"+data.NOMBREARCHIVO+".zip", data.ARCHIVORESPUESTA)
		}
		if data.ARCHIVOENVIADO != nil {
			writeBytesToFile(ubl_path+data.NOMBREARCHIVO+".zip", data.ARCHIVOENVIADO)
		}
		if data.XMLDATA.Valid {
			writeStringToFile(xmldata_path+data.NOMBREARCHIVO+".xml", data.XMLDATA.String)
		}
		// log.Println("ID procesado ", id)
		wg.Done()
	}
}

func writeBytesToFile(filename string, content []byte) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

}

func writeStringToFile(filename string, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

}

package repository

import (
	"database/sql"
	"fmt"

	"github.com/ibmdb/go_ibm_db"
)

type Document struct {
	NOMBREARCHIVO    string
	ARCHIVOENVIADO   []byte
	ARCHIVORESPUESTA []byte
	XMLDATA          sql.NullString
}

func GetDataFromId(id string, table string, dbconnect *go_ibm_db.DBP) (Document, error) {
	query := fmt.Sprintf(`
        SELECT
        NOMBREARCHIVO,
        ARCHIVOENVIADO,
        ARCHIVORESPUESTA,
        XMLDATA
        from PORTALPERU.%s
        where ID = ?;
        `, table)
	var NOMBREARCHIVO string
	var ARCHIVOENVIADO []byte
	var ARCHIVORESPUESTA []byte
	var XMLDATA sql.NullString
	err := dbconnect.QueryRow(query, id).Scan(&NOMBREARCHIVO, &ARCHIVOENVIADO, &ARCHIVORESPUESTA, &XMLDATA)
	if err != nil {
		return Document{}, err
	}
	// fmt.Println(NOMBREARCHIVO)
	// fmt.Println(XMLDATA)

	return Document{NOMBREARCHIVO, ARCHIVOENVIADO, ARCHIVORESPUESTA, XMLDATA}, nil
}

package utils

import (
	"os"
	"io"
	"log"
	"strings"

	"encoding/csv"
)



/*
	- Reads csv file from path. Returns error if any.
	- Passes records on to provided channel.
	- Will screen blank records if screenBlank set to true. May
	come at slight performance hit. 
	rows x cols: O(rows) ==> O(rows x cols)
*/
func ReadCsv(path string, c chan []string, screenBlank bool) error {
	// read from csv file
	file, err := os.Open(path)
	if err != nil { return err }
	defer file.Close()
	reader := csv.NewReader(file)

	//reader.Read() // skip the header
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error reading csv file: ", err)
			return err
		}

		// screen for blank record
		// ! this may penalize performance
		if screenBlank == true {
			if isBlank(record) == true { continue }
		}

		c <- record
	}

	close(c)
	return nil
}

func isBlank(record []string) bool {
	count := 0
	for _, v := range record {
		if len(strings.Trim(v, " \n\t")) == 0 {
			count++
		}
	}

	if count > 0 { return true }
	return false
}
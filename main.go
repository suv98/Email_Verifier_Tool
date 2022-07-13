package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain", "hasMX", "hasSPF", "spfRecord", "hasDMARC", "dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from the input %v", err)
	}
}

func checkDomain(Domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(Domain)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(Domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_DMARC." + Domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}

	}

	fmt.Printf("%v\n,%v\n,%v\n,%v\n,%v\n,%v\n", Domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}

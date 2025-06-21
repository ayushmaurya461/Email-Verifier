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

	for scanner.Scan() {
		VerifyEmail(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Could not read from input: %v\n", err)
	}
}

func VerifyEmail(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, DMARCRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	DMARCRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)

	}

	for _, record := range DMARCRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCRecord = record
			break
		}
	}

	fmt.Printf("DOMAIN: %v \nMX: %v \nSPF: %v \nSPF Record: %v \nDMARC: %v \nDMARC Records: %v\n-------------------------------\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, DMARCRecord)
}

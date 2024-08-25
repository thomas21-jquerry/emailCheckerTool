package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	fmt.Println("Enter your domain");
	scanner := bufio.NewScanner(os.Stdin);

	for scanner.Scan(){
		checkDomain(scanner.Text())
		break
	}
	if err:=scanner.Err(); err != nil {
		log.Fatal("Error could not read from input %v \n", err)
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool;
	var spfRecord, dmarcRecord string;
	mxRecord, err := net.LookupMX(domain);
	if err != nil {
		log.Printf("error; %v \n", err);
	}
	if len(mxRecord) > 0 {
		hasMX = true;
	}
	txtRecords, err := net.LookupTXT(domain);
	if err != nil {
		log.Printf("error; %v \n", err);
	}

	for _,record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true;
			spfRecord = record;
			break;
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc."+domain);
	if err != nil{
		log.Printf("error; %v \n", err);
	}
	for _,record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true;
			dmarcRecord = record;
			break;
		}
	}

	fmt.Printf(" domain: %v \n hasMX: %v \n hasSPF: %v \n hasDMARC: %v \n SPFRecord: %v \n DMARC Record: %v \n", domain, hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)

}
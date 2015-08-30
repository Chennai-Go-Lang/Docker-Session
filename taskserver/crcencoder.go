package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"log"
)

func init() {
	log.Printf("CRC Server is ready")
	http.HandleFunc("/postcrc", crchandler)

	adrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Panic(err)
		return
	}
	for _, adr := range adrs {
		fmt.Printf("\n Open http://%v:%v", strings.Split(adr.String(), "/")[0], 8080)
	}

}

func main() {

	http.ListenAndServe(":8080", nil)
}

func crchandler(rw http.ResponseWriter, req *http.Request) {
	log.Printf("Hit from  %v %v", req.RemoteAddr, req.Method)

	if req.Method == "POST" {
		rawbytes, rerr := ioutil.ReadAll(req.Body)
		if rerr != nil {
			fmt.Fprintf(rw, "Something Error reading body", rerr)
			log.Println("Something Error reading body", rerr)
		} else {
			log.Printf("Thanks %v, Received %d bytes from your body", req.RemoteAddr, len(rawbytes))
			crcresult := crc32.ChecksumIEEE(rawbytes)
			fmt.Fprintf(rw, "%0x", crcresult)

		}
	} else {

		fmt.Fprintf(rw, "Please send a POST method to see something %s, : %v", req.RemoteAddr, req.Method)

	}
}

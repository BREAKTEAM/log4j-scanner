package url_scan

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	//"sync"
	"time"
)

var total_requests = 0
var hostname = ""

var httpClient = &http.Client{
	Transport: transport,
}

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: time.Second,
		DualStack: true,
	}).DialContext,
}

func Url_hi() {
	fmt.Println("url_hi")
}

func Url_Execute(headers, payloads, urls, hostname string) {
		//PIPELINE urls
		//var urls []string
		// sc := bufio.NewScanner(os.Stdin)
		// for sc.Scan() {
		// 	urls = append(urls, sc.Text())
		// }
		// if err := sc.Err(); err != nil {
		// 	log.Fatal(err)
		// }

		//READ urls from file
		url_file, err := os.Open(urls)
		if err != nil {
			log.Fatal("File could not be read\n")
		}
		defer url_file.Close()
		//time.Sleep(time.Millisecond * 10)
		uScanner := bufio.NewScanner(url_file)
		var urls2 []string
		for uScanner.Scan() {
			urls2 = append(urls2, uScanner.Text())
		}
		if err := uScanner.Err(); err != nil {
			log.Fatal(err)
		}

		//NO USING goroutin
		for _, url := range urls2 {
			request(url, headers, payloads)
		}

		//USING goroutin
		//wg.Wait()
		//close(results)
		//results := make(chan string)
		//var wg sync.WaitGroup
		//for _, url := range urls {
		//wg.Add(1)
		//go func(url string) {
		//request(url, headers, payloads)
		//defer wg.Done()
		//}(url)
		//}
		//wg.Wait()
		//close(results)
	}

func request(urls string, headers string, payloads string) {

	file, err := os.Open(headers)
	if err != nil {
		log.Fatal("File could not be read\n")
	}
	defer file.Close()
	//time.Sleep(time.Millisecond * 10)
	hScanner := bufio.NewScanner(file)
	var lines []string
	for hScanner.Scan() {
		lines = append(lines, hScanner.Text())
	}
	if err := hScanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(lines)

	payload_file, err := os.Open(payloads)
	if err != nil {
		log.Fatal("File could not be read")
	}
	defer payload_file.Close()
	//time.Sleep(time.Millisecond * 10)
	pScanner := bufio.NewScanner(payload_file)
	var links []string
	var links2 []string
	for pScanner.Scan() {
		links = append(links, pScanner.Text())
	}
	if err := pScanner.Err(); err != nil {
		log.Fatal(err)
	}

	//change hostname
	for _, payload := range links {
		payload := strings.Replace(payload, "hostname", hostname, -1)
		links2 = append(links2, payload)
	}

	fmt.Println(links2)

	for _, header := range lines {
		for _, payload := range links2 {
			req, err := http.NewRequest("GET", urls, nil)
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36")
			req.Header.Add(header, payload)
			fmt.Printf("[+] Testing: \t %s\n", header)
			fmt.Printf("[+] Requested: \t %d\n", total_requests)
			total_requests += 1
			if err != nil {
				return
			}
			resp, err := httpClient.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			res, err := httputil.DumpRequest(req, true)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(res))
			fmt.Println(resp.StatusCode)
		}

	}

}

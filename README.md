# Log4j-scanner
### URL mode (fuzzing url with header, payload)
#### go run . url -h
```bash
Usage of url:
-hf string
        path to list header file
  -hn string
        your hostname to connect to (dnslog.cn/burp collabarator...)
  -pl string
        path to list payload file
  -ur string
        path to list url file
Example: go run . url -hf headers.txt -pl payloads.txt -ur urls.txt -hn xxx.burpcollaborator.net
```
### Internal mode (scan Log4j inside your server)
#### go run . internal -h
```bash
Usage of internal:
  -include-zi
        include zip files in the scan
  -mode string
        the output mode, either 'report' (every java archive pretty printed) or 'list' (list of potentially vulnerable files) (default 'report')
  -online
        go with server opton
  -server string
        server to listen result return
Example1: go run . internal C:\ -server 127.0.0.1 -online //if need send to server,run in parallel with the server
Example2: go run . internal C:\ //not send to server
Example3: go run . internal \ //not send to server
```
### External mode(listen to the client's results
#### go run . external -h
```bash
Usage of external:
  -port string
        port to listen result from client (default "8080")
Example: go run . external -port 4444
```

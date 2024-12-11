# go-ifconfig-me

Small Golang library to query https://ifconfig.me/ and return the response as a Golang struct.

https://ifconfig.me is an public, free endpoint that returns detailed network information, 
including your public IP address, user-agent, hostname, and more. 


## How to use

### Simple

```go
client := ifconfigme.NewClient()
response, err := client.Get()
if err != nil {
	log.Fatal(err)
}

log.Println(response)
```

### With Options Pattern

```go
client := ifconfigme.NewClient(
	ifconfigme.WithTimeout(350*time.Millisecond),
	ifconfigme.WithTransport(&http.Transport{}),
)
response, err := client.Get()
if err != nil {
	log.Fatal(err)
}

log.Println(response)
```

### Response

The structure of the response is the following:

- **ip_addr**: Your public IP address.
- **remote_host**: Hostname associated with your IP, if resolvable otherwise returns `unavailable`.
- **user_agent**: User agent string of the client making the request.
- **port**: Port used by the client for the connection.
- **method**: HTTP method used (in this case typically `GET`).
- **encoding**: Encoding specified in the request.
- **mime**: MIME type requested.
- **via**: Proxy details if accessed through a proxy (empty if direct connection).
- **forwarded**: Headers indicating forwarded IPs (empty if none).

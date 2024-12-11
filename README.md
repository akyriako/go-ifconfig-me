# go-ifconfig-me

Small Golang package that helps you discover your public IP address by querying https://ifconfig.me/ and return the response as a Golang struct. 

![image](https://github.com/user-attachments/assets/7ad8099f-061b-4bb2-a31a-265d66f1613b)

## How to use

Import the package as: 	

`import ifconfigme "github.com/akyriako/go-ifconfig-me"`

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

```go
type Response struct {
	IpAddr     string `json:"ip_addr"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
	Port       string `json:"port"`
	Language   string `json:"language"`
	Method     string `json:"method"`
	Encoding   string `json:"encoding"`
	Mime       string `json:"mime"`
	Via        string `json:"via"`
	Forwarded  string `json:"forwarded"`
}
```

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

# etcd-metrics-proxy

Small proxy services that proxies incoming, unauthenticated requests to the etcd metrics endpoint with correct certificates.

## Usage

```
  -caFile string
    	path to client ca file
  -certFile string
    	path to client cert file
  -ip string
    	IP address to bind and proxy to
  -keyFile string
    	path to client key file
  -port string
    	port to bind to 	
```

## Releasing a new version

To release a new version create and push a tag and travis will compile it and create a release.

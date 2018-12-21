# etcd-metrics-proxy

[![Build Status](https://travis-ci.org/syseleven/etcd-metrics-proxy.svg?branch=master)](https://travis-ci.org/syseleven/etcd-metrics-proxy)
[![GitHub license](https://img.shields.io/github/license/syseleven/etcd-metrics-proxy.svg)](https://github.com/syseleven/etcd-metrics-proxy/blob/master/LICENSE)

Small proxy services that proxies incoming, unauthenticated requests to the etcd metrics endpoint with correct certificates.

## Usage

```
  -caFile string
    	path to client ca file
  -certFile string
    	path to client cert file
  -proxyIp string
    	IP address to proxy to
  -bindIp string
    	IP address to bind to  
  -proxyPort string
    	port to proxy to
  -bindPort string
    	port to bind to
  -proxyPath string
    	path to proxy to   	  	
  -keyFile string
    	path to client key file	
```

## Releasing a new version

To release a new version create and push a tag and travis will compile it and create a release.

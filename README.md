# Warden

![Logo](https://github.com/GaruGaru/Warden/blob/master/res/logo.png){:height="50%" width="50%"}

[![Build Status](https://travis-ci.org/GaruGaru/Warden.svg?branch=master)](https://travis-ci.org/GaruGaru/Warden)
[![Go Report Card](https://goreportcard.com/badge/github.com/GaruGaru/Warden)](https://goreportcard.com/report/github.com/GaruGaru/Warden)
![license](https://img.shields.io/github/license/GaruGaru/Warden.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/garugaru/warden.svg)
[![Docker image](https://images.microbadger.com/badges/image/garugaru/warden.svg)](https://microbadger.com/images/garugaru/warden "Get your own image badge on microbadger.com")

## Lightweight, container native host monitor

Warden is a lightweight host monitor designed to be run inside containers that retrieves metrics from the local node and push them on a remote aggregator.


### Run Agent Locally 

*Just Logging*
```bash
	warden agent
```
*Statsd*

```bash
	warden agent --reporter=statsd --statsd_host=localhost:8125 
```

### Run Agent in docker-swarm 

```yaml
version: "3"
services:
  agent:
      image: garugaru/warden
      entrypoint: warden agent
      environment:
        REPORTER: statsd
        STATSD_HOST: ${STATSD_HOST}
      volumes:
        - "/proc:/host/proc"
        - "/sys:/host/sys"
        - "/var:/host/var"
        - "/etc:/host/etc"
      deploy:
        mode: global
```

### Supported Metrics

* Host
	* Hostname
	* UpTime
	* OS
	* Platform
	* PlatformFamily
	* PlatformVersion

* Cpu 
	* Vendor
	* Family
	* Model
	* Cores
	* ModelName
	* Frequency
	* UsagePercent
	* UsagePercentTotal

* Memory
	* Total
	* Used
	* Free
	* UsedPercent

* Disks
	* Name              
	* Mount             
	* Total             
	* Free              
	* Used              
	* UsedPercent       
	* InodesTotal       
	* InodesUsed        
	* InodesFree        
	* InodesUsedPercent 







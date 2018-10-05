# Tracebeat

Tracebeat is an [Elastic Beat](https://www.elastic.co/products/beats) that reads traceroute output and indexes them into Elasticsearch.

## Description

> Traceroute prints the route that packets take to a network host.

It uses [github.com/aeden/tracebeat](https://github.com/aeden/traceroute) for sending packets and tracing routes.

v.1.0.0
## Configuration

Adjust the `tracebeat.yml` configuration file to your needs.

### `period`
Defines how often to take traceroute output. Default to `30` s.

### `host`
Defines the destination host. Default to `8.8.8.8`
To trace the route to a network host pass the ip address of the server you want to connect to.

### `maxhops`
Specifies the maximum number of hops (max time-to-live value) traceroute will probe. Default to `64`

### `retries`
Default to `3`

### `timeoutms`
Default to `500` ms

### `packetsize`
Packet size in byte. Default to `60`

## Document Example

<pre>

    "traceroute": [
      {
        "address": "192.168.1.1",
        "elapsedTime": 4.74324,
        "hopNumber": 1,
        "hostName": "gateway",
        "n": 57,
        "success": true,
        "ttl": 1
      },
      {
        "address": "213.14.0.175",
        "elapsedTime": 14.180505,
        "hopNumber": 2,
        "hostName": "host-213-14-0-175.reverse.superonline.net.",
        "n": 57,
        "success": true,
        "ttl": 2
      },
      {
        "address": "10.36.246.137",
        "elapsedTime": 16.202385,
        "hopNumber": 3,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 3
      },
      {
        "address": "10.34.255.194",
        "elapsedTime": 16.622273,
        "hopNumber": 4,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 4
      },
      {
        "address": "10.38.218.73",
        "elapsedTime": 28.081027,
        "hopNumber": 5,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 5
      },
      {
        "address": "10.38.219.34",
        "elapsedTime": 27.367852,
        "hopNumber": 6,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 6
      },
      {
        "address": "10.40.130.254",
        "elapsedTime": 31.402987,
        "hopNumber": 7,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 7
      },
      {
        "address": "10.36.108.66",
        "elapsedTime": 37.904474,
        "hopNumber": 8,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 8
      },
      {
        "address": "10.36.6.121",
        "elapsedTime": 37.738806,
        "hopNumber": 9,
        "hostName": "",
        "n": 60,
        "success": true,
        "ttl": 9
      },
      {
        "address": "72.14.196.80",
        "elapsedTime": 44.931238,
        "hopNumber": 10,
        "hostName": "",
        "n": 56,
        "success": true,
        "ttl": 10
      },
      {
        "address": "108.170.250.177",
        "elapsedTime": 63.602268,
        "hopNumber": 11,
        "hostName": "",
        "n": 57,
        "success": true,
        "ttl": 11
      },
      {
        "address": "216.239.58.207",
        "elapsedTime": 56.558554,
        "hopNumber": 12,
        "hostName": "",
        "n": 57,
        "success": true,
        "ttl": 12
      },
      {
        "address": "8.8.8.8",
        "elapsedTime": 77.484311,
        "hopNumber": 13,
        "hostName": "google-public-dns-a.google.com.",
        "n": 57,
        "success": true,
        "ttl": 13
      }
    ],
    "type": "tracebeat"
}

</pre>

## Usage
Must be run as sudo.

### Run
To run Tracebeat with debugging output enabled, run:

```
sudo ./tracebeat -c tracebeat.yml -e -d "*" -strict.perms=false

```

## Getting Started with Tracebeat

Ensure that this folder is at the following location:
`${GOPATH}/github.com/berfinsari/tracebeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Tracebeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Tracebeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/berfinsari/tracebeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Tracebeat run the command below. This will generate a binary
in the same directory with the name tracebeat.

```
make
```

### Test

To test Tracebeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/tracebeat.template.json and etc/tracebeat.asciidoc

```
make update
```


### Cleanup

To clean  Tracebeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Tracebeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/berfinsari/tracebeat
cd ${GOPATH}/github.com/berfinsari/tracebeat
git clone https://github.com/berfinsari/tracebeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.

## License
Covered under the Apache License, Version 2.0
Copyright (c) 2017 Berfin Sarı

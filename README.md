# Datadumperbeat

Welcome to Datadumperbeat.

This is a simple Beat that is meant to simulate some traffic for a given

Ensure that this folder is at the following location:
`${GOPATH}/github.com/spantree/datadumperbeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.7


### Build

To build the binary for Datadumperbeat run the command below. This will generate a binary
in the same directory with the name datadumperbeat.

```
make
```


### Run

To run Datadumperbeat with debugging output enabled, run:

```
./datadumperbeat -c datadumperbeat.yml -e -d "*"
```


### Test

To test Datadumperbeat, run the following command:

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
To generate etc/datadumperbeat.template.json and etc/datadumperbeat.asciidoc

```
make update
```


### Cleanup

To clean  Datadumperbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Datadumperbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/spantree/datadumperbeat
cd ${GOPATH}/github.com/spantree/datadumperbeat
git clone https://github.com/spantree/datadumperbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.

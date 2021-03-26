# Installation

## Requirements
Columbus is written in golang, and requires go version >= 1.13. Please make sure that the go tool chain is available on your machine. See golang’s [documentation](https://golang.org/) for installation instructions.

Alternatively, you can use docker to build columbus as a docker image. More on this in the next section.

Columbus uses elasticsearch v7 as the query and storage backend. In order to run columbus locally, you’ll need to have an instance of elasticsearch running.  You can either download elasticsearch and run it manually, or you can run elasticsearch inside docker by running the following command in a terminal
```
$ docker run -d -p 9200:9200 -e "discovery.type=single-node" elasticsearch:7.6.1
```

## Installation
Begin by cloning this repository then you have two ways in which you can build columbus
* As a native executable
* As a docker image

To build columbus as a native executable, run `make` inside the cloned repository.
```
$ make
```

This will create the `columbus` binary in the root directory

Building columbus’s Docker image is just a simple, just run docker build command and optionally name the image
```
$ docker build . -t columbus
```

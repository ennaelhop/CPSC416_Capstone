This was the capstone project that I worked on for a distributed systems course. I worked in a team of 4 people over the course of about 4 - 5 weeks to create this design document.

## Installation and Setup
- Install GoLang from https://golang.org/doc/install
- To build the project, run `go build` in the root directory
- To run the project, run `.\volunteer-computing.exe ` in the root directory
- To run the benchmark, run `go test -benchmem -run=^$ -bench ^BENCH_NAME$ volunteer-computing  ` in the root directory for example `go test -benchmem -run=^$ -bench ^BenchmarkTrackerServer_Parallel$ volunteer-computing` will run the benchmark for the BenchmarkTrackerServer_Parallel function in the main_test.go file

## How this relates to the original design
As discussed in the design document, the aim of our implementation is to abstract away most of the complexities of the distributed system implementation and test our job separation ideas in a fate-shared environment. To do this, we have created a simple tracker server that will keep track of the peers in the network and the jobs that are being processed. The assignment of jobs is done on the peer level but uses a more simplified design and assumes that peers are always available and have enough resources. The peers will be responsible for processing the jobs. Most of the complexities like network communication, job assignment, communication with the tracker will be abstracted away in this implementation. In our simplified model, the user starts as the client and the client uses the tracker server to assign jobs to the peers. The peers will then process the jobs and log the results. This helps us test our job separation ideas in a simple environment.

## What to expect from the output
- The output will be a list of the benchmark results for each benchmark function in the main_test.go file
- To understand the benefit of parallelism, compare the results of the BenchmarkTrackerServer_Parallel function to the BenchmarkTrackerServer function
- If you want to experiment with different number of peers or calculation values, run the project using `./volunteer-computing.exe` and input the desired values when prompted. 
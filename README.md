## Go Data Structures

Overview

This package provides implementations of essential data structures in Go, including graphs and other commonly used structures. It is designed to help developers efficiently manage and manipulate data while maintaining code simplicity and readability.

Features

Implementation of Graph data structures.

Utility functions for adding nodes and edges.

Easy-to-use API for creating and managing data structures.

Installation

To use this package in your Go project, first ensure you have Go installed. Then, initialize a Go module and install the package:

# Initialize a new Go module (if not already done)
go mod init github.com/yourusername/yourrepo

# Get the package
go get github.com/beingbravado/go-data-structures

Usage

Creating a Graph

Below is an example demonstrating how to create and use a graph:

package main

import (
	"fmt"
	"github.com/beingbravado/go-data-structures/graph"
)

func main() {
	g := graph.NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	fmt.Println("Graph Nodes:", g.Nodes())
}

Creating a Graph from Prerequisites

prereqs := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
g := graph.CreateGraph(4, prereqs)

Running Tests

To run unit tests, use:

go test -v ./...

Contributing

Contributions are welcome! If you’d like to add new features, improve documentation, or fix bugs, follow these steps:

Fork the repository.

Create a new branch.

Make your changes and commit them.

Push the changes and submit a pull request.

License

This project is licensed under the MIT License - see the LICENSE file for details.

Contact

For questions or suggestions, feel free to reach out via GitHub Issues.


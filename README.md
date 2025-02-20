# Unexpected Behavior When Accessing Keys in Nil Maps in Go

This repository demonstrates a potential issue in Go when working with maps: accessing a key in a nil map.  Unlike some languages, attempting to access a non-existent key in a nil map in Go does *not* result in a panic; rather, it returns the zero value for the map's value type. This can be subtle and lead to unexpected behavior if not handled properly. 

The `bug.go` file shows an example where this behavior can lead to incorrect results.  `bugSolution.go` provides a solution with improved error handling and illustrates safer ways to work with maps.

## How to Run

1. Clone the repository:
   ```bash
git clone <repository_url>
```
2. Navigate to the project directory:
   ```bash
cd <project_directory>
```
3. Run the buggy code:
   ```bash
go run bug.go
```
4. Run the corrected code:
   ```bash
go run bugSolution.go
```
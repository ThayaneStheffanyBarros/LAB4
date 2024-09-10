package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error reading file %s: %v", filePath, err)
        return nil, err
    }
    return data, nil
}

// sum all bytes of a file
func sum(filePath string, ch chan int) {
    data, err := readFile(filePath)
    if err != nil {
        return
    }

    _sum := 0
    for _, b := range data {
        _sum += int(b)
    }
    ch <- _sum
    return
}

// print the totalSum for all files and the files with equal sum
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <file1> <file2> ...")
        return
    }

    var totalSum int
    sums := make(map[int][]string)

    ch := make(chan int, len(os.Args[1:]))
    for _, path := range os.Args[1:] {
        go sum(path, ch)
    }

    var a int
    for _, path := range os.Args[1:] {
        a = <- ch
        totalSum += a
        sums[a] = append(sums[a], path)
    }
    fmt.Println(totalSum)

    for sum, files := range sums {
        if len(files) > 1 {
            fmt.Printf("Sum %d: %v\n", sum, files)
        }
    }
}
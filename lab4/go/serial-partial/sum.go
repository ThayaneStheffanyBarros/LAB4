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
func sum(filePath string) (int, error, []int) {
    data, err := readFile(filePath)
    if err != nil {
        a := make([]int, 1)
        return 0, err, a
    }

    _sum := 0
    cont := 0
    sums := make([]int, len(data) / 100)
    for _, b := range data {
        _sum += int(b)
        cont += 1
        if cont % 100 == 0 {
            sums = append(sums, _sum)
            _sum = 0
        }
    }

    return _sum, nil, sums
}

func similarity(canal1 []int, canal2 []int) (float64) {
    aux := make([]int, len(canal2))
    copy(aux, canal2)
    cont := 0
    for _, sum := range canal1 {
        if aux[sum] {
            cont += 1
            delete(aux, sum)
        }
    }   
    return cont / len(canal1)
}

// print the totalSum for all files and the files with equal sum
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go  run main.go <file1> <file2> ...")
        return
    }

    var totalSum int64
    sums := make(map[int][]string)
    for _, path := range os.Args[1:] {
        _sum, err, canal := sum(path)

        if err != nil {
            continue
        }

        totalSum += int64(_sum)

        sums[_sum] = append(sums[_sum], path)
    }

    fmt.Println(totalSum)

    for sum, files := range sums {
        if len(files) > 1 {
            fmt.Printf("Sum %d: %v\n", sum, files)
        }
    }
}
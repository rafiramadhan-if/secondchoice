package main

import "fmt"

type arrInt [1000000]int

func insertionSort(T *arrInt, n int) {
    var temp, i, j int
    i = 1
    for i <= n-1 {
        j = i
        temp = T[j]
        for j > 0 && temp < T[j-1] {
            T[j] = T[j-1]
            j = j - 1
        }
        T[j] = temp
        i = i + 1
    }
}

func cekJarak(T *arrInt, n int) {
    if n <= 1 {
        fmt.Println("Data berjarak tidak tetap")
        return
    }

    jarak := T[1] - T[0]
    tetap := true

    var i int
    i = 1
    for i <= n-1 {
        if T[i]-T[i-1] != jarak {
            tetap = false
        }
        i = i + 1
    }

    if tetap {
        fmt.Printf("Data berjarak %d\n", jarak)
    } else {
        fmt.Println("Data berjarak tidak tetap")
    }
}

func main() {
    var T arrInt
    var n, input int

    n = 0

    for {
        fmt.Scan(&input)
        if input < 0 {
            break
        }
        T[n] = input
        n = n + 1
    }

    if n > 0 {
        insertionSort(&T, n)

        var i int
        i = 0
        for i < n {
            if i < n-1 {
                fmt.Printf("%d ", T[i])
            } else {
                fmt.Printf("%d", T[i])
            }
            i = i + 1
        }
        fmt.Println()

        cekJarak(&T, n)
    }
}

package main

import "fmt"

type arrInt [1000000]int

func selectionSort(T *arrInt, n int) {
    var i, j, idx_min, t int
    i = 1
    for i <= n-1 {
        idx_min = i - 1
        j = i
        for j < n {
            if T[idx_min] > T[j] {
                idx_min = j
            }
            j = j + 1
        }
        t = T[idx_min]
        T[idx_min] = T[i-1]
        T[i-1] = t
        i = i + 1
    }
}

func main() {
    var n, m int
    var data arrInt

    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&m)

        for j := 0; j < m; j++ {
            fmt.Scan(&data[j])
        }

        selectionSort(&data, m)

        for j := 0; j < m; j++ {
            if j < m-1 {
                fmt.Printf("%d ", data[j])
            } else {
                fmt.Printf("%d\n", data[j])
            }
        }
    }
}

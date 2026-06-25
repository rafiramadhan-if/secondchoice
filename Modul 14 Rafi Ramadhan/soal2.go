package main

import "fmt"

type arrInt [1000000]int

func selectionSortAsc(T *arrInt, n int) {
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

func selectionSortDesc(T *arrInt, n int) {
    var i, j, idx_max, t int
    i = 1
    for i <= n-1 {
        idx_max = i - 1
        j = i
        for j < n {
            if T[idx_max] < T[j] {
                idx_max = j
            }
            j = j + 1
        }
        t = T[idx_max]
        T[idx_max] = T[i-1]
        T[i-1] = t
        i = i + 1
    }
}

func main() {
    var n, m, input int
    var ganjil, genap arrInt

    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var nGanjil, nGenap int
        fmt.Scan(&m)

        for j := 0; j < m; j++ {
            fmt.Scan(&input)
            if input%2 != 0 {
                ganjil[nGanjil] = input
                nGanjil++
            } else {
                genap[nGenap] = input
                nGenap++
            }
        }

        selectionSortAsc(&ganjil, nGanjil)
        selectionSortDesc(&genap, nGenap)

        firstPrint := true
        for j := 0; j < nGanjil; j++ {
            if firstPrint {
                fmt.Printf("%d", ganjil[j])
                firstPrint = false
            } else {
                fmt.Printf(" %d", ganjil[j])
            }
        }

        for j := 0; j < nGenap; j++ {
            if firstPrint {
                fmt.Printf("%d", genap[j])
                firstPrint = false
            } else {
                fmt.Printf(" %d", genap[j])
            }
        }
        fmt.Println()
    }
}

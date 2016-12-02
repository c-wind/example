package main


func main() {
    //map
    println("---------- map ------------ ")
    m := map[string]string{"k1":"v1", "k2":"v2"}
    for k,v := range m {
        println(k, v)
    }

    // key
    for k := range m {
        println(k)
    }

    println("---------- slice ------------ ")
    //slice
    s := []string{"k1","k2"}
    for i,v := range s {
        println(i, v)
    }

    for i := range s {
        println(i)
    }

    for _, v := range s {
        println(v)
    }

}

## Utility to manipulate object with arbitraty types
This is a small library to manipulate structures (golang struct) that comes from database and converts them into slice of bytes or strings
I did it for a job I had to do to generate a CSV file from a database
the problem I found is that the golang **encoding/csv** library requires that it be []string [**csv Write**](https://golang.org/pkg/encoding/csv/#example_Writer)

## Download/Install
go get github.com/joshvoll/utils

## Convert struct to Slice 
Considering the fallowing struct

```go

    package main 

    import (

        "utils"
    )

    type Person struct {
        Name     string `col:"name"`
        LastName string `col:"Last Name"`
        Age      int `col:"age"`
    }

    func main() {
        x := []Result{Result{"josue", "Rodriguez", 78}, Result{"Alanis", "Rodriguez", 12}}

        res := utils.ToSlice(x)

        fmt.Println(res)
    }

```

``` terminal
    [{josue Rodriguez 78} {Alanis Rodriguez 12}]
```

the difference here is that reflect is used to validate concrete type

## Convert struct to [][]String{}
same struct example
```go

    package main 

    import (

        "utils"
    )

    type Person struct {
        Name     string `col:"name"`
        LastName string `col:"Last Name"`
        Age      int `col:"age"`
    }

    func main() {
        x := []Result{Result{"josue", "Rodriguez", 78}, Result{"Alanis", "Rodriguez", 12}}

        res := utils.GenerateRows(x)

        fmt.Println(res)
    }

```


``` terminal
    [[Name Last Name] [josue Rodriguez] [Alanis Rodriguez]]
```

as you can see this will be **extremely** useful to generate **CSV** files, the generate type is ready to work with golang **encoding/csv** standard library

## Generate CSV example

same struct example
```go

    package main 

    import (
        "encoding/csv"
        "os"
        "utils"
    )

    type Person struct {
        Name     string `col:"name"`
        LastName string `col:"Last Name"`
        Age      int `col:"age"`
    }

    func main() {
        x := []Result{Result{"josue", "Rodriguez", 39}, Result{"Alanis", "Rodriguez", 12}}

        // Using utils library to generate [][]string from struct
        res := utils.GenerateRows(obj)

        // creating the file on the file system
        file, err := os.Create("result.csv")

        if err != nil {
            log.Fatal(err)
        }

        writer := csv.NewWriter(file)

        for _, row := range res {
            // here we write the data to the csv file
            writer.Write(row)
        }
        defer writer.Flush()

    }

```
``` terminal

    Name, Lastname, age
    josue, Rodriguez, 78
    alanis, Rodriguez, 12
```





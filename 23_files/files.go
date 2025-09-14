package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1.Reading from a File
// func main() {
//     data, err := os.ReadFile("example.txt") // read whole file
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("File Content:", string(data))
// }

// 3.Read File in Chunks (Buffer)

// func main() {
//     file, err := os.Open("example.txt")
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer file.Close()

//     buffer := make([]byte, 50) // read 50 bytes at a time
//     for {
//         n, err := file.Read(buffer)
//         if err != nil {
//             break
//         }
//         fmt.Print(string(buffer[:n]))
//     }
// }

// 4.Writing to a File

// func main() {
//     file, err := os.Create("newfile.txt") // create new file
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer file.Close()

//     file.WriteString("Hello, Go File Handling!\n")
//     fmt.Println("File written successfully")
// }

// 5.Appending to a File
// func main() {
//     file, err := os.OpenFile("newfile.txt", os.O_APPEND|os.O_WRONLY, 0644)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer file.Close()

//     file.WriteString("Appending new line...\n")
//     fmt.Println("Content appended")
// }

// 6. Check if File Exists

// func main() {
//     if _, err := os.Stat("example.txt"); os.IsNotExist(err) {
//         fmt.Println("File does not exist")
//     } else {
//         fmt.Println("File exists")
//     }
// }

// 7. Delete a File
// func main() {
//     err := os.Remove("newfile.txt")
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Println("File deleted")
// }

// 8. working with Directories

// func main() {
//     os.Mkdir("myfolder", 0755) // create folder

//     files, err := os.ReadDir(".") // read current directory
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     for _, f := range files {
//         fmt.Println(f.Name())
//     }
// }

// 9. using buifo for line reading

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
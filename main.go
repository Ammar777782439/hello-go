package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("❗ استخدم: add/list/remove")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("❗ أدخل وصف المهمة بعد 'add'")
            return
        }
        task := os.Args[2]
        AddTask(task)
    case "list":
        ListTasks()
    case "remove":
        if len(os.Args) < 3 {
            fmt.Println("❗ أدخل رقم المهمة بعد 'remove'")
            return
        }
        index, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("❌ رقم غير صحيح")
            return
        }
        RemoveTask(index)
    default:
        fmt.Println("❌ أمر غير معروف:", command)
    }
}

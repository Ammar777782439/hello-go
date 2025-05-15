package main

import (
    "fmt"
)

var tasks []string

func AddTask(task string) {
    tasks = append(tasks, task)
    fmt.Println("✅ تمت إضافة المهمة:", task)
}

func ListTasks() {
    if len(tasks) == 0 {
        fmt.Println("📭 لا توجد مهام حالياً.")
        return
    }
    fmt.Println("📋 قائمة المهام:")
    for i, task := range tasks {
        fmt.Printf("%d. %s\n", i+1, task)
    }
}

func RemoveTask(index int) {
    if index < 1 || index > len(tasks) {
        fmt.Println("❌ رقم المهمة غير صحيح.")
        return
    }
    removed := tasks[index-1]
    tasks = append(tasks[:index-1], tasks[index:]...)
    fmt.Println("🗑️ تم حذف المهمة:", removed)
}

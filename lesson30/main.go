package main

import (
    "fmt"
    "log"
    "github.com/Go11Group/Dilfuza/lesson30/model"
    "github.com/Go11Group/Dilfuza/lesson30/storage"
)

func main() {
    database, err := storage.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    userRepo := storage.NewUserRepo(database)
    productRepo := storage.NewProductRepo(database)
    // userProductRepo := storage.NewUserProductRepo(database)

    err = userRepo.Create(&model.User{Username: "Aliyor", Email: "aliyor@example.com ", Password: "password1"})
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }



    user, err := userRepo.Get(1)
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }
    fmt.Printf("User: %+v\n", user)

    product, err := productRepo.Get(1)
    if err != nil {
        log.Fatalf("Failed to get product: %v", err)
    }
    fmt.Printf("Product: %+v\n", product)

    err = userRepo.Update(&model.User{ID: 1, Username: "dilfuzaupdated", Email: "dilfuzaupdated@example.com", Password: "newpassword1"})
    if err != nil {
        log.Fatalf("Failed to update user: %v", err)
    }

    err = productRepo.Update(&model.Product{ID: 1, Name: "Newproduct", Description: "Description of Newproduct1", Price: 22.00, StockQuantity: 90})
    if err != nil {
        log.Fatalf("Failed to update product: %v", err)
    }

    // err = userRepo.Delete(1)
    // if err != nil {
    //     log.Fatalf("Failed to delete user: %v", err)
    // }

    // err = productRepo.Delete(1)
    // if err != nil {
    //     log.Fatalf("Failed to delete product: %v", err)
    // }
}
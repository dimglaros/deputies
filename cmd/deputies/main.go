package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Application started...")

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dsn := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err.Error())
	}

	e, _ := casbin.NewEnforcer("pkg/rbac/model.conf", a)
	err = e.LoadPolicy()
	if err != nil {
		panic(err.Error())
	}
	e.AddFunction("isOwner", IsOwner)
	e.SavePolicy()

	r := mux.NewRouter()
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		getAllUsers(w, r, e)
	}).Methods("GET")
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, e)
	}).Methods("GET")
	//r.HandleFunc("/user/{id}", createUser).Methods("POST")
	//r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	//r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	//r.HandleFunc("/users", memberHandler).Methods("GET")
	//r.HandleFunc("/admin", adminHandler)

	log.Fatal(http.ListenAndServe(":8000", r))

	//err = db.AutoMigrate(&models.User{}, &models.Application{})
	//if err != nil {
	//	panic(err.Error())
	//}
}

func getUser(w http.ResponseWriter, r *http.Request, e *casbin.Enforcer) {
	b, err := e.Enforce("1", "teacher", "1", "/user/:id", "GET")
	if err != nil {
		panic(err.Error())
	}
	log.Println(b)
}

func getAllUsers(writer http.ResponseWriter, request *http.Request, e *casbin.Enforcer) {
	b, err := e.Enforce("1", "teacher", "4", "/users", "GET")
	if err != nil {
		panic(err.Error())
	}
	log.Println(b)
}

func memberHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside member handler")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside admin handler")
}

func isOwner(arg1 string, arg2 string) bool {
	return arg1 == arg2
}

func IsOwner(args ...interface{}) (interface{}, error) {
	arg1 := args[0].(string)
	arg2 := args[1].(string)

	return (bool)(isOwner(arg1, arg2)), nil
}

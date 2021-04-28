package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      string `json:"ID"`
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
}

type Postup struct {
	Name string `json:"name"`
	Post string `json:"post"`
}

type comm struct {
	Name      string `json:"name"`
	Commenter string `json:"commenter"`
	Comment   string `json:"comment"`
}

type forget struct {
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
}

var username = db().Database("fullCollection").Collection("username")
var postsbyuser = db().Database("fullCollection").Collection("post")
var comment = db().Database("fullCollection").Collection("commment")

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user.Name)
	if err != nil {
		log.Fatal(err)
	}
	signupuser, err := username.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signup user are", signupuser)
	json.NewEncoder(w).Encode(user.Name)
}

// get all the user in  the database
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M // .M is map string interface
	cursor, err := username.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var element primitive.M
		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, element)
	}
	cursor.Close(context.TODO())
	for i := range results {
		json.NewEncoder(w).Encode(results[i]["name"])
	}

}

func postuser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var postu Postup
	params := mux.Vars(r)["id"]
	er := json.NewDecoder(r.Body).Decode(&postu)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println("postu", postu.Name)

	var user []map[string]interface{}
	cursor, err := username.Find(context.TODO(), bson.D{primitive.E{Key: "id", Value: params}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var element primitive.M
		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		user = append(user, element)
		fmt.Println(user)
	}
	//fmt.Println(user[0]["name"])
	cursor.Close(context.TODO())
	//for i := range user {
	if user[0]["name"] == postu.Name {
		posts, err := postsbyuser.InsertOne(context.TODO(), postu)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(posts)
		json.NewEncoder(w).Encode(postu)
		//  else {
		// 	w.Write([]byte(fmt.Sprintf("please enter correct username")))
		// }
		//}
	}
}

func getpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	params := mux.Vars(r)["user"]
	fmt.Println(params)
	cursor, err := postsbyuser.Find(context.TODO(), bson.D{primitive.E{Key: "name", Value: params}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var element primitive.M
		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, element)
	}
	fmt.Println(len(results))
	cursor.Close(context.TODO())
	// w.Write([]byte(fmt.Sprint("user :", params)))
	for i := range results {
		w.Write([]byte(fmt.Sprint("user :", params, "\n post  :-", results[i]["post"])))
	}
}

func getcomment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["user"]
	var res2 []primitive.M
	var res1 []primitive.M
	cur1, err := postsbyuser.Find(context.TODO(), bson.D{{Key: "name", Value: params}})
	if err != nil {
		log.Fatal(err)
	}
	for cur1.Next(context.TODO()) {
		var element primitive.M
		err := cur1.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		res1 = append(res1, element)
	}
	cur2, err := comment.Find(context.TODO(), bson.D{{Key: "name", Value: params}})
	if err != nil {
		log.Fatal(err)
	}
	for cur2.Next(context.TODO()) {
		var element primitive.M
		err := cur2.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		res2 = append(res2, element)
	}
	fmt.Println(res2)

	w.Write([]byte(fmt.Sprint("user ", params, "\n has posted  ", res1[0]["post"], "\n and these are the comments ", res2[0]["comment"])))
}
func postcomment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var com comm
	params := mux.Vars(r)["user"]
	err := json.NewDecoder(r.Body).Decode(&com)
	if err != nil {
		fmt.Println(err)
	}
	commm, err := comment.InsertOne(context.TODO(), com)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprint("your comment is live on ", params, "and your comment id is ", commm)))
}

func forgetid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body forget
	var result []primitive.M
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	name1 := body.Name
	email1 := body.EmailId
	cursor, err := username.Find(context.TODO(), bson.D{primitive.E{Key: "name", Value: name1}, {Key: "emailid", Value: email1}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var element primitive.M
		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, element)
	}
	cursor.Close(context.TODO())
	for i := range result {
		w.Write([]byte(fmt.Sprint("\n Your ID is -:  ", result[i]["id"])))
	}
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")
	params := mux.Vars(r)["user"]
	del, err := username.DeleteOne(context.TODO(), bson.M{"name": params})
	if err != nil {
		log.Fatal(err)
	}
	//json.NewEncoder(w).Encode("user", del.DeletedCount)
	w.Write([]byte(fmt.Sprint("user   ", del.DeletedCount)))
	del, err = postsbyuser.DeleteOne(context.TODO(), bson.M{"name": params})
	if err != nil {
		log.Fatal(err)
	}
	//json.NewEncoder(w).Encode("Post", del.DeletedCount)
	w.Write([]byte(fmt.Sprint("Post    ", del.DeletedCount)))
	del, err = comment.DeleteOne(context.TODO(), bson.M{"name": params})
	if err != nil {
		log.Fatal(err)
	}
	//json.NewEncoder(w).Encode("Comment", del.DeletedCount)
	w.Write([]byte(fmt.Sprint("Comment   ", del.DeletedCount)))

}

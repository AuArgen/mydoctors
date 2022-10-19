package main

import(
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"database/sql"
	_"github.com/lib/pq"
)

type Article struct {
	Id uint64
	Title, Anons, Full_text string
}
type Oi_pikir struct {
	Id uint64
	Name, Email, Adress, Message string
}
type Kezek struct {
	Id uint64
	Name, Email, Message string
}
var posts = []Article{}
var showPosts = Article{}

var oiPikir = []Oi_pikir{}
var show_ioPikir = []Oi_pikir{}

var kezeks = []Kezek{}
var showKezeks = []Kezek{}

func index(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()

	res, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}

	posts = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		posts = append(posts,post)
	}

	t.ExecuteTemplate(w, "index.html",posts)

}

func daryger(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/daryger.html")
	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()

	res, err := db.Query("SELECT * FROM oi_pikir")
	if err != nil {
		panic(err)
	}

oiPikir = []Oi_pikir{}

	for res.Next() {
		var post Oi_pikir
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Adress, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		oiPikir = append(oiPikir,post)
	}

	t.ExecuteTemplate(w, "daryger.html",oiPikir)

}

func kezek(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/kezek.html")
	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()

	res, err := db.Query("SELECT * FROM kezek")
	if err != nil {
		panic(err)
	}

	kezeks = []Kezek{}

	for res.Next() {
		var post Kezek
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		kezeks = append(kezeks,post)
	}

	t.ExecuteTemplate(w, "kezek.html",kezeks)

}

func oi_pikir(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/oi-pikir.html")
	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()

	res, err := db.Query("SELECT * FROM oi_pikir")
	if err != nil {
		panic(err)
	}

	oiPikir = []Oi_pikir{}

	for res.Next() {
		var post Oi_pikir
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Adress, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		oiPikir = append(oiPikir,post)
	}

	t.ExecuteTemplate(w, "oi-pikir.html",oiPikir)

}


func search(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/daryger.html")
	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	search :=r.PostFormValue("search")
	// fmt.Println(search)
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()
	if search  == "" {
		res, err := db.Query(fmt.Sprintf("SELECT * FROM oi_pikir"))
			if err != nil {
				panic(err)
			}
		oiPikir = []Oi_pikir{}

	for res.Next() {
		var post Oi_pikir
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Adress, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		oiPikir = append(oiPikir,post)
	}

	t.ExecuteTemplate(w, "daryger.html",oiPikir)
	} else {
		res, err := db.Query(fmt.Sprintf("SELECT * FROM oi_pikir WHERE name LIKE '%s'", search))
			if err != nil {
				panic(err)
			}
		oiPikir = []Oi_pikir{}

	for res.Next() {
		var post Oi_pikir
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Adress, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		oiPikir = append(oiPikir,post)
	}

	t.ExecuteTemplate(w, "daryger.html",oiPikir)
	}

	

}

func sea(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/kezek.html")
	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	search :=r.PostFormValue("search")
	// fmt.Println(search)
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()
	if search  == "" {
		res, err := db.Query(fmt.Sprintf("SELECT * FROM kezek"))
			if err != nil {
				panic(err)
			}
		oiPikir = []Oi_pikir{}

	for res.Next() {
		var post Oi_pikir
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		oiPikir = append(oiPikir,post)
	}

	t.ExecuteTemplate(w, "kezek.html",oiPikir)
	} else {
		res, err := db.Query(fmt.Sprintf("SELECT * FROM kezek WHERE name LIKE '%s'", search))
			if err != nil {
				panic(err)
			}
		oiPikir = []Oi_pikir{}

	for res.Next() {
		var post Oi_pikir
		err = res.Scan(&post.Id, &post.Name, &post.Email, &post.Message)
		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Full_text, post.Id))
		oiPikir = append(oiPikir,post)
	}

	t.ExecuteTemplate(w, "kezek.html",oiPikir)
	}

	

}


func create(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/create.html")

	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 
	t.ExecuteTemplate(w, "create",nil)

}

func save_article(w http.ResponseWriter, r*http.Request) {
	title :=r.PostFormValue("title")
	anons :=r.PostFormValue("anons")
	full_text :=r.PostFormValue("full_text")
	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w,"Pleace enter all input")
	} else {
		connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		} 
		defer db.Close()


		insert,err := db.Query(fmt.Sprintf("INSERT INTO articles(title,anons,full_text) VALUES('%s','%s','%s')",title,anons, full_text))

		if err != nil {
			panic(err)
		} 
		defer insert.Close()

		http.Redirect(w,r,"/", http.StatusSeeOther)
	}
}

func save_kezek(w http.ResponseWriter, r*http.Request) {
	name :=r.PostFormValue("name")
	email :=r.PostFormValue("email")
	message :=r.PostFormValue("message")
	if name == "" || email == "" || message == "" {
		fmt.Fprintf(w,"Pleace enter all input")
	} else {
		connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		} 
		defer db.Close()


		insert,err := db.Query(fmt.Sprintf("INSERT INTO kezek(name,email,message) VALUES('%s','%s','%s')",name,email, message))

		if err != nil {
			panic(err)
		} 
		defer insert.Close()
		http.Redirect(w,r,"/", http.StatusSeeOther)
	}
}

func save_oi_pikir(w http.ResponseWriter, r*http.Request) {
	name :=r.PostFormValue("name")
	email :=r.PostFormValue("email")
	adress :=r.PostFormValue("address")
	message :=r.PostFormValue("message")
	if name == "" || email == "" || message == "" || adress == "" {
		fmt.Fprintf(w,"Pleace enter all input")
	} else {
		connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		} 
		defer db.Close()


		insert,err := db.Query(fmt.Sprintf("INSERT INTO oi_pikir(name,email,message,adress) VALUES('%s','%s','%s','%s')",name,email, message,adress))

		if err != nil {
			panic(err)
		} 
		defer insert.Close()
		http.Redirect(w,r,"/", http.StatusSeeOther)
	}
}

func show_post(w http.ResponseWriter, r*http.Request) {
	t,err := template.ParseFiles("./templates/show.html")

	if err != nil {
		fmt.Fprintf(w,err.Error())
	} 

	vars := mux.Vars(r)
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	} 
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM public.articles WHERE id = %s",vars["id"]))
	if err != nil {
		panic(err)
	}

	showPosts = Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err)
		}

		// fmt.Println("Post: %s with id: %d", post.Full_text, post.Id)
		showPosts = post
	}

	t.ExecuteTemplate(w, "show",showPosts)
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/",index).Methods("GET")
	rtr.HandleFunc("/daryger/",daryger).Methods("GET")
	rtr.HandleFunc("/kezek/",kezek).Methods("GET")
	rtr.HandleFunc("/oi-pikir/",oi_pikir).Methods("GET")
	rtr.HandleFunc("/sea/",sea).Methods("POST")
	rtr.HandleFunc("/search/",search).Methods("POST")
	rtr.HandleFunc("/save_kezek/",save_kezek).Methods("POST")
	rtr.HandleFunc("/save_oi_pikir/",save_oi_pikir).Methods("POST")
	rtr.HandleFunc("/save_article/",save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}",show_post).Methods("GET")
	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)	

}


func main() {
	handleFunc()
}
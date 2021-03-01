// go build -ldflags "-H=windowsgui"
package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/zserge/lorca"
)

// ContactDetails ...
type ContactDetails struct {
	Email   string
	Subject string
	Message string
	Color1  []string
	Color2  []string
}

// ReturnedResult ...
type ReturnedResult struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("%v", "index.html"))) // forms.html
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	/*
		r.ParseForm()
		fmt.Println(r.Form) // print information on server side.
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)

		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", v)
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
			Color1:  r.Form["color1"],
			Color2:  r.Form["color2"],
		}
		fmt.Printf("Post from website! r.PostFrom = %v\n", r.PostForm)
		fmt.Printf("Details = %v\n", details)

		//r.FormValue("username")
		fmt.Println()
		// do something with details
		sheetID := "AKfycbxfMucXOzX15tfU4errRSAa9IzuTRbHzvUdRxzzeYnNA8Ynz8LJuBuaMA/exec"
		url := "https://script.google.com/macros/s/" + sheetID + "/exec"
		bytesRepresentation, err := json.Marshal(details)
		if err != nil {
			log.Fatalln(err)
		}
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			log.Fatalln(err)
		}
		// read all response body
		data, _ := ioutil.ReadAll(resp.Body)

		// close response body
		resp.Body.Close()

		webReturn := ReturnedResult{}
		if err := json.Unmarshal([]byte(data), &webReturn); err != nil {
			panic(err)
		}
		fmt.Println(webReturn.Message)

		webReturn.Message = strings.ReplaceAll(webReturn.Message, "&export=download", "")
		//tmpl.Execute(w, struct{ Success bool }{webReturn.Result})
		tmpl.Execute(w, webReturn)
	*/
}

func main() {
	// Start Host goroutine
	go func() {
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
		http.HandleFunc("/", index)
		http.ListenAndServe(":8090", nil)
	}()

	// Start UI
	ui, err := lorca.New("http://localhost:8090/index", "", 1200, 800)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	ui.Bind("setLanguage", func(data string) {
		// perform the server side task here
		fmt.Println(data)
	})

	<-ui.Done()
}

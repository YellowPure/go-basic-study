package main

import (
	"errors"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	templates = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html"))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	addr      = flag.Bool("addr", false, "find open address and print to final-port.txt")
	search    = regexp.MustCompile("\\[([a-zA-Z]+)\\]")
)

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("./data/"+filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("./data/" + filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(w http.ResponseWriter, req *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(req.URL.Path)
	if m == nil {
		http.NotFound(w, req)
		return "", errors.New("404")
	}
	return m[2], nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, req *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, req, "/edit/"+title, http.StatusFound)
		return
	}
	p.Body = search.ReplaceAllFunc(p.Body, func(s []byte) []byte {
		group := search.ReplaceAllString(string(s), `$1`)
		newGroup := "<a href='/view/" + group + "'>" + group + "</a>"
		return []byte(newGroup)
	})

	renderTemplate(w, "view", p)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, req *http.Request, title string) {
	// title := req.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	p := &Page{
		Title: title,
		Body:  []byte(body),
	}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		m := validPath.FindStringSubmatch(req.URL.Path)
		if m == nil {
			http.NotFound(w, req)
			return
		}
		fn(w, req, m[2])
	}
}

func main() {
	// p1 := &Page{
	// 	Title: "TestPage",
	// 	Body:  []byte("这是实力页面。"),
	// }
	// p1.Save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	flag.Parse()
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.Redirect(rw, r, "/view/FrontPage", http.StatusFound)
	})
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}

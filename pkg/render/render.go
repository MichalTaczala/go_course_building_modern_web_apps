package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
		return
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("not found")
		return
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("render failure")
		return
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error")
	// 	return
	// }
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCashe := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCashe, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCashe, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCashe, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCashe, err
			}
		}
		myCashe[name] = ts
	}
	return myCashe, nil

}

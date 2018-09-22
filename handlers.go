package main 

import(
 "fmt"
 "html"
 "net/http"
 "encoding/json"
)

func Help(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, This is the health page");
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func GetAge(w http.ResponseWriter, r *http.Request){
	fmt.Println("GET params are: ", r.URL.Query())

	dob := r.URL.Query().Get("dob")
	//fmt.Fprintf(w, "Received age, %q", dob)
	//fmt.Fprintf(w, "Received age, %q", GetAgeInYear(dob))
	m := make(map[string]string)
	m["Age"] = GetAgeInYear(dob)

	if err := json.NewEncoder(w).Encode(m); err != nil {
        panic(err)
    }

}

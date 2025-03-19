// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"html/template"
// )

// // Template for the HTML form
// var formTemplate = `
// <!DOCTYPE html>
// <html>
// <head>
//     <title>Welcome Page</title>
// </head>
// <body>
// 	<form action="/welcome" method="POST">
// 		<label for="fname">First Name:</label><br>
// 		<input type="text" id="fname" name="firstname" placeholder="Enter your first name" required><br>
// 		<label for="lname">Last name:</label><br>
// 		<input type="text" id="lname" name="lastname" placeholder="Enter your last name" required><br>
// 		<label for="dob">DOB:</label><br>
// 		<input type="date" id="dob" name="dob"><br>
// 		<button type="submit">Submit</button>
// 	</form> 
// </body>
// </html>
// `

// // Template for the welcome message
// var welcomeTemplate = `
// <!DOCTYPE html>
// <html>
// <head>
//     <title>Welcome</title>
// </head>
// <body>
//     <h1>Welcome, {{.}}</h1>
// </body>
// </html>
// `

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "This is the home page for terateck\n")
// 		fmt.Fprintf(w, "This is my second paragraph\n")
// 		fmt.Fprintf(w, formTemplate)
// 	})

// 	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodPost {
// 			r.ParseForm()
// 			name := r.FormValue("firstname")

// 			// Render the welcome message
// 			tmpl := template.Must(template.New("welcome").Parse(welcomeTemplate))
// 			tmpl.Execute(w, name)
// 		} else {
// 			http.Redirect(w, r, "/", http.StatusSeeOther)
// 		}
// 	})

// 	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "This is my admin page\n")
// 		fmt.Fprintf(w, "This is my second go line program")
// 	})

// 	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "This is my about page for terateck\n")
// 		fmt.Fprintf(w, "This is my second paragraph")
// 	})

// 	fmt.Println("Server started on :8082")
// 	http.ListenAndServe(":8082", nil)
// }


package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// HTML template for the form
var formTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Welcome Page</title>
</head>
<body>
	<h1>Enter Your Details</h1>
	<form action="/welcome" method="POST">
		<label for="fname">First Name:</label><br>
		<input type="text" id="fname" name="firstname" placeholder="Enter your first name" required><br><br>
		
		<label for="lname">Last Name:</label><br>
		<input type="text" id="lname" name="lastname" placeholder="Enter your last name" required><br><br>
		
		<label for="dob">Date of Birth:</label><br>
		<input type="date" id="dob" name="dob"><br><br>
		
		<button type="submit">Submit</button>
	</form> 
</body>
</html>
`

// HTML template for the welcome page
var welcomeTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Welcome</title>
</head>
<body>
    <h1>Welcome, {{.FirstName}} {{.LastName}}</h1>
    <p>Your Date of Birth: {{.DOB}}</p>
</body>
</html>
`

// Struct to hold form data
type User struct {
	FirstName string
	LastName  string
	DOB       string
}

func main() {
	// Parse form template
	tmplForm := template.Must(template.New("form").Parse(formTemplate))
	// Parse welcome template
	tmplWelcome := template.Must(template.New("welcome").Parse(welcomeTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html") // Ensure content type is HTML
		tmplForm.Execute(w, nil)                    // Render the form template
	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			user := User{
				FirstName: r.FormValue("firstname"),
				LastName:  r.FormValue("lastname"),
				DOB:       r.FormValue("dob"),
			}
			w.Header().Set("Content-Type", "text/html") // Ensure content type is HTML
			tmplWelcome.Execute(w, user)                // Render the welcome page
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	fmt.Println("Server started on :8082")
	http.ListenAndServe(":8082", nil)
}


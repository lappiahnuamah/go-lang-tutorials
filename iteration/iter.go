package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
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

// HTML template for displaying all users
var welcomeTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Welcome</title>
</head>
<body>
	<h1>Welcome, Here are all registered users:</h1>  
	<ol>
        {{range $index, $user := .}}
            <li><strong>{{$user.FirstName}} {{$user.LastName}}</strong> </li>
        {{end}}
    </ol>
    <br>
    <a href="/">Go Back</a>
</body>
</html>
`

// User struct
type User struct {
	FirstName string
	LastName  string
}

// Slice to store users
var users []User
var mu sync.Mutex // Prevents race conditions when modifying users slice

func main() {
	// Parse templates
	tmplForm := template.Must(template.New("form").Parse(formTemplate))
	tmplWelcome := template.Must(template.New("welcome").Parse(welcomeTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmplForm.Execute(w, nil)
	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			newUser := User{
				FirstName: r.FormValue("firstname"),
				LastName:  r.FormValue("lastname"),
			}

			// Store user in slice
			mu.Lock()
			users = append(users, newUser)
			mu.Unlock()

			w.Header().Set("Content-Type", "text/html")
			tmplWelcome.Execute(w, users) // Pass users slice to template
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	fmt.Println("Server started on :8082")
	http.ListenAndServe(":8082", nil)
}

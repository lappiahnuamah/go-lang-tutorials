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

// HTML template for displaying all users (with numbering & auto-incremented usernames)
var welcomeTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Welcome</title>
</head>
<body>
    <h1>Welcome, Here are all registered users:</h1>
    <ol>
        {{range .}}
            <li><strong>{{.Username}}</strong> - {{.FirstName}} {{.LastName}} (DOB: {{.DOB}})</li>
        {{end}}
    </ol>
    <br>
    <a href="/">Go Back</a>
</body>
</html>
`

// User struct
type User struct {
	Username  string
	FirstName string
	LastName  string
	DOB       string
}

// Slice to store users
var users []User
var mu sync.Mutex // Prevents race conditions when modifying users slice
var userCount int // Keeps track of the number of users

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

			mu.Lock()
			userCount++ // Increment user count
			newUser := User{
				Username:  fmt.Sprintf("user%d", userCount), // Assign unique username
				FirstName: r.FormValue("firstname"),
				LastName:  r.FormValue("lastname"),
				DOB:       r.FormValue("dob"),
			}

			users = append(users, newUser) // Store user in slice
			mu.Unlock()

			w.Header().Set("Content-Type", "text/html")
			tmplWelcome.Execute(w, users) // Pass users slice to template
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	fmt.Println("Server started on :8081")
	http.ListenAndServe(":8081", nil)
}

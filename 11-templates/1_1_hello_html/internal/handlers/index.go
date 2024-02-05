package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func (Server) index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	_, err := io.WriteString(w, `
<head>
	<title>Learn Go</title>
</head>
<body>
	<h2>YAY</h2>
	<p>Way to go!</p>
</body>
`)

	if err != nil {
		fmt.Println("Error in index:", err)
	}
}

// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/kahunacohen/songs/controllers"
)

func initDB(ctx context.Context) (*pgx.Conn, error) {
	connStr := "postgresql://postgres:password@postgres:5432/songs?sslmode=disable"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn, err
}
func main() {
	ctx := context.Background()
	conn, err := initDB(ctx)
	if err != nil {
		log.Fatalf("failed to intialize db: %v", err)
	}
	defer conn.Close(ctx)
	router := gin.Default()
	router.Use(ResponseFormatMiddleware)
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/api/v1/users/:user_id/:song_id", controllers.SongsByUserHandler(conn))
	router.GET("/users/:user_id/:song_id", controllers.SongsByUserHandler(conn))
	router.GET("/", func(c *gin.Context) {
		data := gin.H{
			"Title": "Gin Example",
		}

		// Render the HTML template with the provided data
		c.HTML(http.StatusOK, "base.html", data)
	})

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// package main

// import (
//     "html/template"
//     "net/http"
//     "github.com/gin-gonic/gin"
// )

// type PageData struct {
//     Title   string
//     Content string
// }

// func main() {
//     router := gin.Default()

//     // Define routes
//     router.GET("/page_a", func(c *gin.Context) {
//         data := PageData{
//             Title:   "Page A",
//             Content: "This is the content for Page A.",
//         }

//         // Execute the "base.html" template with the "content_a.html" template
//         if err := executeTemplates(c.Writer, "base.html", "content_a.html", data); err != nil {
//             c.String(http.StatusInternalServerError, err.Error())
//         }
//     })

//     router.GET("/page_b", func(c *gin.Context) {
//         data := PageData{
//             Title:   "Page B",
//             Content: "This is the content for Page B.",
//         }

//         // Execute the "base.html" template with the "content_b.html" template
//         if err := executeTemplates(c.Writer, "base.html", "content_b.html", data); err != nil {
//             c.String(http.StatusInternalServerError, err.Error())
//         }
//     })

//     router.Run(":8080")
// }

// // executeTemplates executes the base template with the specified content template
// func executeTemplates(w http.ResponseWriter, baseTemplate, contentTemplate string, data PageData) error {
//     // Load the base template and content template
//     templates := template.Must(template.ParseFiles("templates/" + baseTemplate, "templates/" + contentTemplate))

//     // Execute the base template with the content template
//     return templates.ExecuteTemplate(w, baseTemplate, data)
// }

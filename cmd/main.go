package main
import (
	"context"
	"github.com/gorilla/mux"
	"github.com/roobiuli/MongoCRUD/internal/comment"
	"github.com/roobiuli/MongoCRUD/internal/pkg/storage/mongodb"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Spawn DB Conf
	DbCF := mongodb.DatabaseConfig{
		Auth: os.Getenv("AUTH"),
		Host: os.Getenv("DBHOST"),
		Port: os.Getenv("DBPORT"),
		User: os.Getenv("DBUSER"),
		Pass: os.Getenv("DBPASS"),
		Name: os.Getenv("DBNAME"),
	}

	// DB Connection
	MDB, err := mongodb.NewDatabase(ctx, DbCF)

	if err != nil {
		log.Fatalln(err)
	}

	defer MDB.Client().Disconnect(ctx)

	// Setting up the Storage Sys
	StorManager := mongodb.CommentStorage{
		Database: MDB,
		Timeout:  time.Second * 5,
	}

	CommentController := comment.Controller{Storage: StorManager}


	r := mux.NewRouter()

	r.HandleFunc("/api/v1/comments", CommentController.Create)
	r.HandleFunc("/api/v1/comments:{uuid}", CommentController.Find)
	r.HandleFunc("/api/v1/comments:{uuid}", CommentController.Delete)
	r.HandleFunc("/api/v1/comments:{uuid}", CommentController.Update)


	log.Fatalln(http.Server{
		Addr:              os.Getenv("APPPORT"),
		Handler:           r,
	})
}

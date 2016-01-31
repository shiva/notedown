package main

import (
    "os"
    "fmt"
    "log"
    "net/http"

	"github.com/astaxie/beego/session"
	"github.com/codegangsta/negroni"
    "github.com/jessevdk/go-flags"
)

type Options struct {
    // Slice of bool will append 'true' each time the option
    // is encountered (can be set multiple times, like -vvv)
    Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information."`
    MongoUri string `short:"s" long:"mongo-server" required:"true" description:"Mongo server URI."`
    MongoUser string `short:"u" long:"mongo-user" required:"true" description:"Mongo user name."`
    MongoPassword string `short:"p" long:"mongo-password" required:"true" description:"Mongo password."`
}

var (
	GlobalSessions *session.Manager
    repo            Repo
    opts            Options
)

var parser = flags.NewParser(&opts, flags.Default)

func main() {

    if _, err := parser.Parse(); err != nil {
        panic(err)
        os.Exit(1)
    }

    if len(opts.Verbose) >= 1 {
        fmt.Printf("Mongo Uri: %s\n", opts.MongoUri)
        fmt.Printf("Mongo User: %s\n", opts.MongoUser)
        fmt.Printf("Mongo Password: %s\n", opts.MongoPassword)
    }

	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go GlobalSessions.GC()

    repo.Init(MongoProperties{opts.MongoUri, opts.MongoUser, opts.MongoPassword})
    rtr := NewRouter()
    n := negroni.Classic()
    n.UseHandler(rtr)

    log.Fatal(http.ListenAndServe(":8080", n))
}

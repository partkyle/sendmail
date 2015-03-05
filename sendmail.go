package main

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
)

var (
	host = flag.String("host", "localhost", "host to connect")
	port = flag.Int("port", 25, "port to connect to")

	user = flag.String("user", "", "user to connect as")
	pass = flag.String("pass", "", "pass to connect as")

	domain = flag.String("domain", "testing", "domain for EHLO")
)

func main() {
	flag.Parse()

	client, err := smtp.Dial(fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	{
		err := client.Hello(*domain)
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		auth := smtp.PlainAuth("", *user, *pass, *host)
		err := client.Auth(auth)
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		err := client.Mail("kyle.partridge@sendgrid.com")
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		err := client.Rcpt("kyle.partridge@sendgrid.com")
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		w, err := client.Data()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "hello you")
		if err := w.Close(); err != nil {
			log.Fatal(err)
		}
	}

	{
		err := client.Quit()
		if err != nil {
			log.Fatal(err)
		}
	}
}

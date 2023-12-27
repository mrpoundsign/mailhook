package main

import (
	"flag"
	"log"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/mrpoundsign/mailhook/pkg/config"
	"github.com/mrpoundsign/mailhook/pkg/mail"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "c", "/etc/mailhook.yaml", "config file path")
	flag.Parse()

	conf, err := config.NewMailHook(configPath)
	if err != nil {
		log.Fatal(err)
	}

	be := mail.NewBackend(conf.Auth.Username, conf.Auth.Password)
	for _, hook := range conf.Hooks {
		log.Println("Added hook", hook.Name)
		be.AddHook(hook)
	}

	s := smtp.NewServer(be)

	s.Addr = conf.ListenString()
	s.Domain = "localhost"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 1
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

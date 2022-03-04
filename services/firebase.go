package services

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitializeFirebase() {
	serviceAccountKeyPath, err := filepath.Abs("./firebaseServiceAccountKey.json")
	if err != nil {
		panic("[ERROR] Unable to load firebaseServiceAccountKey.json")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("[ERROR] Firebase load error")
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("[ERROR] Firebase load error")
	}

	FirebaseAuth = auth
}

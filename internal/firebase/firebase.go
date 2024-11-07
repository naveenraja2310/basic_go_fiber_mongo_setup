/*
firebase/firebase.go
Author: Naveen Raj
Description:This file is responsible for initilizing the firebase.
*/

package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

// InitFirebase is to read and initialize the firebase
func InitFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("../../firebase-config.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	FirebaseApp = app
}

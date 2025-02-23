package firebase

import (
	"context"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"github.com/huydq/gokits/libs/utilities/ijson"
)

type FirebaseAuth struct {
	*auth.Client
}

func InstallFirebaseAuth(configKeys ...string) *FirebaseAuth {
	if firebaseConfig == nil {
		getFirebaseConfigs(configKeys...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	opt := option.WithCredentialsJSON(ijson.ToJsonByte(firebaseConfig))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, errAuth := app.Auth(ctx)
	if errAuth != nil {
		panic(errAuth)
	}

	return &FirebaseAuth{Client: client}
}

func (f *FirebaseAuth) GetUser(jwtToken string) (*auth.UserRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	token, err := f.VerifyIDToken(ctx, jwtToken)
	if err != nil {
		return nil, err
	}

	user, err := f.Client.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

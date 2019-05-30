// Package main implements a client for Greeter service.
package main

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	napi "np-10/rpc-test/api"
	ngrpc "np-10/rpc-test/apigrpc"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	//ngrpc "github.com/heroiclabs/nakama/apigrpc"
	//napi "github.com/heroiclabs/nakama/api"
	// wrappers "np-10/rpc-test/wrappers"
	//vgrpc "github.com/heroiclabs/nakama/vender/google.golang.org/grpc"
)

const (
	address = "x.x.x.x:7349"
)

type basicAuth struct {
	username string
	password string
}

func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.username + ":" + b.password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

func (basicAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		//grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithPerRPCCredentials(basicAuth{"defaultkey", ""}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := ngrpc.NewNakamaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. populate urlpath = "/v2/account/authenticate/email?";
	// 2. populate message body is create, username,
	// 3. compose http request message, http method "POST", header type
	// 4. basic authentication user name & passwrod
	// var credentials = Encoding.UTF8.GetBytes(basicAuthUsername + ":" + basicAuthPassword);
	// 5. Basic credentials
	// 		var header = string.Concat("Basic ", Convert.ToBase64String(credentials));
	// 		request.Headers.Authorization = AuthenticationHeaderValue.Parse(header);
	// to json format
	// 6. var response = await _httpClient.SendAsync(request);
	// 7. var contents = await response.Content.ReadAsStringAsync();

	// Login with email account
	// populate account email request message
	pAcc := &napi.AccountEmail{Email: "ab1@abc.com", Password: "12345678"}
	bCreate := &wrappers.BoolValue{Value: true}
	pEmailReq := &napi.AuthenticateEmailRequest{Account: pAcc, Create: bCreate, Username: "ab1"}

	session, error := c.AuthenticateEmail(ctx, pEmailReq)
	if error != nil {
		log.Fatalf("could not AuthenticateEmail: %v", error)
	}
	log.Println(session)

	// try add friend
	pAddFriends := &napi.AddFriendsRequest{Ids: nil, Usernames: []string{"ab2"}}
	_, error = c.AddFriends(ctx, pAddFriends)
	if error != nil {
		log.Fatalf("could not AddFriends: %v", error)
	}
}

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"

	napi "np-10/grpc-test/api"
	ngrpc "np-10/grpc-test/apigrpc"
)

const (
	address = "x.x.x.x:7349"
)

type basicAuth struct {
	username string
	password string
}

// to setup http header: --user 'defaultkey:'
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

type bearerAuth struct {
	sessionToken string
}

// to setup http header:  -H 'Authorization: Bearer <session token>'
func (b bearerAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + b.sessionToken,
	}, nil
}

func (bearerAuth) RequireTransportSecurity() bool {
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

	conn2, err2 := grpc.Dial(address, grpc.WithInsecure(),
		//grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithPerRPCCredentials(bearerAuth{session.GetToken()}))
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	defer conn2.Close()
	c2 := ngrpc.NewNakamaClient(conn2)

	// try add friend
	pAddFriends := &napi.AddFriendsRequest{Ids: nil, Usernames: []string{"ab2"}}
	_, error = c2.AddFriends(ctx, pAddFriends)
	if error != nil {
		log.Printf("could not AddFriends: %v\n", error)
	}

	friendList, errLF := c2.ListFriends(ctx, &empty.Empty{})
	if errLF != nil {
		log.Fatalf("could not ListFriends: %v", errLF)
	}
	log.Println(friendList)
}

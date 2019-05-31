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

	"np-10/grpc-test/api"
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

//　NakamaGrpcAuth　Email authenticate and get user's session token
func NakamaGrpcAuth(svrAddr string, userName string, email string, pwd string, timeout time.Duration) (*api.Session, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(svrAddr, grpc.WithInsecure(),
		//grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithPerRPCCredentials(basicAuth{"defaultkey", ""}))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := ngrpc.NewNakamaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 1. populate urlpath = "/v2/account/authenticate/email?";
	// 2. populate message body is create, username,
	// 3. compose http request message, http method "POST", header type
	// 4. basic authentication user name & passwrod "defaultkey:"
	// 5. Basic credentials
	// 		var header = string.Concat("Basic ", Convert.ToBase64String(credentials));
	// 		request.Headers.Authorization = AuthenticationHeaderValue.Parse(header);
	// to json format
	// 6. var response = await _httpClient.SendAsync(request);

	// Login with email account
	// populate account email request message
	pAcc := &napi.AccountEmail{Email: email, Password: pwd}
	bCreate := &wrappers.BoolValue{Value: true}
	pEmailReq := &napi.AuthenticateEmailRequest{Account: pAcc, Create: bCreate, Username: userName}

	return c.AuthenticateEmail(ctx, pEmailReq)
}

// NewNakamaGrpcClientBySessionToken: use user's session token to get Nakama grpc client
func NewNakamaGrpcClientBySessionToken(svrAddr string, sessionToken string) (*grpc.ClientConn, ngrpc.NakamaClient, error) {
	conn, err := grpc.Dial(svrAddr, grpc.WithInsecure(),
		//grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithPerRPCCredentials(bearerAuth{sessionToken}))
	if err != nil {
		return nil, nil, err
	}
	return conn, ngrpc.NewNakamaClient(conn), nil
}

// RpcFunc to call namaka server plugin rpc functions
func RpcFunc(ctx context.Context, nc ngrpc.NakamaClient, id string, payload string) {
	pRPCReq := &napi.Rpc{Id: id, Payload: payload, HttpKey: "defaultkey"}
	pRPCRes, errRpc := nc.RpcFunc(ctx, pRPCReq)

	if errRpc != nil {
		log.Printf("could not call Rpc: %v\n", errRpc)
	} else {
		log.Println(pRPCRes.Id + "\n" + pRPCRes.Payload)
	}
}

func main() {
	session, error := NakamaGrpcAuth(address, "ab1", "ab1@abc.com", "12345678", 3*time.Second)
	if error != nil {
		log.Fatalf("could not AuthenticateEmail: %v", error)
	}
	log.Println(session)

	conn2, c2, err2 := NewNakamaGrpcClientBySessionToken(address, session.GetToken())
	if err2 != nil {
		log.Fatalf("did not connect: %v", err2)
	}
	defer conn2.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

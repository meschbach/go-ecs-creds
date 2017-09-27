package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	fmt.Println( "ECS credentials" )

	sess := session.Must(session.NewSession())
	svc := sts.New( sess )
	input := &sts.GetCallerIdentityInput{}

	http.HandleFunc( "/", func( w http.ResponseWriter, r *http.Request ) {
		result, err := svc.GetCallerIdentity(input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					default:
						fmt.Fprintf( w, aerr.Error() )
					}
			} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
				fmt.Fprintf( w, err.Error() )
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder( w ).Encode( result )
	} )
	log.Fatal(http.ListenAndServe(":8080", nil))
}


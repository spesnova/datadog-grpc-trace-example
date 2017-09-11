/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	grpctrace "github.com/DataDog/dd-trace-go/contrib/google.golang.org/grpc"
	"github.com/DataDog/dd-trace-go/tracer"
)

const (
	address               = "localhost:50051"
	defaultName           = "world"
	datadogAPMServiceName = "greeter-client"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpctrace.UnaryClientInterceptor(datadogAPMServiceName, tracer.DefaultTracer),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootSpan := tracer.DefaultTracer.NewRootSpan("web.request", datadogAPMServiceName, "/")
		defer rootSpan.Finish()
		ctx := rootSpan.Context(r.Context())

		// Contact the server and print out its response.
		name := defaultName
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			rootSpan.FinishWithErr(err)
			log.Fatalf("[ERROR] Could not greet: %v", err)
			return
		}

		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp.Message + "\n"))
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("[ERROR] Failed to start server: %s", err)
	}
	log.Println("[INFO] Listening server on :3000")
}

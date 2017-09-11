# Datadog gRCP Tracing Example
This repo shows you how to trace gRPC with Datadog APM (go tracer).

## Instrumentation
### Client Side

```go
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpctrace.UnaryClientInterceptor(datadogAPMServiceName, tracer.DefaultTracer),
		),
	)
```

You can see the details here: [greeter_client/main.go](https://github.com/spesnova/datadog-grpc-trace-example/blob/master/greeter_client/main.go)

### Server Side

```go
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpctrace.UnaryServerInterceptor(datadogAPMServiceName, tracer.DefaultTracer),
		),
	)
```

You can see the details here: [greeter_server/main.go](https://github.com/spesnova/datadog-grpc-trace-example/blob/master/greeter_server/main.go)


## Screenshots

![2017-09-11 17 33 35](https://user-images.githubusercontent.com/261700/30265613-9c358926-9717-11e7-8a7d-c3c95f00d73d.png)
![2017-09-11 17 33 52](https://user-images.githubusercontent.com/261700/30265612-9c30f7d0-9717-11e7-9692-96f4bdf57c9e.png)
![2017-09-11 17 33 56](https://user-images.githubusercontent.com/261700/30265615-9c38e5e4-9717-11e7-95b9-0726a00ded92.png)
![2017-09-11 17 34 01](https://user-images.githubusercontent.com/261700/30265616-9c3d81d0-9717-11e7-9096-47f6c9533bb5.png)
![2017-09-11 17 34 10](https://user-images.githubusercontent.com/261700/30265614-9c36f1ee-9717-11e7-9bb5-917fe93acec6.png)

## Resources

- [Tracing Overview - Datadog docs](https://docs.datadoghq.com/tracing/)
- [Tracing Go Applications - Datadog docs](https://docs.datadoghq.com/tracing/go/)
- [package tracer - GoDoc](https://godoc.org/github.com/DataDog/dd-trace-go/tracer)

package main

import ("log"
       "net"
	   "google.golang.org/grpc"
      ) 


func main(){
	lis,err:=net.Listen("tcp",":5000")
	if err!=nil{
		log.Fatal(err)
	}

	grpcserver:=grpc.NewServer()

	if err := grpcserver.Serve(lis); err != nil{
		log.Fatal("Failed To serve Grpc",err)
	}
}

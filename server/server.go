package server

import (
 "context"
 "log"
 "os"

 protos "github.com/kamalbahsine/simple-grpc-app/protos/podinfos"
)

type PodInfo struct {
 protos.UnimplementedPodInfoServer
}

func NewPodInfo() *PodInfo {
 return &PodInfo{}
}

func (t *PodInfo) GetPodInfo(
 ctx context.Context,
 i *protos.EmptyMessage,
) (*protos.PodInfoOutput, error) {

    pod_name := os.Getenv("POD_NAME")
 worker_name := os.Getenv("WORKER_NAME")
 version := os.Getenv("VERSION")

 log.Printf(
   `GetPodInfo for : "%s" hosted on : "%s" having application version : "%s"`,
   pod_name, worker_name, version,
 )

 info := &protos.PodInfoOutput{
   PodName:        pod_name,
   WorkerName:  worker_name,
   Version:  version,
 }

 return info, nil
}

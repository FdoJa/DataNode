package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	pb "github.com/FdoJa/DataNode1/proto"
	"google.golang.org/grpc"
)

var lock = &sync.RWMutex{}
var err error

type dataNodeServer struct {
	pb.UnimplementedDataNodeServer
}

func (s *dataNodeServer) RegistrarNombre(ctx context.Context, registro *pb.Registro) (*pb.Recepcion, error) {
	lock.Lock()
	defer lock.Unlock()

	filePath := "/app/Data.txt"
	var file *os.File

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Printf("Agregando a txt: %s %s %s", registro.Id, registro.Nombre, registro.Apellido)
	data := registro.Id + " " + registro.Nombre + " " + registro.Apellido + " \n"
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}

	return &pb.Recepcion{
		Ok: "OK",
	}, nil
}

func (s *dataNodeServer) Solicitud_Info_DataNode(ctx context.Context, idList *pb.Id) (*pb.Lista_Datos_DataNode, error) {
	filePath := "/app/Data.txt"
	var file *os.File

	file, err = os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	personaMap := make(map[string]*pb.Datos_DataNode)

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		id := data[0]
		nombre := data[1]
		apellido := data[2]
		personaMap[id] = &pb.Datos_DataNode{
			Nombre:   nombre,
			Apellido: apellido,
		}
	}

	var listaDatos []*pb.Datos_DataNode

	log.Println("Personas en el mapa:")
	for id, datos := range personaMap {
		log.Printf("ID: %s, Nombre: %s, Apellido: %s", id, datos.Nombre, datos.Apellido)
	}

	log.Println("idList/ListaId: %s", idList.ListaId)
	for _, id := range idList.ListaId {
		log.Printf("Id siendo revisado: %s", id)
		if persona, ok := personaMap[id]; ok {
			log.Printf("Persona agregada: %s %s %s", id, persona.Nombre, persona.Apellido)
			listaDatos = append(listaDatos, persona)
		}
	}

	response := &pb.Lista_Datos_DataNode{
		ListaDatos_DataNode: listaDatos,
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Fallo en escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataNodeServer(s, &dataNodeServer{})

	fmt.Println("Servidor DataNode1 escuchando en :80")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}
}

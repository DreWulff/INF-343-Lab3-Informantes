package client

import (
	"Informantes/pb"
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
)

type register struct {
	planet    string
	city      string
	value     string // -1 is deleting
	clock     []int32
	direction string
}

func Sum(array []int32) int32 {
	var result int32
	result = 0
	for _, v := range array {
		result += v
	}
	return result
}

func GetCommand(command string) int32 {
	if command == "AddCity" {
		return 1
	} else if command == "UpdateName" {
		return 2
	} else if command == "UpdateNumber" {
		return 3
	} else if command == "DeleteCity" {
		return 4
	}
	return 0
}

func CommandConn(command int32, conn *grpc.ClientConn, split []string) *pb.Clock {
	var res *pb.Clock
	var err error
	if command == 1 {
		req := &pb.City{
			NombrePlaneta: split[1],
			NombreCiudad:  split[2],
			NuevoValor:    split[3],
		}
		client := pb.NewAddCityClient(conn)
		res, err = client.AddCity(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	} else if command == 2 {
		req := &pb.City{
			NombrePlaneta: split[1],
			NombreCiudad:  split[2],
			NuevoValor:    split[3],
		}
		client := pb.NewUpdateNameClient(conn)
		res, err = client.CambiarNombre(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	} else if command == 3 {
		req := &pb.City{
			NombrePlaneta: split[1],
			NombreCiudad:  split[2],
			NuevoValor:    split[3],
		}
		client := pb.NewUpdateNumberClient(conn)
		res, err = client.CambiarValor(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	} else if command == 4 {
		req := &pb.DelCt{
			NombrePlaneta: split[1],
			NombreCiudad:  split[2],
		}
		client := pb.NewDeleteCityClient(conn)
		res, err = client.DeleteCity(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	}
	conn.Close()
	return res
}

func StartInfo() {
	// Se declaran las variables
	var conn *grpc.ClientConn
	var err error
	var clock []int32
	var function int32
	val := " "
	logs := []register{}
	for {
		// Se ingresa un comando
		fmt.Println("Ingrese un comando: ")
		var command string
		fmt.Scanln(&command)
		split := strings.Split(command, " ")

		// Se verifica la validez del comando
		function = GetCommand(split[0])
		if function == 0 {
			fmt.Println("Comando invalido.")
			continue
		}

		// Conexion
		conn, err = grpc.Dial("localhost:5041", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client := pb.NewPetitionClient(conn)
		req := &pb.PetitionReq{
			Function: function,
			Planet:   split[1],
			City:     split[2],
		}
		res, err := client.Petition(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()

		// Se realiza peticion a direccion recibida
		conn, err = grpc.Dial(res.Direction, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		clock = CommandConn(function, conn, split).Reloj

		// Se verifica respuesta
		if Sum(clock) < 0 {
			fmt.Println("Respuesta del servidor: No se pudo realizar su peticion")
			continue
		}

		// Registro de comando realizado
		if function != 4 {
			val = split[3]
		} else {
			val = "-1"
		}
		reg := register{
			planet:    split[1],
			city:      split[2],
			value:     val,
			clock:     clock,
			direction: res.Direction,
		}
		logs = append(logs, reg)
		// De momento no se usan los logs registrados
	}
}

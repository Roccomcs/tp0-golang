package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	//log.Println("Soy un logcito")


	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("/home/utnso/Desktop/tp0-golang/client/config.json")
	log.Println(globals.ClientConfig.Ip)
	log.Println(globals.ClientConfig.Puerto)
	log.Println(globals.ClientConfig.Mensaje)


	// validar que la config este cargada correctamente

	// loggeamos el valor de la config

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip,globals.ClientConfig.Puerto,globals.ClientConfig.Clave)

	// leer de la consola el mensaje
	//utils.LeerConsola()	
	//utils.LeerConsolaHastaVacio()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}

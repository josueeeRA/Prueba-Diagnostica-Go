package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)
	fmt.Println("Venta registrada con exito")
}

func MostrarEstadisticas() {
	if len(productosVendidos) == 0 {
		fmt.Println("No existen ventas registradas")
		return
	}

	var total float64 = 0
	for i := 0; i < len(subtotales); i++ {
		total = total + subtotales[i]
	}
	fmt.Println("Total recaudado:", total)
}

func main() {
	var opcion int

	for {
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadisticas")
		fmt.Println("3. Salir")
		fmt.Print("Ingrese una opcion: ")
		fmt.Scan(&opcion)

		if opcion == 1 {
			fmt.Println("1. Arroz - $1.15")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan - $0.50")
		
		 var selecion int 
		 fmt.Println("Seleccione el numero del producto")
		 fmt.Scan(&selecion)
		}
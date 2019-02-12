package main

import ( 
	"strings"
	"fmt"
	"os"
	"bufio"
)
// entrada de dados via argumentos utilizando a biblioteca os.Args
// aceita argumentos em forma de slice(fatias) s[m:n]
// -- ":=" simbolo utilizado para criação de variaveis curtas (shot variable);
// variaveis com valores não declaradas tem seu valor default como 0 ou vazio 
// para numeros ou strings respectivamente;
func  main()  {
	var s , sep string
	for i := 1 ; i < len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
	//usingJoin()
	//usingFmt()
	//viewIndexAndValue()
	//Bufio()
	inputDataUser()
}
func usingJoin(){
	// string.join() é equivalmente ao algoritmo desenvolvido acima.
	fmt.Println(strings.Join(os.Args[1:], " "))
}
func usingFmt(){
	// exibe o comando que foi chamado durante a execução do os.Args;
	fmt.Println(os.Args[0])
}
func viewIndexAndValue(){
	// Função range produz um par de valores (indice, valor) do elemento;
	// neste exemplo de função: é criado duas variaveis locais (index e arg)
	// caso não seja necessario recuperar o valor do indice pode-se utilizar o 
	// underscore(_)
	// for _, arg := range os.Args[1:]
	s, sep := "", ""
	for index, arg := range os.Args[1:]{
		s += sep + arg
		sep = ""
		index += 1
		fmt.Println("index: " , index ,"value: " ,s)
	}
	
}
func inputDataUser(){
	// fmt.Scanln: Semelhante a entrada de dados do Scanner em java;
	// espera a entrada do usuario; para poder executar a proxima linha;
	var name, age string
	fmt.Println("say to name and Age")
	fmt.Scanln(&name,&age)
	fmt.Println("ok. Name" , name , " Age:" , age)
}
func Bufio(){
	input := bufio.NewScanner(os.Stdin)
	fmt.Println(input)
}

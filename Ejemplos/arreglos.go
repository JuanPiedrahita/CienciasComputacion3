package main
import "fmt"
func main(){

	var a [5]int
	fmt.Println("emp:", a)

	a[4]= 100
	fmt.Println("emp:", a)

	fmt.Println("Lon:", len(a))

	b:=[5] int{1,2,3,4,5}
	fmt.Println(b)

	var matriz [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			matriz[i][j] = i+j;
		}	
	}
	fmt.Println(matriz)
}

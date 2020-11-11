package main

import "fmt"

func main(){
	// var slice1 []int=make([]int,4)
	// slice1[0]=1
	// slice1[1]=2
	// slice1[2]=3
	// slice1[3]=4
	// for ix ,value:=range slice1 {
	// 	fmt.Printf("slice at %d is %d\n",ix,value)
	// }
	// example2()
	//
	// var screen [1][2]int
	// screens:=examplearr2(screen)
	// fmt.Println(screens,"screens")
	//

	//exam
	// fmt.Println(exam75())
    // f32:=[]float32{2.2,3.3,2.1,4.3}
	// fmt.Println(SumAndAverage(f32))
	fmt.Println(minSlice([]int{3,3,2}))
}
func example2(){
	//春天 夏天 秋天 冬天
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	} 
}
func exam75()[5]int{
	items := [...]int{10, 20, 30, 40, 50}
	for i, item := range items {
		items[i]= item*2
	}
	return items
}
func examplearr2(screen [1][2]int)[1][2]int{
	
	for row := range screen{
		for column := range screen[row] {
			screen[row][column] = 1
		}
	}
	return screen;
}
func examdian(arr []int){
	
}
func examsumarray(arrF []float32)float32{
	 var sum float32
	for _,v:=range arrF{
		sum+=v
	}
	return sum
}
func SumAndAverage(arrF []float32)(int, int, float32,float32){
	var sum float32
	var average float32
	// var sum1 int
	for _,v:=range arrF{
		sum+=v
	}
	average=sum/float32(len(arrF))
	var y int=int(sum)
	var x int=int(average)
	return y,x,sum,average
}

func minSlice(arr []int)int{
	arr0:=arr[0]
	for _,v:=range arr {
		if v<arr0{
			arr0=v
		}
	}
	return arr0
}
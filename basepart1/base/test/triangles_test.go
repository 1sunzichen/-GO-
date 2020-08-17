package main

import "testing"

func TestTriangles(t *testing.T){
	tests:=[]struct{
		a,b,c int
	}{
		//{3,4,5},
		//{12,5,13},
		{3000000000,4000000000,5000000000},
	}
	for _,tt:=range tests{
		if actual:=calctriangle(tt.a,tt.b);
		actual!=tt.c{
			t.Errorf("calcTrianele(%d,%d);"+
				"got %d; expected %d",
				tt.a,tt.b,actual,tt.c)
		}
	}
}
func BenchmarkSubstr(b *testing.B){
	//tests:=[]struct{
	//	a,b,c int
	//}{
	//	//{3,4,5},
	//	//{12,5,13},
	//	{3000000000,4000000000,5000000000},
	//}
	s:=[3] int{3,4,5}
	for i:=0;i<b.N;i++{


			if actual:=calctriangle(s[0],s[1]);
				actual!=s[2]{
				b.Errorf("calcTrianele(%d,%d);"+
					"got %d; expected %d",
					s[0],s[1],actual,s[2])
			}

	}

}
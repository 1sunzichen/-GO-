package main

import "testing"

func TestTriangles(t *testing.T){
	tests:=[]struct{
		a,b,c int
	}{
		{3,4,5},
		{12,5,13},
		{9,12,15},
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
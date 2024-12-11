package Calculator

import "testing"

func TestCalcultor(t *testing.T) {
	tests := []struct {
		a float64
		b float64
		operator string
		output      float64
	}{
		{5, 4, "+", 9},
		{5,4,"-",1},
		{5,4,"*",20},
		{5,4,"/",1.25}
	}
	for _,test:=range tests{
		res:=Calculator(test.a,test.b,test.operator)
		if res!=test.output{
			t.Errorf("test failed calculator(%f,%f,%s)=%f; %f",test.a,test.b,test.operator,res,test.output)

		}
	}
}

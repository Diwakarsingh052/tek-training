package sum

import "testing"

func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(x)

	if got != want {
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got) // test fails but continue to check others
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got) // stop the current test
	}

	//x = nil
	//got = SumInt(x)
	//want = 0
	//if got != want {
	//	t.Errorf("sum of nil should be %v; got %v", want, got)
	//}

	x = []int{1, -1}
	got = SumInt(x)
	want = 0
	if got != want {
		t.Errorf("sum of 1,-1 should be %v; got %v", want, got)
	}

}

package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	testRec := Rectangle{height: 10,width: 20}
	got := Perimeter(testRec)
	want := 60.0

	if got != want {
			t.Errorf("got %.2f want %.2f", got,want)
	}
}
		func TestTableTestArea(t *testing.T) {
			areaTests := []struct{
				name string
				shape Shape
				want float64

			}{
				{name:"Rectangle",shape:Rectangle{12, 6}, want:72.0},
				{name:"Circle", shape: Circle{10}, want: 314.1592653589793},
				{name: "Triangle",shape:Triangle{12, 6}, want:36.0},
			}
			
			for _, tt := range areaTests {
				t.Run(tt.name, func(t *testing.T){
				got := tt.shape.Area()
				if got != tt.want {
					t.Errorf("got %g want %g", got, tt.want)
				}
				})
			}
		}
 
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		want := 72.0

		checkArea(t,rectangle,want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		
		want := 314.1592653589793

		checkArea(t,circle,want)
	})

}
package forms

import "testing"
import "math"

func test_area(t *testing.T) {
	t.Run("Retangulo", func (t *testing.T)  {
		ret := Rectangle{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Error("A area diferente do esperado")
		}
	})

	t.Run("Circulo", func (t *testing.T)  {
		circ := Circle{ 10 }

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Error("A area diferente do esperado")
		}
	})
}
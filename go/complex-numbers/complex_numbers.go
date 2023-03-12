package complexnumbers

import "math"

type Number struct {
	r, i float64
}

func (n Number) Real() float64 {
	return n.r
}

func (n Number) Imaginary() float64 {
	return n.i
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.r + n2.r, n1.i + n2.i}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.r - n2.r, n1.i - n2.i}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.r*n2.r - n1.i*n2.i, n1.i*n2.r + n1.r*n2.i}
}

func (n Number) Times(factor float64) Number {
	return Number{n.r * factor, n.i * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		(n1.r*n2.r + n1.i*n2.i) / (n2.r*n2.r + n2.i*n2.i),
		(n1.i*n2.r - n1.r*n2.i) / (n2.r*n2.r + n2.i*n2.i),
	}
}

func (n Number) Conjugate() Number {
	return Number{n.r, -n.i}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.r*n.r + n.i*n.i)
}

func (n Number) Exp() Number {
	// e^(a + i * b) = e ^ a * e^(i * b)
	// e^(i * b) = cos(b) + i * sin(b)
	// e^a * cos(b) + e^a * sin(b) * i
	factor := math.Exp(n.r)
	return Number{
		factor * math.Cos(n.i),
		factor * math.Sin(n.i),
	}
}

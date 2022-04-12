package vectors

import 

type Vector []numeric

func (v Vector) Add(v2 Vector) Vector {
	var out Vector
	for i, s := range v {
		out = append(out, s+v2[i])
	}
	return out
}

func (v Vector) Sub(v2 Vector) Vector {
	var out Vector
	for i, s := range v {
		out = append(out, s-v2[i])
	}
	return out
}

func (v Vector) Dot(v2 Vector) float64 {
	var out float64
	for i, s := range v {
		out += s * v2[i]
	}
	return out
}

func (v Vector) Scale(scalar float64) Vector {
	var out Vector
	for _, s := range v {
		out = append(out, s*scalar)
	}
	return out
}

type Matrix []Vector

func (m Matrix) Add(m2 Matrix) Matrix {
	var out Matrix
	for i, v := range m {
		out = append(out, v.Add(m2[i]))
	}
	return out
}

func (m Matrix) Sub(m2 Matrix) Matrix {
	var out Matrix
	for i, v := range m {
		out = append(out, v.Sub(m2[i]))
	}
	return out
}

func (m Matrix) Transform(v Vector) Vector {
	var out Vector
	for i, vec := range m {
		out = out.Add(vec.Scale(v[i]))
	}
	return out
}

func (m Matrix) Mutliply(m2 Matrix) Matrix {
	var out Matrix
	for _, v := range m2 {
		out = append(out, m.Transform(v))
	}
	return out
}

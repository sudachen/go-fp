package fu

import "reflect"

/*
Min returns minimal value
*/
func Min(a ...interface{}) interface{} {
	return a[MinIndex(reflect.ValueOf(a))]
}

/*
MaxValue returns maximal value
*/
func MinValue(a reflect.Value) reflect.Value {
	return a.Index(MinIndex(a))
}

/*
MinIndex returns index of minimal value
*/
func MinIndex(a reflect.Value) int {
	if a.Kind() != reflect.Slice {
		panic("only slice is allowed as an argument")
	}
	N := a.Len()
	d := 0
	r := a.Index(0)
	for i := 1; i < N; i++ {
		j := a.Index(i)
		if Less(j, r) {
			r = j
			d = i
		}
	}
	return d
}

/*
Max returns maximal value
*/
func Max(a ...interface{}) interface{} {
	return a[MaxIndex(reflect.ValueOf(a))]
}

/*
MaxValue returns maximal value
*/
func MaxValue(a reflect.Value) reflect.Value {
	return a.Index(MaxIndex(a))
}

/*
MaxIndex returns index of maximal value
*/
func MaxIndex(a reflect.Value) int {
	if a.Kind() != reflect.Slice {
		panic("only slice is allowed as an argument")
	}
	N := a.Len()
	d := 0
	r := a.Index(0)
	for i := 1; i < N; i++ {
		j := a.Index(i)
		if Less(r, j) {
			r = j
			d = i
		}
	}
	return d
}

/*
IndexOfMin returns index of minimal value
*/
func IndexOfMin(a interface{}) int {
	v := reflect.ValueOf(a)
	return MinIndex(v)
}

/*
Mini returns minimal int value
*/
func Mini(a ...int) int {
	q := a[0]
	for _, x := range a[1:] {
		if x < q {
			q = x
		}
	}
	return q
}

/*
Minf returns minimal float value
*/
func Minf(a ...float32) float32 {
	q := a[0]
	for _, x := range a[1:] {
		if x < q {
			q = x
		}
	}
	return q
}

/*
IndexOfMax returns index of maximal value
*/
func IndexOfMax(a interface{}) int {
	v := reflect.ValueOf(a)
	return MaxIndex(v)
}

/*
Maxi returns maximal int value
*/
func Maxi(a ...int) int {
	q := a[0]
	for _, x := range a[1:] {
		if x > q {
			q = x
		}
	}
	return q
}

/*
Maxf returns maximal float32 value
*/
func Maxf(a ...float32) float32 {
	q := a[0]
	for _, x := range a[1:] {
		if x > q {
			q = x
		}
	}
	return q
}

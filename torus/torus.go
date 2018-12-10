//Package torus implements two dimensional fixed sized torus
package torus

import (
	"fmt"
	"reflect"
	"strings"
)

//Torus represent a two dimensional torus using a matrix
type Torus struct {
	tp         reflect.Type
	cols, rows int
	data       [][]interface{}
}

//New allocates and retuns a torus with r rows and c columns, and v in every index
func New(v interface{}, c, r int) *Torus {
	var torus Torus
	torus.cols, torus.rows = c, r
	torus.tp = reflect.TypeOf(v)
	torus.data = make([][]interface{}, r)
	for i := range torus.data {
		torus.data[i] = make([]interface{}, c)
		for j := range torus.data[0] {
			torus.data[i][j] = v
		}
	}
	return &torus
}

//Cols returns the columns len of t
func (t *Torus) Cols() int {
	return t.cols
}

//Rows returns the rows len of t
func (t *Torus) Rows() int {
	return t.rows
}

//Type returns the type of the elements in the torus t
func (t *Torus) Type() reflect.Type {
	return t.tp
}

//Init initialize the torus with v in every index,
//returns error if the type of v is diferent to t.Type()
func (t *Torus) Init(v interface{}) error {
	t1 := reflect.TypeOf(v)
	if t1 != t.tp {
		return fmt.Errorf("Diferent types %v and %v", t.tp, t1)
	}
	for i := range t.data {
		for j := range t.data[0] {
			t.data[i][j] = v
		}
	}
	return nil
}

//Get return the element in the ith row and jth column
func (t *Torus) Get(i, j int) interface{} {
	i %= t.rows
	j %= t.cols
	if i < 0 {
		i += t.rows
	}
	if j < 0 {
		j += t.cols
	}
	return t.data[i][j]
}

//Set sets the element in the ith row and jth column to v value,
//returns error if the type of v diferent to t.Type()
func (t *Torus) Set(i, j int, v interface{}) error {
	t1 := reflect.TypeOf(v)
	if t1 != t.tp {
		return fmt.Errorf("Diferent types %v and %v", t.tp, t1)
	}
	i %= t.rows
	j %= t.cols
	if i < 0 {
		i += t.rows
	}
	if j < 0 {
		j += t.cols
	}
	t.data[i][j] = v
	return nil
}

func (t *Torus) String() string {
	var sb strings.Builder
	for _, r := range t.data {
		sb.WriteString(fmt.Sprint(r, "\n"))
	}
	return sb.String()
}

package autofillDB

import (
	"strconv"
	"os"
	"reflect"
	"strings"

)
// csvdata complements the csv package by allowing you to map a custom structure to
// the columns of data in a CSV file. The struct needs to be annotated so that each
// field can match a column in the data
//
//    type Person struct {
//       FirstName string `field:"First Name"`
//       Second_Name string
//       Age int
//    }
//
// The name of the column can be inferred from the field name; any underscores in the
// name are converted to spaces when comparing. Otherwise, you must provide a tag
// 'field' with the name of the column.
//
//     r := csv.NewReader(os.Stdin)
//     p := new (Person)
//     rs,_ := NewReaderIter(r,p)
//     for rs.Get() {
//        fmt.Println(p.FirstName,p.Second_Name,p.Age)
//     }
//     if rs.Error != nil {
//        fmt.Println("error",rs.Error)
//    }

// The data source is any object that has a Read method which can
// return a row as a slice of strings. This matches csv.Reader in particular.
type Reader interface {
	Read() ([]string, os.PathError)
}

// Custom data types can be implemented by implementing Value; these
// methods must be defined on a pointer receiver.
// The interface is also used by flag package for a similar purpose.
type Value interface {
	String() string
	Set(string) bool
}

// ReadIter encapsulates an iterator over a Reader source that fills a
// pointer to a user struct with data.
type ReadIter struct {
	Reader       Reader
	Headers      []string
	Error        os.Error
	Line, Column int
	fields       []reflect.Value
	kinds        []int
	tags         []int
}

const (
	none_k = iota
	string_k
	int_k
	float_k
	uint_k
	value_k
)

// Creates a new iterator from a Reader source and a user-defined struct.
func NewReadIter(rdr Reader, ps interface{}) (this *ReadIter, err os.Error) {
	this = new(ReadIter)
	this.Line = 1
	this.Headers, err = rdr.Read()
	this.Reader = rdr
	if err != nil {
		this = nil
		return
	}
	st := reflect.TypeOf(ps).Elem()
	sv := reflect.ValueOf(ps).Elem()
	nf := st.NumField()
	this.kinds = make([]int, nf)
	this.tags = make([]int, nf)
	this.fields = make([]reflect.Value, nf)
	for i := 0; i < nf; i++ {
		f := st.Field(i)

		val := sv.Field(i)
		// get the corresponding field name and look it up in the headers
		tag := f.Tag.Get("field")
		if len(tag) == 0 {
			tag = f.Name
			if strings.Contains(tag, "_") {
				tag = strings.Replace(tag,"_"," ",-1)
			}
		}
		itag := -1
		for k, h := range this.Headers {
			if h == tag {
				itag = k
				break
			}
		}
		if itag == -1 {
			err = os.NewError("cannot find this field " + tag)
			this = nil
			return
		}
		kind := none_k
		Kind := f.Type.Kind()
		// this is necessary because Kind can't tell distinguish between a primitive type
		// and a type derived from it. We're looking for a Value interface defined on
		// the pointer to this value
		_, ok := val.Addr().Interface().(Value)
		if ok {
			val = val.Addr()
			kind = value_k
		} else {
			switch Kind {
			case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
				kind = int_k
			case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
				kind = uint_k
			case reflect.Float32, reflect.Float64:
				kind = float_k
			case reflect.String:
				kind = string_k
			default:
				kind = value_k
				_, ok := val.Interface().(Value)
				if !ok {
					err = os.NewError("cannot convert this type ")
					this = nil
					return
				}
			}
		}
		this.kinds[i] = kind
		this.tags[i] = itag
		this.fields[i] = val
	}
	return
}

// The Get method reads the next row. If there was an error or EOF, it
// will return false.  Client code must then check that ReadIter.Error is
// not nil to distinguish between normal EOF and specific errors.
func (this *ReadIter) Get() bool {
	row, err := this.Reader.Read()
	this.Line = this.Line + 1
	if err != nil {
		if err != os.EOF {
			this.Error = err
		}
		return false
	}
	var ival int64
	var fval float64
	var uval uint64
	var v Value
	var ok bool

	for fi, ci := range this.tags {
		vals := row[ci] // string at column ci of current row
		f := this.fields[fi]
		switch this.kinds[fi] {
		case string_k:
			f.SetString(vals)
		case int_k:
			ival, err = strconv.Atoi64(vals)
			f.SetInt(ival)
		case uint_k:
			uval, err = strconv.Atoui64(vals)
			f.SetUint(uval)
		case float_k:
			fval, err = strconv.Atof64(vals)
			f.SetFloat(fval)
		case value_k:
			v, ok = f.Interface().(Value)
			if !ok {
				err = os.NewError("Not a Value object")
				break
			}
			v.Set(vals)
		}
		if err != nil {
			this.Column = ci + 1
			this.Error = err
			return false
		}
	}
	return true
}

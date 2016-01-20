package mysql
import (
	"reflect"
	"fmt"
)


func InsertRecord(tableName string, i interface{}) string  {
	s := reflect.ValueOf(i).Elem()
	typeOfT := s.Type()
	values := ""
	fields := ""
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fields += typeOfT.Field(i).Tag.Get("mysql")
		values += f.Interface()
		//retS += fmt.Sprintf("\n%d: %s %s = %v %s", i,
		//	typeOfT.Field(i).Name, f.Type(), f.Interface(), typeOfT.Field(i).Tag.Get("mysql"))
	}
	return fmt.Sprintf("insert into %s (%s) values (%s)", )
}
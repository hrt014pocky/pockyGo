/*
 * @Date: 2022-01-27 17:45:22
 * @LastEditors: Szang
 * @LastEditTime: 2022-01-27 18:03:06
 * @FilePath: \hrt014pocky\encoder\main.go
 */

package main
package rlp

import "reflect"


type writer func(reflect.Value, *encbuf) error
var theTC = newTypeCache()
func cachedWriter(typ reflect.Type) (writer, error) {
	info := theTC.info(typ)
	return info.writer, info.writerErr
}


func (w *encbuf) encode(val interface{}) error {
	rval := reflect.ValueOf(val)
	writer, err := cachedWriter(rval.Type())
	if err != nil {
		return err
	}
	return writer(rval, w)
}

func main() {

}

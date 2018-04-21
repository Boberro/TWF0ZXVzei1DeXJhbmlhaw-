package main

type MyObject struct {
	Key         string
	Value       []uint8
	ContentType string
}

var myObjects = []*MyObject{
	{"abc", []uint8("Tralalala"), "text/plain"},
}

func getObject(list []*MyObject, key string) (int, *MyObject) {
	for i, obj := range list {
		if obj.Key == key {
			return i, obj
		}
	}
	return -1, nil
}

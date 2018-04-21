package main

type MyObject struct {
	Key   string
	Value string
}

var myObjects = []*MyObject{
	{Key: "abc", Value: "123"},
}

func getObject(list []*MyObject, key string) (int, *MyObject) {
	for i, obj := range list {
		if obj.Key == key {
			return i, obj
		}
	}
	return -1, nil
}

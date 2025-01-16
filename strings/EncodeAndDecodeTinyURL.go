package main

import (
	"fmt"
	"strconv"
)

type Codec struct {
	encodedMap map[string]string
	decodedMap map[string]string
	counter    int
}

func Constructor() Codec {
	return Codec{encodedMap: map[string]string{}, decodedMap: make(map[string]string), counter: 0}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if val, ok := this.encodedMap[longUrl]; ok {
		return val
	}
	this.counter++
	tinyUrl := "http://tinyurl.com/" + strconv.Itoa(this.counter)
	this.encodedMap[longUrl] = tinyUrl
	this.decodedMap[tinyUrl] = longUrl
	return tinyUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.decodedMap[shortUrl]
}

func main() {
	obj := Constructor()
	url := obj.encode("/ravindra")
	ans := obj.decode(url)
	fmt.Println(ans)
	fmt.Println(obj)
}

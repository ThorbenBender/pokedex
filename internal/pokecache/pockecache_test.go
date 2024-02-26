package pokecache

import (
	"testing"
	"time"
)


func TestCreateCache(t *testing.T) {
  interval := time.Millisecond * 10
  cache := NewCache(interval)
  if cache.cache == nil {
    t.Error("cache is nil")
  }

  cases := []struct {
    inputKey string
    inputVal []byte
  }{
    {
      inputKey: "key1",
      inputVal: []byte("val1"),
    },
  }


  for _, cas := range cases {

    cache.Add(cas.inputKey, cas.inputVal)

    actual, ok := cache.Get(cas.inputKey)

    if !ok {
      t.Error("key1 not found")
      continue
    }

    if string(actual) != string(cas.inputVal) {
      t.Errorf("%s does not match %s", string(actual), string(cas.inputVal))
      continue
    }
  }
}


func TestReap(t *testing.T) {
  interval := time.Millisecond * 10
  cache := NewCache(interval)

  keyOne := "key1"
  cache.Add(keyOne, []byte("val1"))

  time.Sleep(interval + time.Millisecond)

  _, ok := cache.Get(keyOne)

  if ok {
    t.Errorf("%s should have been reaped", keyOne)
  }
}



func TestReapFail(t *testing.T) {
  interval := time.Millisecond * 10
  cache := NewCache(interval)

  keyOne := "key1"
  cache.Add(keyOne, []byte("val1"))

  time.Sleep(interval / 2)

  _, ok := cache.Get(keyOne)

  if !ok {
    t.Errorf("%s should have not been reaped", keyOne)
  }
}

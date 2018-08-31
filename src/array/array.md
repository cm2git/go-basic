#### *GO Array and Slice*
```
    Go array is value type,the length n and type
    T compose the array type
```
1. *definition*
```
    [n]T
    n: the length of array
    T: type of array
```
2. *create array*
```
    a:=[]int{1,2,3}
    <=>
    a:=[3]int{1,2,3}
```
3. *loop array*
```
    for i,v:=range a{
        fmt.Printf("index %d ,value %v",i,v)
    }
```


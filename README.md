# comparejson

### Context
After implementing database backups, we need to ensure that the data is restored correctly,
and no corruption happens.

To verify equality, we compare the exported data from before and after the backup.

Exporting data from the DB produces a list of complex JSON objects. The objects and keys
inside can be sorted randomly due to the specificity of the database.

### Task
Write a small program in Go that accepts two JSON files and prints out if they are equal.


### Solution

[comparejson.go/EqualDataSets](https://github.com/JFixby/comparejson/blob/main/comparejson.go#L15)

Usage:
[usageexample_test.go](https://github.com/JFixby/comparejson/blob/main/usageexample_test.go)

```Go
entriesSet1, _ := ParseJsonDataSet(json_string_1)
entriesSet2, _ := ParseJsonDataSet(json_string_2)

result := EqualDataSets(entriesSet1, entriesSet2)
```


 


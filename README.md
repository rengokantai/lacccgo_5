####lacccgo_5
recursion, file manipulation

######1
resource
[the way to Go](https://sites.google.com/site/thewaytogo2012/Downhome/Topic3)

######5
use third party
```
go get github.com/rwcarlsen/goexif/exif  //download
go get -u github.com/rwcarlsen/goexif/exif  //update
```
######11
finished 11

######15
```
f.err:=os.Openfile(path,os.O_APPEND | os.O_WRONLY,0600)
if _,err=f.WriteString(text);err!=nil{}
```
######16
```
strconv.Itoa(int)
os.Rename(old,new)
files,_:=ioutil.ReadDir(dir)
for i,f:=range files{
fmt.Print(f.Name())
}
```
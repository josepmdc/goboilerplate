## How to use

Run the following command to replace goboilerplate with the name of your
project: 
```
find . -type f -print0 | xargs -0 sed -i 's/goboilerplate/<name of your project>/g'
```

Create the database

```
createdb <name of your project>
```

## HOW TO RUN API 

```bash
# run in docker container
make up 
make migrate_up

-------------------------

# run in localhost 
make run 
# change port of database in Makefile
make migrate_up  
# or 
# execute sql scripts of migrations dir in your postgres database
# and change postgres host in config.yml (database -> localhost) 
```



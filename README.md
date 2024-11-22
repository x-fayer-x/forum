# Forum

this is a program for a forum write in golang with a database

## forum folder

there is 9 file 

**AddFunction** is a file which contain all function which add somethings in database

**database** is the file which contain all function which create table in database 

**GetInDb** is the file which contain all function which take somethings from the data base

**HandlerCon** is the file which contain all Handler for the server when user is connected

**HandlerNoCon** is the file which contain all Handler for the server when user is not connected

**PrimaryFunction** is the file which contain all function for take all data each handler need

**Structure** is the file which contain all structure for the handlers 

**subfunction** is the file which contain all function which delete somethings in database

**updatefunction** is the file which contain all function wich update somethings in database

-----------Stress test---------------------
// effectue plusieurs requete simultanement pour voir si on gere le DDOS
for i in {1..20}; do curl -s -o /dev/null -w "%{http_code}\n" http://localhost:8080/; done
ou
siege -b -t 10s http://127.0.0.1:8080
ou 
// avec apache2
ab -n 100 -c 10 http://localhost:8080/
------------verifier luuid de mon user------------
run "sqlite3 forum.db" and rafter un "SELECT * FROM users;" to select all users.
```
--- api/todos ‹master* M› » 
--- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/detail -d '{"id": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/edit -d '{"id": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/del -d '{"id": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/get -d '{"id": "1000"}'
{"id":"","code":0,"detail":"rpc: can't find method Get","status":""}%                                                                                                                                                                                                           --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{"id": "1000"}'
{"success":true}%                                                                                                                                                                                                                                                               --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/del -d '{"id": "1000"}' 
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/del -d '{"idd": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/del -d '{"iddd": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/detail -d '{"iddd": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/edit -d '{"iddd": "1000"}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/edit -d '}'               
{"id":"go.micro.api.todos.todos.Edit","code":500,"detail":"body json error","status":"Internal Server Error"}%                                                                                                                                                                  --- api/todos ‹master* M› » 
--- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/edit -d '' 
{"id":"go.micro.api.todos.todos.Edit","code":500,"detail":"body json error","status":"Internal Server Error"}%                                                                                                                                                                  --- api/todos ‹master* M› » 
--- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/edit -d '{}'
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master* M› » 
--- api/todos ‹master* M› » 
--- api/todos ‹master* M› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/add -d '{"title": "item", "content":"sssss"}'
{"success":true}%                                                                                                                                                                                                                                                               --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{}'                                 
{"success":true,"todos":[{"id":"1014","title":"item","content":"sssss"}]}%                                                                                                                                                                                                      --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/edit -d '{"title":"item2222", "content":"sssssssss", "id":"1014"}'
{"success":true}%                                                                                                                                                                                                                                                               --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{}'                                                      
{"success":true,"todos":[{"id":"1014","title":"item2222","content":"sssssssss"}]}%                                                                                                                                                                                              --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/add -d '{"title": "item", "content":"sssss"}'                     
{"success":true}%                                                                                                                                                                                                                                                               --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{}'                                                      
{"success":true,"todos":[{"id":"1014","title":"item2222","content":"sssssssss"},{"id":"1015","title":"item","content":"sssss"}]}%                                                                                                                                               --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/detail -d '{"iddd": "1015"}'                 
{"error":{"code":404,"detail":"not found"}}%                                                                                                                                                                                                                                    --- api/todos ‹master› » 
--- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/detail -d '{"id": "1015"}' 
{"success":true,"todo":{"id":"1015","title":"item","content":"sssss"}}%                                                                                                                                                                                                         --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/del -d '{"id": "1015"}'
{"success":true}%                                                                                                                                                                                                                                                               --- api/todos ‹master› » curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{}'                
{"success":true,"todos":[{"id":"1014","title":"item2222","content":"sssssssss"}]}%                                                                                                                                                                                              --- api/todos ‹master› » 

```
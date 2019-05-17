links 应该可以保证启动的依赖关系




docker-compose up -d --scale todos-api=3  
docker-compose up -d --scale todos-srv=3



docker-compose scale todos-srv=10
docker-compose scale todos-api=10

curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{"title":"item123", "content":"asaaaaaaaa"}' | jq .todos



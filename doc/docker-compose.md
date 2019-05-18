links 应该可以保证启动的依赖关系




docker-compose up -d --scale todos-api=3  
docker-compose up -d --scale todos-srv=3



docker-compose scale todos-srv=10
docker-compose scale todos-api=10

curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{"title":"item123", "content":"asaaaaaaaa"}' | jq .todos






docker-compose scale之后，docker进程会关掉重新建立，直接关掉的docker进程，导致没有向consul取消服务，导致有时路由到关闭的docker，导致访问失败。
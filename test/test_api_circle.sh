#!/usr/bin/env bash
i=0
times=50
while  (( $i <= $times ));
do
#    curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/add -d '{"title":"item123", "content":"asaaaaaaaa"}'
    curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{"title":"item123", "content":"asaaaaaaaa"}'
    let i++
    echo '\n'
    echo $i
    echo '\n'
done;

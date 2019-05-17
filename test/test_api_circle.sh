#!/usr/bin/env bash
i=0
times=50
while  (( $i <= $times ));
do
(
    #    curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/add -d '{"title":"item123", "content":"asaaaaaaaa"}'
    curl -XPOST -H 'Content-Type: application/json'  http://localhost:8080/todos/todo/list -d '{}'

    echo '\n'
    echo $i
    echo '\n'
)& # 小括号里面的会fork一个子进程执行
let i++
echo $!
done;
wait;

```bash
10332  hub clone https://github.com/mingz2013/demo-todos-go-micro
10333  cd demo-todos-go-micro
10334  mkdir doc
10335  mkdir srv api web fnc
10336  touch doc/create.md doc/doc.md
10337  touch Makefile
10338  cd srv 
10339  micro new todos --gopath=false --type="srv"
10340  cd ..
10341  cd api
10342  micro new todos --gopath=false --type="api"
10343  cd ..
10344  cd web
10345  micro new todos --gopath=false --type="web"
```
# HelloWorld

An Hello World application in Go with *negroni*.

## Local setup
Install and run locally on port 3000
```sh
sh install.sh && go run src/index.go
```

Remember to define your ``$GOPATH``.

## Heroku setup
First create your application using the Heroku Dashboard, and follow the guide for setting up the project using the Heroku Toolkit. E.g:

```sh
heroku git:remote -a <heroku project name>
```

Then add the build pack for running go applications
```sh
heroku create -b https://github.com/heroku/heroku-buildpack-go.git
```

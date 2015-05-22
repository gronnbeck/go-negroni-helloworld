# HelloWorld

An Hello World application in Go with *negroni*.

## Local setup
To install and run locally the project locally, you need to install [godep](https://github.com/tools/godep) first. Then run ``godep restore`` to restore the go dependencies. 

## Heroku setup
First create your application using the Heroku Dashboard, and follow the guide for setting up the project using the Heroku Toolkit. E.g:

```sh
heroku git:remote -a <heroku project name>
```

Then add the build pack for running go applications
```sh
heroku create -b https://github.com/heroku/heroku-buildpack-go.git
```

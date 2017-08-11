# Binar Academy - Heroku API
by: Yogi Agnia Dwi Saputro

## Penjelasan
Ini adalah source code Heroku API untuk asessment test Binar Academy kelas back end.
 - Package models:
      - model.go
      - user.go
 - Package routers
      - router.go
 - Package controllers
      - subscribe.go
      - user.go
 - Package main
      - main.go

## Note
Aplikasi yang ada di source code sesuai kriteria soal berbeda dengan aplikasi yang terdeploy di Heroku.
Yang sudah berhasil dilakukan adalah deploy aplikasi mengikuti petunjuk Getting Started with Go.
Deploy aplikasi sesuai assessment test tidak dapat dilakukan karena berbagai issue yang muncul dari sisi Go maupun deploy di Heroku.

Berikut adalah isi readme.md dari Heroku: Getting Started with Go
-----------------------------------------------------------------
-----------------------------------------------------------------
-----------------------------------------------------------------
# go-getting-started

A barebones Go app, which can easily be deployed to Heroku.

This application supports the [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) article - check it out.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/heroku/go-getting-started
$ cd $GOPATH/src/github.com/heroku/go-getting-started
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

You should also install [govendor](https://github.com/kardianos/govendor) if you are going to add any dependencies to the sample app.

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)

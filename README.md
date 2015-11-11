
# Go sample application

See [Goに入門してRedis+PostgresなアプリをHerokuにデプロイするまで](http://leko.jp/archives/763)

## Usage
Install [Vagrant](https://www.vagrantup.com/downloads.html), [Ansible](http://docs.ansible.com/intro_installation.html) and [Bower](http://bower.io/#install-bower)

```
$ git clone git@github.com:Leko/godemo.git
$ cd godemo
$ bower install
$ vagrant up --provision
$ vagrant ssh

$ vagrant@precise64:~$ cd go/src/godemo
$ vagrant@precise64:~/go/src/godemo$ go get
$ vagrant@precise64:~/go/src/godemo$ go get github.com/codegangsta/gin
$ vagrant@precise64:~/go/src/godemo$ gin -p 8080
```

Then open `http://localhost:3000`

## Deploy

```sh
$ rm -rf .git
$ git init
$ git add .
$ git commit -m "initial commit"
$ heroku apps:create --addons heroku-postgresql:hobby-dev,rediscloud:30 --buildpack heroku/go
$ git push heroku master

$ # Wait a minutes
$ heroku open
```

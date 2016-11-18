# Golang Api Command
golang API with command line

Requirements
------------

* docker-engine
* docker-compose

Commands
---------

Build and Up the application
~~~
$ docker-compose build
$ docker-compose up -d
~~~

Run the API
~~~
$ docker exec golang_api app --port 4000 (or default 9000)
~~~

Other commands
~~~
$ docker exec golang_api app -h (optional)
$ docker exec golang_api app -v (optional)
~~~

Stop the container
~~~
$ docker-compose stop
~~~
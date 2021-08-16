# TodoApp

Todo is a simple Go webapp to manage todos, I built this tool to manage tasks which I wanted to do, there are many good to-do lists, but also I wanted to learn the Go webapp development.

How to use?
==================
This will generate the binary and set up the database. If you want, you can copy the binary and the public folder into a folder of your choice.

Manually:

1. If you desire to run the application in local, change the .env file according the docker's postgresql credentials. 
2. `docker-compose up -d`
3. open [localhost:8080](http://localhost:8080)
4. Production url [http://34.116.133.134:8080/ping](http://34.116.133.134:8080/ping)

## Decisions

1. Chose the Go language due to it is used in the company.
2. Chose the Gin among dozens of frameworks because of Faster performance, middleware support, json validation, routes and popularization. 
3. Chose the MVC pattern which used in other modern frameworks to better understanding of the project.
4. Chose GCP because of that free 300$ :)
5. Chose The github action to able to learn it.

## Features

1. Add, list, delete a task.
2. Migration runs automatically when the app is bootstrapped.
3. If a commit pushes into main branch or a branch merges with the main branch, the github action works to test, build and deploy the application to the Google Cloud Kubernetes env.

# License

The MIT License (MIT)

Copyright (c) 2015 Suraj Patil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
# Docker Excercise One

This is a simple challenge that involves publishing ports when running a Docker container.

You will **NOT** need to modify `main.go`

You will **NOT** need to modify the `Dockerfile`

## What do I do?

Build this directory into a docker container. Remember to build the go binary for Linux first!

When running your docker container you will need to:

- Publish port `4000` on the container to port `4000` on you machine

Take another look at [Dr. Stearn's Docker tutorial](https://drstearns.github.io/tutorials/docker/) if you need hints!

## Test it out

Open a browser window and go to [localhost:4000](http://localhost:4000). If everything went well you should see a success message. If you aren't able to connect then you didn't publish the ports quite right.

## Stop and remove the container

You likely don't need or want this container to be running any more so let's stop and remove it.

Run `docker ps` to list all currently running containers and copy the `CONTAINER ID` of the docker exercise container.

Stop the conainer:

`docker stop <container-id>`

Remove the container:

`docker rm <container-id>`

*When you run `docker stop` and `docker rm` you will see the container ID printed to your terminal window. You can treat this as a success message.*

You can also stop and remove a running container in one command:

`docker rm -f <container-id>`

This forces the removal of the container, even if it is currently running.

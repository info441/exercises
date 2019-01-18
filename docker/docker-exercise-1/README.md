# Solution

You do not need to modify either the go source code, or the Dockerfile.

```
# build the go executable for linux
GOOS=linux go build

# build your docker container. Make sure to replace "brendankellogg" with your Dockerhub name!
docker build -t brendankellogg/docker-exercise-1 .

# clean up the built Go executable
go clean

# run your newly built container. Make sure to replace "brendankellogg" with your Dockerhub name!
docker run -d -p 4000:4000 brendankellogg/docker-exercise-1

# You should now be able to see the container on http://localhost:4000/
```

## Things To Note

The go executable is hardcoded to listen on `:4000`, so that is why you are able to connect to the container on port `4000`.

You can remove the built executable (with `go clean`) after the container is built since the executable has already been copied into the container.

`-d` in the `docker run` means to run the container in "detached" mode. This means that the container will be run as a background process, and will not consume your terminal window.

`-t` in the `docker build` command (with no extra arguments) tags the container your building as the `latest` version.

`docker build` requires a directory to build the container from as its last argument. This is typically the current directory, so you can use `.` as the directory.

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

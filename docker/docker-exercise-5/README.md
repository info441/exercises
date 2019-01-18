# Solution

You do not need to modify the go source code, but you will need to modify the `Dockerfile`.

```
# build the go executable for linux
GOOS=linux go build

# set up your Dockerfile! The Dockerfile on this branch is correct!

# build your docker container. Make sure to replace "brendankellogg" with your Dockerhub name!
docker build -t brendankellogg/docker-exercise-5 .

# clean up the built Go executable
go clean

# run your newly built container. Make sure to replace "brendankellogg" with your Dockerhub name!
docker run -d \
-p 4000:4000 \
-e PORT=4000 \
-e STATICDIR=/static/ \
brendankellogg/docker-exercise-5

# You should now be able to see the container on http://localhost:4000/
```

## Things To Note

We do not need to specify a specific file with the `STATICDIR` environment variable since we are serving the entre directory

We do not need to volume mount the static directory since we added the static directory with the Dockerfile. However, you can do it this way if you'd like. Remove the COPY command from the Dockerfile that adds the static directory in and add a volume mount to the run command above that allows the container to see the static directory at the location specified by the STATICDIR environment variable in the run command above.

`STATICDIR` expects an absolute filepath (including the file name) to the location of the file within the context of the container. That is, the filepath provided should be where that file lives inside of the container and NOT where it is on your host machine.

# Docker Excercise Five

This is a simple challenge that involves publishing ports and setting environment variables, as well as mounting volumes and creating a Dockerfile when building and running a Docker container. You also need to write the go program.

You **WILL** need to modify `main.go`

You **WILL** need to modify the `Dockerfile`

## What do I do?

Create a go webserver that serves files out of the `static` diretory within this exercise. A piece of code needed to do this has been given to you in `main.go`. Remember to publish the ports you need, the environment variables you need, and to mount the correct volume into your container.

You need to:

- Add the correct server code to `main.go`
- Configure the `Dockerfile` with the correct build instructions
- Run your container with all the necessary flags

*You should not be hardcoding values for your webserver, instead read them in from environment variables*

Take another look at [Dr. Stearn's Docker tutorial](https://drstearns.github.io/tutorials/docker/) if you need hints!

Feel free to modify the contents of the `static` directory!

## Test it out

Open a browser window and go to `localhost:<port>/index.html`. If everything went well you should see an HTML page. If you aren't able to connect then you there is an issue with your ports!

If your container crashed, use `docker logs <container-id>` to see what the issue is and fix it!

*If you named your container, you will need to remove it before you can start another with the same name.*

## Stop and remove the container

You likely don't need or want this container to be running any more so let's stop and remove it.

Run `docker ps` to list all currently running containers and copy the `CONTAINER ID` of the docker exercise container.

Stop the conainer:

`docker stop <container-id>` or `docker stop <container-name>`

Remove the container:

`docker rm <container-id>` `docker rm <container-name>`

*When you run `docker stop` and `docker rm` you will see the container ID printed to your terminal window. You can treat this as a success message.*

Remember that you can also stop and remove a running container in one command:

`docker rm -f <container-id>` or `docker rm -f <container-name>`

This forces the removal of the container, even if it is currently running.

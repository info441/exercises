# Put your docker build instructions here!

# use the scratch image since we do not need any linux utilities
FROM scratch

# copy the executable that we'll be running into the container
COPY docker-exercise-4 docker-exercise-4

# tell docker to use the executable as the main entrypoint when the
# container starts. That is, what should the container do when it starts?
ENTRYPOINT ["/docker-exercise-4"]
# Go k8s cli

A simple k8s cli written in Golang.

## Requirements

* For the Cli the Kubeconf must be working.
* For the API it image must be avaliable on kubernetes.

## Assumptions

### Api server

I was not sure if it was to Extend the Kubernetes API Server or to create a new API only, but because it didn't ask to extend splicity, I decided to create a API Server with some shared code with the cli.

I decide to use Chi framework because I have more affinity with it and it is a well balanced framework.

It asked to replay with the candidate name which I was not sure if it was to be the release candidate name, but because it didn't ask specifically about the release candidate, I conclude that the candidate name was the candidate to the role which in this care is the person developming the test, which is me.

### Cli

The cli itself is simple but still I decide to use Cobra framework because it is simple to start and if the cli grow big it is already in a good pattern.

I decide to use a interface before print out the messages to help doing tests.
Same way to os.Exit any exit in a cli should have a number different from zero, but calling os.Exit directly it will stop the tests, that is why I create a option to fake the exit by trhowing a panic.


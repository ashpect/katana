# katana
An advanced yet simple attack/defence CTF infrastructure in Go build upon k8s.

## Setup
- To start, you must have the following installed:
  - Go 1.18+
  - Minikube & kubectl
- Run `make set-env` to setup the environment
- To start katana, run `./bin/katana run`

## Docs
- To get the most stable docs : 
  - Setup docs locally by running `make setup-docs` followed by `hugo serve` from the `docs` directory.
  - The docs are also available at [https://blog.sdslabs.co/katana](https://blog.sdslabs.co/katana/Katana/getting-started/). 
- To visit the most updated docs, it's advised to setup the docs locally from the docs branch.
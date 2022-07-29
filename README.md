# Backend Test Project
1. Create a new branch named 'feature/project-set-up'
2. Copy the directories and files into the project. Small explanations:
   - README.md – the project documentation
   - Makefile – the file that makes shortcuts for the most useful commands. Usage: `make <command name>`. For example, `make run`, `make build`
     - `run` - run the project locally
     - `build` - build the project locally
     - `compose-up` - run/re-run the whole project using Docker
     - `compose-up-app-only` - run/re-run only the application using Docker
     - `compose-up-no-app` - run/re-run the whole stuff but application using Docker
     - `compose-down` - stop the whole project using Docker
   - `.gitignore` - list of the files that are ultimately ignored and wouldn't be committed to the repo
   - `src` - the working directory
     - `src/cmd` - the directory with the executable files
       - `src/cmd/api` - the directory with the executable files related to API
         - `src/cmd/api/main.go` - the main executable API file
   - `deployment` - the directory with deployment and DevOps stuff
     - `deployment/Dockerfile` - file-instruction for Docker how to set up and build the application
     - `deployment/docker-compose.yml` - file-instruction for Docker how to run the whole stuff related to the project (application, database, etc)
   - `example.env` - the example file with the required environment variables. Environment variables are the values that depend on the machine/environment it running (for example, the database credentials for the local development and the production would be different). You should create a `.env` file in the same directory and fill every one with your unique values (or just copy the ones from the example)
   - `wait-for-it.sh` - the utility file for DB running.
3. Try to run the project locally and using Docker.
4. Commit and push the changes. Create a PR on the Github.
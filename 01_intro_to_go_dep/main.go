package main

func main() {

/*
	WHY dep?
	The dep tool is the “official experiment” dependency management tool for the go programming language. 
	It helps you to manage the ever-growing list of dependencies your project needs to maintain without a lot of overhead 
	and it can pin you to specific versions of dependencies to ensure stability in your systems.

	1. Installation
		brew install dep
		brew upgrade dep

	2. dep init
		When you call dep init, the tool does a few things:
		It identifies the dependencies of your current project
		It validates whether or not these dependencies use the dep tool
		It picks the highest compatible version for each of these dependencies

	3. Creating a New Project: 
		The first and possibly best option is to create your project within your $GOPATH, 
		much like you normally would, cd into that directory and then call dep init:
		$ mkdir -p $GOPATH/src/github.com/my/project
		$ cd $GOPATH/src/github.com/my/project
		$ dep init
		$ ls
		Gopkg.toml Gopkg.lock vendor/

	4. Gopkg.toml
		The Gopkg.toml file is where you specify your dependencies and the particular versions of these dependencies that you wish your project to use. 
		Think of this as your package.json if you are coming from a NodeJS background, or your pom.xml if you are coming from a Java background.

	5. Gopkg.lock
		The Gopkg.lock file is a transitively complete snapshot of your project’s dependency graph that is expressed as a series of [[project]] stanzas.
		In layman’s terms this is a list of every dependency and the particular revision of that dependency.

	6. The vendor/ Directory
		The vendor/ directory is where your dependencies are stored. It’s the equivalent to the node_modules/ directory in your NodeJS projects.

	7. Helpful Commands: The dep command features 5 commands in total
		init - Sets up a new Go project
		status - Reports the status of a project’s dependencies
		ensure - Ensures a dependency is safely vendored in the project
		prune - Prunes your dependencies, this is also done automatically by ensure
		version - Shows the dep version information

		You’ll typically only work with the first 3 commands so I’ll just be covering these in more detail.

		7.1 dep ensure
			The dep ensure command is quite possibly the most important command you will need to come to grips with when it comes to working with the dep dependency management tool.
		7.2 Adding Dependencies
			If you want to add new dependencies to your project you can do so by calling the dep ensure -add command and specifying the source for the project.
			$ dep ensure -add github.com/foo/bar github.com/another/project ...
		7.3 Updating Dependencies
			Should you wish to update some of the dependencies within your project you can do that using the -update flag when calling dep ensure:
			// dry run testing an update
			$ dep ensure -update -n
			// non-dry run
			$ dep ensure -update
			// updates a specific package
			$ dep ensure -update github.com/gorilla/mux
			// updates to a specific version
			$ dep ensure -update github.com/gorilla/mux@1.0.0
		7.4 dep status
			The dep status command reports the status of a project’s dependencies:
			$ dep status
			// output
			...
*/

}

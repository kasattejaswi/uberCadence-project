# Project on Uber Cadence

Cadence is a workflow orchestrator tool designed by Uber Engineering team. It basically helps focussing on coding workflows rather than handling situations which can be out of control quickly. 

## Basic use case

Let's take an example of a use case. After taking a ride in a cab, we may want to tip the driver. How this thing can be handled? We can debit the user's account and credit the driver's account using this basic function:
```
function tip() {
    debitUser()
    creditDriver()
}
```
Till everything is working fine, this function is OK. Since debit and credit requires the system to connect to external payment APIs, there are chances that these external services may fail causing inconsistent states. For example, what if debit from user's account was successful but credit to driver's account gets failed. It's an inconsistent state and must be handled.<br>
How to handle it?<br>
A simple way to introduce DB which will store the states of each transaction. So it will look something like this:
```
function debitUser() {
    debit()
    updateDB()
}

function creditDriver() {
    credit()
    updateDB()
}
```
This code looks perfect but it has introduced multiple problems. What if debit() was successful but DB entry failed? Or what if credit() was successful but DB entry gets failed? We have to write more functions to handle these situations. Things will get complicated quickly and will result in a codebase that will be difficult to manage.

## Uber Cadence to the rescue

Uber Cadence allows us to write code in this format:
```
function tip() {
    debitUser()
    creditDriver()
}
```
And rest of the complexities it handles itself. But how? When we run this function, cadence will make sure that debitUser() will run no matter what happens until the result is successful. Its so safe that even if the payment services fail, or cadence process itself goes down because of some outage, it will still run after everything goes back normal.<br>
The functions debitUser() and creditDriver() will be termed as activities and function trip() will be termed as a workflow in Cadence terminology.

## Document links
Currently Cadence supports only two languages for coding: Go and Java. Support for other languauges are still in development phase.<br>
Core concepts of Cadence:  https://cadenceworkflow.io/docs/concepts/<br>
Go client: https://cadenceworkflow.io/docs/go-client/<br>
Github: https://github.com/uber/cadence<br>
Official examples: https://github.com/uber-common/cadence-samples/tree/master/cmd/samples/recipes<br>

## What's in this project?

This project contains the initial setup required in order to run your first workflow. A sample hello world workflow is also given. The workflows can be started via command line. All you need to write your workflows under <b>workflows</b> folder and register that workflow and its activities in <b>localworker/registries</b> as given in helloworld sample.

The command line usage is as follows:<br>

### Basic
```
Usage:
  uberCadence-project [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Writes a default configuration file under user's home directory
  worker      Start a worker registering all available workflows
  workflow    Helps in running workflows

Flags:
  -h, --help   help for uberCadence-project
```

### init
```
Usage:
  uberCadence-project init [flags]

Flags:
  -f, --force         Use force to replace existing configuration and generate a new one.
  -h, --help          help for init
  -p, --path string   Path where config file will be written. Default will be user's home directory.
```

### worker
```
Usage:
  uberCadence-project worker [command]

Available Commands:
  start       Perform start action

Flags:
  -h, --help   help for worker
```

### workflow
```
Usage:
  uberCadence-project workflow [command]

Available Commands:
  list        List all the available workflows
  run         Run a workflow using name or run all available workflows

Flags:
  -h, --help   help for workflow
```

I would recommend to spend time with <a href="https://github.com/uber-common/cadence-samples/tree/master/cmd/samples/common">common</a> folder. It will help to give a clear idea of workers. 
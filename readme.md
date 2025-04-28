# proman - protocmanager
    protocmanager is cli tool to simplify the process of
        * setting up your machine with protobuff compiler(protoc)
        * installing compiler's language specific plugins
        * generating source files from proto files

# install help
for now install using `go install`

# how to use
if proman supports your prefferred language, when you run `proman gen` for the first time it'll setup things for you, like downloading and installing the protobuff compiler (protoc), and downloading and installing the language specific plugins.

when you want to generate code for your `*.proto` files all you have to invoke is
`proman --lang=<1st-lang>,<second-lang> --in=./protos-folder --out=./generated-code-folder`
or
`proman --lang=go,dart --in=./protos-folder --out=./generated-code-folder`

# config
if you like doing things from configuration file, that is possible
## init a config file
`proman cfg init` inside a project
## make the apropriate changes
a file with name `.proman` will be generated, modify the config as required
## genrate the code
run `proman gen` without any arguments or with arguments to override whats in the config file

# 🚀 Language Support

| Language  | Codegen Plugin(s)                   | Status           | Notes                                |
|-----------|--------------------------------------|------------------|--------------------------------------|
| Go        | `protoc-gen-go`, `protoc-gen-go-grpc` | ✅ Supported     | Standard Go + gRPC generation.       |
| Dart      | `protoc_plugin`                      | ✅ Supported      | Dart/Flutter gRPC support.           |
| Java      | *(planned)*                         | 🚧 In Progress    | Will use `protoc-gen-java`.           |
| Python    | *(planned)*                         | 🚧 In Progress    | Will use `protoc-gen-python`.         |
| TypeScript (gRPC-Web) | *(planned)*              | 🚧 In Progress    | Will use `protoc-gen-ts`, `protoc-gen-grpc-web`. |
| C++       | *(future)*                          | 🕒 Planned         | Native C++ support.                  |


# proman - protocmanager
this tool keeps track of `protoc` and and its language specific plugins in a raw and hacky way, rather than relying on the path variables.

    protocmanager is cli tool to simplify the process of
        * setting up your machine with protobuff compiler(protoc)
        * installing compiler's language specific plugins
        * maintains a copy of protos by google `https://github.com/protocolbuffers/protobuf`
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
| TypeScript (gRPC-Web) | *(planned)*              | 🚧 In Progress      | Will use `protoc-gen-ts`, `protoc-gen-grpc-web`. |
| Java      | *(planned)*                         | 🕒 Planned      | Will use `protoc-gen-java`.           |
| Python    | *(planned)*                         | 🕒 Planned      | Will use `protoc-gen-python`.         |
| C++       | *(future)*                          | 🕒 Planned         | Native C++ support.                  |


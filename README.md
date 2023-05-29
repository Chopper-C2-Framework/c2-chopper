# C2 Chopper Framework

<a name="readme-top"></a>



<!-- PROJECT LOGO -->
<br />
<div align="center">
<img src="https://github.com/Chopper-C2-Framework/c2-chopper/assets/62627838/f810ef27-2448-4d44-ba3e-7231c48c1071" width=300>

  <h3 align="center">C2 Chopper - Framework</h3>

  <p align="center">
    This framework for the final term project for the Personal Professional Project of SE 3 2023-2022 at INSAT ⭐
    <br />
    <a href="https://docs-c2-chopper.vercel.app"><strong>Explore the docs »</strong></a>
    <br />
    <br /><a href="https://github.com/Chopper-C2-Framework/c2-chopper/issues">Report Bug</a>
    <a href="https://github.com/Chopper-C2-Framework/c2-chopper/issues">Request a feature</a>

  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![image](https://github.com/Chopper-C2-Framework/c2-chopper/assets/62627838/8ceac991-e3c3-4b14-b339-d579db6d3caa)

Our C2 framework project is designed to provide a flexible and extensible platform for managing teams and operations throughout penetration testing missions. 

The framework is built upon a plugin-based architecture, which allows users to easily add new functionality and customize the platform to suit their specific needs and extends the basic functionalities to enhance their . With a wide variety of plugins available, including ones for remote administration, data collection and analysis, and automation, the framework is highly adaptable and can be tailored to a variety of use cases. Additionally, the framework includes robust security features to ensure that all communication and data transfer is encrypted and protected from unauthorized access. Whether you're looking to manage a small network of devices or a large-scale industrial control system, our C2 framework project provides the flexibility and scalability you need to get the job done.


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.

* [**Golang**](https://go.dev)
* ![React][React.js](https://react.dev/)
* [**GRPC**](https://grpc.io/)
* [**Shadcn**](https://ui.shadcn.com)
* [**React Query**](https://tanstack.com/query/v3/docs/react/overview)
* [**Go Fiber**](https://gofiber.io/)

and more ..

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To setup the framework on your team server do the following

### Prerequisites

#### Installing go v1.20
Make sure you have golang v1.20 installed if not do the following

##### Ubuntu

```shell
sudo apt-get update 
sudo apt-get upgrade 
wget https://go.dev/dl/go1.20.linux-amd64.tar.gz 
tar -xvf go1.19.linux-amd64.tar.gz 
mv go /usr/local
```

Export environment variables 

```shell
export GOROOT=/usr/local/go 
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
```

#### Verifying installation

```shell
$ go version 
go version go1.20 linux/amd64
```

### Installation

```shell
cd ~/c2-chopper
go build -o c2-chopper cmd/main.go
./c2-chopper 
```

![image](https://github.com/Chopper-C2-Framework/c2-chopper/assets/62627838/1f4ccd8b-e6de-412e-947c-3b10e7113253)

You can copy the binary to your `/usr/bin/` and use it in the system

```shell
sudo cp ./c2-chopper /usr/bin/c2-chopper
c2-chopper --help 
```



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage


### Configuration 

To start using the framework you have to launch the first time as it will setup the configuration needed. Do not worry you can later modify the configuration. 
A configuration file will be generated at `$HOMEDIR/.c2-chopper/config.yaml` and it will looks like the following

```yaml
plugins_path: .c2-chopper/plugins
client_port: 9001
server_port: 9002
server_http_port: 4000
host: localhost
server_cert_path: ./cert/server-cert.pem
sever_cert_key_path: ./cert/server-key.pem
use_tls: false
server_db_path: server.db
secret_token: super_duper_secret_token
```

It's short, clear and concise you can modify it and rerun your binary so the modifications take effect.

#### To generate configuration

```c2-chopper
c2-chopper gen-conf
```

#### Checking configuration correctness

```c2-chopper
c2-chopper cconf
```


### Server

### Launching server

```shell
c2-chopper server
```

![image](https://github.com/Chopper-C2-Framework/c2-chopper/assets/62627838/a588aeac-a7ba-4ecb-8b9c-2360a05cbb25)

### Launcing client

```shell
c2-chopper client
```
![image](https://github.com/Chopper-C2-Framework/c2-chopper/assets/62627838/61cb2d46-c2bf-4be0-be79-d409939112b9)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Create documentation website
- [ ] Add more control on exploited hosts visibility per team. 
- [ ] Make plugins loading better and more ergonomic
- [ ] Improve Frontend shell experience
- [ ] Multi-language Support
    - [ ] Chinese
    - [ ] French

See the [open issues](https://github.com/Chopper-C2-Framework/c2-chopper/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Yassine Belkhadem ~ fir3cr4ckers - [@YBK_FireLights](https://twitter.com/YBK_FireLights) - yassine.belkhadem@insat.ucar.tn
Mohamed Mongi Saidane ~ M0ngi - [@M0ngio](https://twitter.com/M0ngii) - mohamedmongi.saidane@insat.ucar.tn

Project Link: [C2-Chopper](https://github.com/Chopper-C2-Framework/c2-chopper/edit/development/README.md)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## Contribution: Starting point

### Prerequisites

* Protocol buffer compiler, `protoc`. [Installation Guide.](https://grpc.io/docs/protoc-installation/)
* Go plugins for the protocol compiler.
    1. Install the protocol compiler plugins for Go using the following commands:
    ```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```

    1. Install gRPC-Gateway plugin:
    ```
    $ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
    $ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
    ```

    1. Update your PATH so that the protoc compiler can find the plugins:
    ```
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```

## HTTP -> GRPC Gateway

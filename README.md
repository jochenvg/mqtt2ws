[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache License][license-shield]][license-url]



# MQTT to MQTT over WebSocket Bridge

A simple MQTT to MQTT over WebSocket bridge written in go.

Will listten on a configurable TCP port for MQTT messages. When a client connects a Websocket connection is established to a configurable remote server.
The MQTT messages from the TCP socket are bridged to the Websocket connection and vice versa.


## Getting Started

To get a local copy up and running follow these simple steps.


### Installation

#### Using go get

1. Install go 
2. Go get module
   ```sh
   go get github.com/jochenvg/mqtt2ws
   ```

#### Download release

Alternatively, you can download a binary for your operation system on the releases page.


## Usage

Operation is controller by two flags:
```
Usage of ./mqtt2ws:
  -listen string
    	listen tcp socket (default ":1883")
  -url string
    	websocket url (default "ws://broker.hivemq.com:8000/mqtt")
```


## License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.


[contributors-shield]: https://img.shields.io/github/contributors/jochenvg/mqtt2ws.svg?style=for-the-badge
[contributors-url]: https://github.com/jochenvg/mqtt2ws/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jochenvg/mqtt2ws.svg?style=for-the-badge
[forks-url]: https://github.com/jochenvg/mqtt2ws/network/members
[stars-shield]: https://img.shields.io/github/stars/jochenvg/mqtt2ws.svg?style=for-the-badge
[stars-url]: https://github.com/jochenvg/mqtt2ws/stargazers
[issues-shield]: https://img.shields.io/github/issues/jochenvg/mqtt2ws.svg?style=for-the-badge
[issues-url]: https://github.com/jochenvg/mqtt2ws/issues
[license-shield]: https://img.shields.io/github/license/jochenvg/mqtt2ws.svg?style=for-the-badge
[license-url]: https://github.com/jochenvg/mqtt2ws/blob/master/LICENSE.txt

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache License][license-shield]][license-url]



# mqtt2ws - Transparently proxy MQTT to MQTT-over-Websocket

This is a simple MQTT to MQTT-over-Websocket proxy written in go which preserves MQTT packet alignment with websocket frames.

Mqtt2ws will listten on a configurable TCP port for MQTT messages. When a client connects a Websocket connection is established to a configurable remote server.
The MQTT messages from the TCP socket are then bridged to the Websocket connection in both directions.

## Background

MQTT is a light-weight publish/subscribe protocol frequently used for communication with IoT devices. It can be transported over TCP, over TLS or over WebSocket.
The MQTT payload is identical regardless of the the protocol used to transport it.

TCP is a stream based protocol, meaning that as far the MQTT protocol using it is concerned, it is a continuous stream of bytes. 
Websocket however is a message based a message/frame based protocol. 

The MQTT-over-WebSocket spec does not require alignment of MQTT messages with WebSocket frames, and WebSocket frames can contain more than one MQTT message or partial messages.
It appears however that not all MQTT server implementations, including Mosquitto when used with the "protocol websockets" directive adhere to this permissive soec, and connections fail when not aligned.

Many MQTT-over-WebSocket client inplementaations ensure alignment of MQTT messages witg WebSocket frames, including the Eclipse Paho client libraries.

Mqtt2ws uses the Eclipse Paho library to read MQTT messages from the TCP session and send them over the WebSocket so alignment is preserved.


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

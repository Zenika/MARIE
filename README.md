# MARIE

Module d'Administration Régissant les Interfaces Electroniques

This software exists to administrate electronic things in a domotic solution. 

## Architecture

The backend folder contains the server. It connects to MQTT broker and starts a HTTP and a WS server.

The backoffice folder contains a Vue.JS application. It can create new things and see the things that was created, if they are online and can launch actions on them.

The things folder contains scripts for things for different type of hardware: Arduino, Raspberry Pi and NodeMCU.

## Authors 

* **Valentin STERN** - [Sehsyha](https://github.com/Sehsyha)

## License 

This project is under the MIT License - see the [LICENSE](https://github.com/Zenika/MARIE/blob/master/LICENSE) file for details.
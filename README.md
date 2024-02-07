# Mercenary - A Reliable Process Killer

Welcome to Mercenary, a robust and efficient process killer tool developed in Go. It's purpose-built to terminate processes post-execution or in response to specific events that cause them to zombify or exhibit undesirable behavior. It's designed to provide a reliable solution for managing system processes and maintaining optimum resource usage.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Features

- Efficient process termination
- Handling of zombified processes
- Response to undesirable process behavior
- Structured logging

## Installation

### Prerequisites

Ensure you have Go 1.20 or later installed on your system to build using Go.

#### Installation via Go

```shell
go get gitlab.servicenow.net/cce/hweng/hardware-test/mercenary.git
```

#### Installation via SSH

```shell
git clone https://gitlab.servicenow.net/cce/hweng/hardware-test/mercenary.git
```

#### Installation via HTTPS

```shell
git clone https://gitlab.servicenow.net/cce/hweng/hardware-test/mercenary.git
```

#### Post Cloning Steps with pre-compiled binary
Navigate to the /bin directory and execute the binary according to your system architecture:
```shell
# For Linux
./bin/mercenary-linux puppet iptables mariadb
# For MacOS
./bin/mercenary-darwin puppet iptables mariadb
```

#### Post Cloning Steps
Navigate to the cloned repository and run the following command to install the Mercenary binary:
```shell
go install
```

## Usage

```shell
# Allow to ask nicely first before killing (SIGTERM)
./mercenary <PID or process name>
```

```shell
# To kill without asking nicely first (SIGKILL)
./mercenary kill <PID or process name>
```

## Examples

```shell
# Kill a process with pid 1234
./mercenary kill 1234
```

```shell
# Kill a process with name "mercenary"
./mercenary mariadb
```

```shell
# Kill multiple processes with PID's
mercenary kill 1234 5678 9012
```

```shell
# Kill multiple processes with names
./mercenary kill swtichEngage mariadb firewalld
```

## License

This software is property of ServiceNow. All rights reserved.

## Contact

- Kenny Sheridan

For any troubles or inquiries, reach out to us at hwtest@servicenow.com





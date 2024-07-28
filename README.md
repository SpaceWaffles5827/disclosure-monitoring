# Disclosure Server

## Overview

The Disclosure Server is a server-side application that listens for incoming TCP connections and logs received data. It is intended to be used in conjunction with the Remote Monitoring and Control Project, allowing for keylogging, webcam access, and remote desktop control data to be sent from clients and logged on the server.

**IMPORTANT: This software is intended solely for ethical and legal purposes. Unauthorized use of remote monitoring tools is illegal and unethical. Do not use this software for malicious purposes. Always obtain explicit consent before running this software in any environment.**

## Setup

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/YourUsername/disclosure-server.git
   cd disclosure-server
   ```

2. **Install Dependencies:**
   - Ensure you have Python 3 installed.
   - This script uses standard libraries, so no additional packages should be needed.

3. **Run the Server:**
   - Execute the server script to start listening for incoming connections.
     ```sh
     python server.py
     ```

## Usage

- The server will listen on all available interfaces (`0.0.0.0`) and port `9002` by default.
- When a client connects, the server will create a log file named based on the client's IP address and hostname.
- All received data will be logged to this file.

## Important Notes

- **Ethical Use:** Use this software only for ethical and legal purposes. Unauthorized remote monitoring is illegal and unethical.
- **Explicit Consent:** Always obtain explicit consent before running this software in any environment.
- **Security:** Ensure that the server is securely managed and protected against unauthorized access.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Disclaimer

This software is provided "as is" without any warranties or guarantees. The authors are not responsible for any misuse of this software. Use it at your own risk and responsibility.

---

For any issues or contributions, please open an issue or a pull request on GitHub.

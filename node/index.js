import { createConnection } from "net";

export class AppleDoreClient {
  options;
  client;
  connected;

  constructor(options) {
    this.options = options;
    this.client = null;
  }

  connect(handler) {
    return new Promise((resolve, reject) => {
      this.client = createConnection(this.options, handler);

      this.client.on("connect", () => {
        this.connected = true;
        resolve("connected");
      });

      this.client.on("error", (err) => {
        reject(err);
      });
    });
  }

  sendCommand(command, ...args) {
    return new Promise((resolve, reject) => {
      if (!this.connected) {
        reject(new Error("Not connected to Appledore"));
        return;
      }

      const redisCommand = this.#buildCommand(command, args);
      this.client.write(redisCommand);

      let responseData = "";

      this.client.once("data", (data) => {
        responseData += this.#encoder(data.toString());
        resolve(responseData);
      });
    });
  }

  #buildCommand(command, args) {
    return `*${1 + args.length}\r\n$${command.length}\r\n${command}\r\n${args
      .map((v) => `$${v.length}\r\n${v}\r\n`)
      .join("")}\r\n`;
  }

  #encoder(data) {
    const arr = data.split("\r\n")
    arr.pop()
    if (arr.length === 0) return ""
    const statusString = arr[0];
    if (statusString[0] === "-") return ""
    if (statusString[0] === "+") return "OK"
    if (statusString[0] === "$") {
      if (arr.length === 1) return ""
      if (arr.length === 2) return arr[1]
    }
    return ""
  }

  disconnect() {
    this.client.end();
    this.connected = false;
    console.log("Disconnected from AppleDore");
  }
}


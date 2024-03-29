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
    let noEscapeString = data.replace(/[\r\n]/g, "").replace(/\$\d+/g, "");
    if (noEscapeString.startsWith("+")) {
      noEscapeString = noEscapeString.substring(1);
    }
    if (noEscapeString.startsWith(":")) {
      noEscapeString = noEscapeString.substring(1);
    }
    return noEscapeString;
  }

  disconnect() {
    this.client.end();
    this.connected = false;
    console.log("Disconnected from AppleDore");
  }
}

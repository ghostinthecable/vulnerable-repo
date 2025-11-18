const http = require('http');
const crypto = require('crypto');

const SECRET = "123456";

http.createServer((req, res) => {
    let body = "";

    req.on("data", chunk => body += chunk);
    req.on("end", () => {
        const hash = crypto.createHash("md5").update(body + SECRET).digest("hex"); // weak
        res.end("Hash: " + hash);
    });
}).listen(8080);

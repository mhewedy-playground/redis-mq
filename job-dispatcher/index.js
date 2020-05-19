const redis = require("redis");
const client = redis.createClient({
    address: "localhost:6379"
});

client.on("error", function (error) {
    console.error(error);
});

let myArgs = process.argv.slice(2);
let start = Number(myArgs[0])
let end = start + 10000

console.log(end)

for (let i = start; i < end; i++) {
    client.rpush("myJobQueue", i);
}

client.quit();

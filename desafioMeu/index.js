const numbers = require("fs").readFileSync("input.txt", "utf8");


const re = /\d+/g

var set = new Set();

count = 0;

do {

    m = re.exec(numbers);
    if (m) {
        count += 1;
        set.add(m[0])
    }
} while (m);

console.log(count);
console.log(set.size);
var fs = require('fs');

const data = fs.readFileSync('input.txt', 'utf8')
const lines = data.split(/\r?\n/);

var count = 0;


// part 1 
lines.forEach((line) =>{
    var [pswdInfo, pswd] = line.split(':')
    var [atLeast, atMost, letter] = pswdInfo.split(/[- ]/)

    const regex =  new RegExp(letter,'g');

    matched = pswd.match(regex) !== null ? pswd.match(regex) : [];
   
    if((matched.length >= atLeast)  && (matched.length <= atMost)){
        count = count +1;
    }
    
}) 

console.log(count);
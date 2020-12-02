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


//part 2 
// part 1

validCounter = 0;
lines.forEach((line) =>{
    var [pswdInfo, pswd] = line.split(':')
    var [position1, position2, letter] = pswdInfo.split(/[- ]/)

    const regex =  new RegExp(letter,'g');

    var positionsArray =[]
    while(match = regex.exec(pswd)){
        positionsArray.push(match.index);
    }
       if(positionsArray.includes(Number(position1)) && !positionsArray.includes(Number(position2))){
        validCounter = validCounter +1;
       }
    
       if(!positionsArray.includes(Number(position1)) && positionsArray.includes(Number(position2))){
        validCounter = validCounter +1;
       }


}) 


console.log(validCounter)




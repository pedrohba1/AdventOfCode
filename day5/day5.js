
var fs = require('fs')

const input = fs.readFileSync('input.txt', 'utf8')
var passes =  input.split('\n');

const search  = (min, max, pass) => {
    var upper = 127;
    var lower = 0;

    [...pass].forEach(letter => {
        console.log(letter);
        mid = min + ((max - min) /2);
        useTop = letter === 'B' || letter === 'R';
        min = useTop ? Math.ceil(mid) : min; 
        max = useTop ? max : Math.floor(mid);
    });
    return min;
}

const ids = passes.map(pass =>{
    const row = search(0,127, pass.slice(0,7));
    const column = search(0,7, pass.slice(7));
    return (row*8 ) + column;
})

// part 1
const highestId = Math.max(...ids);
console.log(highestId);
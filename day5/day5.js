
var fs = require('fs')

const input = fs.readFileSync('input.txt', 'utf8')
var passes =  input.split('\n');

function search (min,max,codes){
    const code = codes[0];
    const mid = min + ((max - min) /2);
    const useTop = code === 'B' || code === 'R';

    let newMin = useTop ? Math.ceil(mid) : min;
    let newMax = useTop ? max : Math.floor(mid);

    if(codes.length ===1) return useTop ? newMax : newMin;
     else return search(newMin, newMax, codes.slice(1));
}


const ids = passes.map(pass =>{
    const row = search(0,127, pass.slice(0,7));
    const column = search(0,7, pass.slice(7));
    return (row*8 ) + column;
})

const highestId = Math.max(...ids);
console.log(highestId);
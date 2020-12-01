var fs = require('fs');

const data = fs.readFileSync('input.txt', 'utf8')
const lines = data.split(/\r?\n/);


//part 1

lines.forEach((line) =>{
    lines.forEach((lineSec) =>{                  
        mul = line*lineSec;
        sum = Number(line) + Number(lineSec);
        if(sum === 2020){
            console.log(sum)
            console.log(mul)
        }
    })
})




//part 2
lines.forEach((line) =>{
    lines.forEach((lineSec) =>{        
        lines.forEach((lineTerc) =>{            
            mul = line*lineSec*lineTerc;
            sum = Number(line) + Number(lineSec) + Number(lineTerc)
            if(sum === 2020){
                console.log(sum)
                console.log(mul)
            }
        })
    })
})



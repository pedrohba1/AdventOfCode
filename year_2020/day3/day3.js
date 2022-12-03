var fs = require('fs')

const input = fs.readFileSync('input.txt', 'utf8')
const lines =  input.split('\n')

var count = 0
var map2D = []


lines.forEach((line) =>{
    arrayedLine =line.split('')
    arrayedLine.pop('\r')
    map2D.push(arrayedLine)
});


function doPath( right , down, map ) {
    let trees = 0
    let x = 1
    let y = 1
    while( y <= map.length ) {
        if( ( map[ y-1 ] )[ x-1 ] == '#') trees++
        x += right
        y += down
        if( x > map[0].length) x -= map[0].length
    }
    return trees
}

/// First Star
console.log(doPath( 3 , 1, map2D))

//Second star

let result = 1
result *=  doPath( 1 , 1, map2D )
result *= doPath( 3 , 1,map2D )
result *= doPath( 5 , 1, map2D )
result *= doPath( 7 , 1, map2D )
result *= doPath( 1 , 2, map2D )

console.log(result)
var fs = require('fs')

const input = fs.readFileSync('input.txt', 'utf8')
var documents =  input.split('\n\n');

const regex = /(\w{3}):([\w#]+)/g


documentJSONs =  documents.map((line) =>{
    documentJSON = {}
    let found
    do {
        found = regex.exec(line)
        if(found) documentJSON[found[1]] = found[2]  
    }while(found)
    return documentJSON
})


// part 1
function isValid (documentJSON) {
    keys = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]
   fieldsAmount =  Object.keys(documentJSON).length

   validKeys = true
    keys.forEach((key) =>{
        if (!documentJSON.hasOwnProperty(key)) validKeys = false
    })
    if(validKeys) return true
    else return false
}

count = 0
documentJSONs.forEach((json) =>{
    if(isValid(json)) count += 1
})

console.log(count)




//part 2 

function validateHeight(hgt){
    if (hgt === undefined ) return false
    m = hgt.match(/^(\d+)(cm|in)$/)
    if (m && m[2] && m[2] == "in") {
      if( m[1] >= 59 && m[1] <= 76) return true
    }
    if (m && m[2] && m[2] == "cm") {
        if(m[1] >= 150 && m[1] <= 193) return true
    }
    return false;
}



function validateDoc(document){
    const {byr, iyr, eyr, hgt, hcl, ecl, pid} = document
    if (byr >= 1920 && byr <= 2002 ) validate = true; else  return false
    if (iyr >= 2010 && iyr <= 2020 ) validate =  true ; else return  false
    if ( eyr >= 2020 && eyr <= 2030 ) validate = true ; else  return  false 
    if ( validateHeight(hgt)) validate =  true; else  return  false
    if ( /^#[0-9a-f]{6}$/.test(hcl)) validate =  true; else return false
    if ( /^brn|amb|blu|gry|grn|hzl|oth$/.test(ecl)) validate =  true; else return  false
    if ( /^\d{9}$/.test(pid)) validate =  true; else return  false
    return validate
}


count = 0
documentJSONs.forEach((json) =>{
  if(validateDoc(json)) count+= 1
})

console.log(count)





const input = require("fs").readFileSync("input.txt", "utf8").split("\n");

const bags = {}
const goldBag = "shiny gold"


input.forEach((line) =>{
    const parts = line.split(' contain ');
    const id = parts[0].replace('bags', '').trim();
    const rules = parts[1].matchAll(/(\d+) (\w+ \w+)/g);
    bags[id] = [...rules].map(rule => [parseInt(rule[1]), rule[2]])
})


const numContain = (id, target) => {
    let total = 0;
    for (const rule of bags[id]){
        if(rule[1] === target){
            total += rule[0]
        }
        total += numContain(rule[1], target);
    }
    return total;
}


const numInnerBags = id =>{
    let total = 0;
    for (const rule of bags[id]){
        total += rule[0] + numInnerBags(rule[1]) *rule[0]
    }
    return total;
}

const totals = Object.keys(bags).map(id => numContain(id, goldBag))
const canContain = totals.filter(n => n >0 ).length;
const canContainAmount = totals.reduce((a,b) => a+b, 0);
console.log('canContain', canContain);
console.log('inner bags', numInnerBags(goldBag));
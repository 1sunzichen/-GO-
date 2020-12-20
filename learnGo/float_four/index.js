let obj={}
Object.defineProperty(obj, 'name', {
    value: 42,


});
for( let o in obj){
    console.log(obj[o],"qqq")
}
// console.log(obj,"obj")
#!/usr/bin/env node
const questionsQuantity = 40
const prefix = process.argv[2];
const quantityText = process.argv[3];

const possibleResponses = {
  eleitor: ['CP', 'C', 'I', 'D', 'DP'],
  candidato: ['CP', 'C', 'D', 'DP'],
}

const help = () =>
  console.log(`
      Uso:
        gerador.js <modo> <quantidade>
        gerador.js eleitor 10
        gerador.js candidato 1000
      `)

const validate = (prefix, quantity) =>
  new Promise(
    (resolve, reject) =>
      (!prefix || isNaN(quantity) ) ?
        reject() :
        resolve()
  )

const getRandomItemFromArray = (array) =>
  array[ Math.floor(Math.random() * array.length) ]

const generateResponses = (prefix, index) =>
   Array.from(
     {length: questionsQuantity},
     () => getRandomItemFromArray(possibleResponses[prefix])
   )

const generatePeople = (prefix, quantity) =>
  Array.from(
    {length: quantity},
    (value, index) => ({
      [`${prefix}-${index}`]: generateResponses(prefix)
    })
  )

const run = (prefix, quantity) =>
  validate(prefix, quantity)
    .then( () => generatePeople(prefix, quantity) )
    .then( (responses) => Object.assign({}, ...responses) )
    .then(JSON.stringify)
    .catch(help)


run(prefix, quantityText)
  .then(console.log)

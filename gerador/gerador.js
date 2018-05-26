#!/usr/bin/env node
const questions_number = 40
const modo = process.argv[2]
const num = parseInt(process.argv[3])

let help = () => {
    console.log(`
        Uso:
          gerador.js <modo> <quantidade>
          gerador.js eleitor 10
          gerador.js candidato 1000
        `)
    process.exit()
}

let validate = () => {
    if (process.argv.length !== 4 || !Number.isInteger(num)) {
        help()
    }
}

let generate = (idPrefix, num, alternatives) => {
    responsesPerSubject = {}

    for (subject = 1; subject <= num; subject++) {
        subjectReponses = []
        for (question = 0; question < questions_number; question++) {
            subjectReponses[question] = alternatives[Math.floor(Math.random() * alternatives.length)];
        }
        responsesPerSubject[idPrefix+'-'+subject] = subjectReponses
    }
    return responsesPerSubject
}

let run = () => {
    validate()
    switch(modo) {
        case 'eleitor':
            data = generate('eleitor', num, ['CP', 'C', 'I', 'D', 'DP'])
            break;
        case 'candidato':
            data = generate('candidato', num, ['CP', 'C', 'D', 'DP'])
            break;
        default:
            help()
    }

    console.log(JSON.stringify(data));
}

run()









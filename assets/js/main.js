import nexGenerationButton from "./nextGenerationButton.js"
import generateRulesButton from "./generateRulesButton.js"
import playButton from "./playButton.js"
import newAutomatonButton from "./newAutomatonButton.js"
import {resizeCanvas} from "./canvas.js"

document.addEventListener("DOMContentLoaded", function(){
    nexGenerationButton()
    generateRulesButton()
    playButton()
    newAutomatonButton()
    resizeCanvas()
});
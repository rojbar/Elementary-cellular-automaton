/**
 * 
 * @param {string} generation 
 * @param {int} size 
 * @param {int} posGeneration 
 */
 function drawGeneration(generation, size, posGeneration) {
    const canvas = document.getElementById('canvas');
    if (canvas.getContext) {
        const ctx = canvas.getContext('2d');
        ctx.font = `${size}px arial`;
        ctx.fillText(generation, size, 10 + (posGeneration * (size + 5)));
    }
}

export default drawGeneration
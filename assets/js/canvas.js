export function resizeCanvas() {
    const canvas = document.getElementById('canvas');
    canvas.style.width = "1080px";
    canvas.style.height = "920px";
    canvas.width = canvas.offsetWidth;
    canvas.height = canvas.offsetHeight;
}

export function clearCanvas() {
    const canvas = document.getElementById('canvas');
    if (canvas.getContext) {
        const ctx = canvas.getContext('2d');
        ctx.clearRect(0, 0, canvas.width, canvas.height);
    }
}

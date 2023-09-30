const canvas = document.querySelector('canvas');
const ctx = canvas.getContext('2d');
const [width, height] = [canvas.width, canvas.height];
const go = new Go()

async function loadWasm() {
    const arraySize = (width * height * 4) >>> 0;
    const nPages = ((arraySize + 0xffff) & ~0xffff) >>> 16;
    const memory = new WebAssembly.Memory({ initial: nPages });

    const wasm = await WebAssembly
        .instantiateStreaming(fetch('main.wasm'), go.importObject).
        then((result) => {
        mod = result.module;
        inst = result.instance;
    }).catch((err) =>{
        console.error(err)
    });

    const solveMazePromise = new Promise(resolve => {
        setSolve = resolve
    })
    const run = go.run(inst)

    const solveMaze = await solveMazePromise
    solveMaze()
    await run

    inst = await WebAssembly.instantiate(mod, go.importObject)

    displayMaze()
}

async function displaySolvedMaze() {
    const solveMazePromise = new Promise(resolve => {
        setSolve = resolve
    })

    const solveMaze = await solveMazePromise

    solveMaze()
    await run

    inst = await WebAssembly.instantiate(mod, go.importObject)

    displayMaze()
}

function displayMaze() {
    const img = new Image();
    img.src = './maze.png';
    img.crossOrigin = 'anonymous';
    img.onload = () => ctx.drawImage(img, 0, 0, width, height);
}

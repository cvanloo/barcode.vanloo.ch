const btn_generate = document.getElementById('btn-generate')
const scratch_pad = document.getElementById('scratch-pad')
let idx = 0;

function generateBarcode(e, bc_img_url, bc_text) {
    e.preventDefault()

    const div = document.createElement('div')
    div.classList.add('barcode')
    div.id = `bc-${idx}`

    const top_bar = document.createElement('div')
    top_bar.id = "top-bar"

    const btn_move = document.createElement('button')
    const btn_close = document.createElement('button')
    btn_move.type = "button"
    btn_close.type = "button"
    btn_move.innerHTML = "Move"
    btn_close.innerHTML = "Close"

    const cidx = idx
    btn_close.onclick = () => removeBarcode(cidx)

    top_bar.appendChild(btn_move)
    top_bar.appendChild(btn_close)
    div.appendChild(top_bar)

    const img = new Image(312, 80)
    img.src = bc_img_url
    div.appendChild(img)

    const p = document.createElement('p')
    p.innerHTML = bc_text

    div.appendChild(p)

    scratch_pad.appendChild(div)

    idx++

    e.target.reset()
}

function removeBarcode(idx) {
    const d = document.getElementById(`bc-${idx}`)
    scratch_pad.removeChild(d)
}

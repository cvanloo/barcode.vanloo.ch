const btn_generate = document.getElementById('btn-generate')
const scratch_pad = document.getElementById('scratch-pad')
let idx = 0;

function generateBarcode(bc_img_url, bc_text, bc_name) {
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

    const name_tag = document.createElement('p')
    name_tag.innerHTML = bc_name

    const cidx = idx
    btn_close.onclick = () => removeBarcode(cidx)

    top_bar.appendChild(btn_move)
    top_bar.appendChild(name_tag)
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
}

function removeBarcode(idx) {
    const d = document.getElementById(`bc-${idx}`)
    scratch_pad.removeChild(d)
}

function saveBarcode(code) {
    const session = sessionStorage.getSession()
    const bcs = localStorage.getObject(session) ?? []
    bcs[bcs.length] = code
    localStorage.setObject(session, bcs)
}

function createSession() {
    scratch_pad.textContent = ''
    const session = sessionStorage.newSession()
    localStorage.setObject(session, null)
}

function loadSession(session) {
    scratch_pad.textContent = ''
    sessionStorage.setSession(session)
    localStorage.getObject(session)?.forEach((bc) => {
        generateBarcode(bc.image, bc.data, bc.name)
    })
}

function sessionSelect(select) {
    const session = sessionStorage.getSession()

    select.textContent = ''

    const cn = document.createElement('option')
    cn.value = 'create_new'
    cn.innerHTML = 'Create New'
    select.appendChild(cn)

    Object.keys(localStorage)?.forEach((key) => {
        const opt = document.createElement('option')
        opt.value = key
        opt.innerHTML = key
        if (session === key) {
            opt.selected = true
        }
        select.appendChild(opt)
    })
}

import 'api/storage'

let _onBarcodesUpdate = () => {}
let _barcodes = []
let _moveAction = null;

function renderFunc(fun) {
    _onBarcodesUpdate = fun
}

function renderOnto(div) {
    _barcodes.forEach((bc, i) => {
        div.appendChild(_render(i, bc))
    })
}

/* Returns a div filled out with the barcode's data.
 * barcode = { name; url; text; }
 */
function _render(id, barcode) {
    // Order matters!

    const div = document.createElement('div')
    div.classList.add('barcode')

    const btn_move = document.createElement('button')
    btn_move.type = "button"
    if (_moveAction === null) {
        btn_move.innerHTML = "Move"
        btn_move.onclick = () => {
            _moveAction = {}
            _moveAction.from = id
            _onBarcodesUpdate()
        }
    } else if (_moveAction.from === id) {
        btn_move.innerHTML = "Stop"
        btn_move.onclick = () => {
            _moveAction = null
            _onBarcodesUpdate()
        }
    } else {
        btn_move.innerHTML = "Here"
        btn_move.onclick = () => {
            move((_moveAction.to = id, _moveAction))
            _moveAction = null
            _onBarcodesUpdate()
        }
    }

    const name_tag = document.createElement('p')
    name_tag.innerHTML = barcode.name

    const btn_close = document.createElement('button')
    btn_close.type = "button"
    btn_close.innerHTML = "Close"
    btn_close.onclick = () => remove(id)

    const top_bar = document.createElement('div')
    top_bar.id = "bc-top-bar"
    top_bar.appendChild(btn_move)
    top_bar.appendChild(name_tag)
    top_bar.appendChild(btn_close)
    div.appendChild(top_bar)

    const img = new Image(312, 80)
    img.src = barcode.url
    div.appendChild(img)

    const text_tag = document.createElement('p')
    text_tag.innerHTML = barcode.text
    div.appendChild(text_tag)

    return div
}

/* Add a barcode
 * barcode = { name; url; text; }
 */
function add(barcode) {
    _barcodes.push(barcode)
    _save(_barcodes.length-1, barcode)
    _onBarcodesUpdate()
}

function remove(id) {
    _delete(id, _barcodes.splice(id, 1))
    _onBarcodesUpdate()
}

function move(moveAction) {
    _barcodes.splice(moveAction.to, 0, _barcodes.splice(moveAction.from, 1)[0])
    localStorage.replaceObject(sessionStorage.getSession(), _ => _barcodes)
    _onBarcodesUpdate()
}

//
// Barcode local storage
//

function _save(id, barcode) {
    localStorage.replaceObject(sessionStorage.getSession(), old => {
        old = old ?? []
        old[id] = barcode
        return old
    })
}

function _delete(id, barcode) {
    localStorage.replaceObject(sessionStorage.getSession(), old => {
        old.splice(id, 1)
        return old.length > 0 ? old : null
    })
}

//
// Sessions
//

function sessionSelect(select) {
    select.textContent = ''

    const cn = document.createElement('option')
    cn.value = 'create_new'
    cn.innerHTML = 'Create New'
    select.appendChild(cn)

    const curr = document.createElement('option')
    const session = sessionStorage.getSession()
    curr.value = session
    curr.innerHTML = `[CURRENT] ${session}`
    curr.selected = true
    select.appendChild(curr)

    Object.keys(localStorage)?.filter(el => el !== session).forEach((key) => {
        const opt = document.createElement('option')
        opt.value = key
        opt.innerHTML = key
        select.appendChild(opt)
    })
}

function loadSession(session) {
    session = session ?? sessionStorage.newSession()
    sessionStorage.setSession(session)
    _barcodes = localStorage.getObject(session) ?? []
    _onBarcodesUpdate()
}

export { loadSession, sessionSelect, add, remove, move, renderFunc, renderOnto }

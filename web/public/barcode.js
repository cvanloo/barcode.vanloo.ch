import 'api/storage'

let _onBarcodesUpdate = () => {}
let _barcodes = []
let _moveTarget = null;
let _deleteTarget = null;

/* fun = (e) => { ... }
 * e is the function that caused the update (e.g., add, move, remove, or loadSession).
 */
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
    barcode.id = id

    const div = document.createElement('div')
    div.classList.add('barcode')

    const name_tag = document.createElement('p')
    name_tag.innerHTML = barcode.name

    const top_bar = document.createElement('div')
    top_bar.id = 'bc-top-bar'
    top_bar.appendChild(name_tag)
    div.appendChild(top_bar)

    const img = new Image(312, 80)
    img.src = barcode.url
    div.appendChild(img)

    const text_tag = document.createElement('p')
    text_tag.innerHTML = barcode.text
    div.appendChild(text_tag)

    div.onmousedown = (e) => {
        if ([name_tag, text_tag, img].includes(e.target)) return // click must be made on outer div, not its contents
        e.preventDefault()                                       // prevent text selection
        switch (e.buttons) {
            case 1: _moveTarget = barcode; break
            case 4: _deleteTarget = barcode; break
        }
    }
    div.onmouseover = () => {
        if (_moveTarget !== null && _moveTarget !== barcode) {
            move(_moveTarget.id, id)
        }
    }
    div.onmouseup = () => {
        if (_moveTarget !== null && _moveTarget !== barcode) {
            move(_moveTarget.id, id)
        } else if (_deleteTarget === barcode) {
            remove(id)
        }
        _deleteTarget = _moveTarget = null
    }

    return div
}

/* Add a barcode
 * barcode = { name; url; text; }
 */
function add(barcode) {
    _barcodes.push(barcode)
    localStorage.setObject(sessionStorage.getSession(), _barcodes)
    _onBarcodesUpdate(add)
}

function remove(id) {
    _barcodes.splice(id, 1)
    localStorage.setObject(sessionStorage.getSession(), _barcodes)
    _onBarcodesUpdate(remove)
}

function move(from, to) {
    _barcodes.splice(to, 0, _barcodes.splice(from, 1)[0])
    localStorage.setObject(sessionStorage.getSession(), _barcodes)
    _onBarcodesUpdate(move)
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
    _onBarcodesUpdate(loadSession)
}

export { loadSession, sessionSelect, add, remove, move, renderFunc, renderOnto }

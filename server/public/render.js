const COLORS = [
    "#00D8C5", "#00C389", "#3DBE47", "#89C94A", "#D3D848",
    "#AA00FF", "#7750FF", "#4D81FF", "#2D9BF7", "#00BFFF",
    "#FFD03A", "#FFB62B", "#FF9436", "#FF7747", "#FF5C5C",
    "#FF3E3E", "#E63E3E", "#DC143C", "#FF6B8B", "#FF5189",
    "#FF3877", "#FF2E63", "#FF5733", "#FF8000", "#FFA000",
    "#FFC100", "#FFD500", "#FFEA00", "#FFFA00", "#FFFF00",
    "#E1E100", "#B7E200", "#89C64A", "#64C7A9", "#5ED4B3",
    "#53D1C9", "#59BDE5", "#4F9EE6", "#4E79B8", "#6C7A89",
    "#6D6D6D", "#4F4F4F", "#363636", "#FF6E40", "#E65100",
    "#FF5733", "#FF8A65", "#FFBFA1", "#FFDAB9", "#FFB6C1"
];

const getTheme = ()=>{
    if(!window.PATTERNS) return {}
    const classes = window.PATTERNS.map(p=>p.Class)
    classes.sort()
    const theme = {}
    classes.map((cls,i)=>{
        theme[cls] = COLORS[i%COLORS.length]
    })
    return theme
}

const codePreviewElement = document.getElementById("preview")
const outputPreviewElement = document.getElementById("output")
const theme = getTheme()
const appendSpan = (content="",cls="")=>{
    const el = document.createElement("span")
    el.innerText = content
    if(cls){
        el.style.color = theme[cls]
        el.className = cls
    }
    outputPreviewElement.append(el)
}

const render = ()=>{
    if(!window.TOKENS || !window.CODE || !window.PATTERNS) return;

    const code = window.CODE
    codePreviewElement.innerText = code

    if(!window.TOKENS.length){
        outputPreviewElement.innerText = code
        return;
    }

    let prevEnd = 0;
    window.TOKENS.forEach(token=>{
        if(token.Start - prevEnd>0){
            appendSpan(code.substring(prevEnd,token.Start),null)
        }
        appendSpan(code.substring(token.Start,token.End),token.Class)
        prevEnd = token.End
    })

    if(prevEnd !==code.length){
        appendSpan(code.substring(prevEnd,code.length),null)
    }
}

render()
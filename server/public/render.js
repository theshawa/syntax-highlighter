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

const render = ()=>{
    if(!window.TOKENS || !window.CODE) return;
    const code = window.CODE
    const renderElement = document.getElementById("output")
    const previewElement = document.getElementById("preview")
    previewElement.innerText = code

    if(!window.TOKENS.length){
        renderElement.innerText = code
        return;
    }
    const parts = []
    let prevEnd = 0;
    const classNames = []
    const theme = {}
    let colorIndex = -1
    window.TOKENS.forEach(token=>{
        if(token.Start - prevEnd>0){
            parts.push({
                content:code.substring(prevEnd,token.Start),
                highlighted:null
            })
        }
        parts.push({
            content:code.substring(token.Start,token.End),
            highlighted:token.Type
        })
        if(!Object.keys(theme).includes(token.Type)){
            theme[token.Type] = COLORS[(colorIndex+1)%COLORS.length]
            colorIndex++
        }
        prevEnd = token.End
    })
    console.log(classNames)
    parts.forEach(part=>{
        const el = document.createElement("span")

        if(part.highlighted){
            el.style.color = theme[part.highlighted]
        }
        el.innerText = part.content
        renderElement.append(el)
    })

}

render()
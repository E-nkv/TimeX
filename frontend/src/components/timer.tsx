export function Timer({secs}: {secs: number}) {
    const ft = formatSecs(secs)
    return <div className="text-8xl">
        {ft}
    </div>
}

function formatSecs(secs: number) {
    let m = Math.floor(secs / 60)
    let s = secs % 60
    
    
    let fs = String(s).length == 1? "0" + String(s) : String(s) 
    let fm = String(m).length == 1? "0" + String(m) : String(m) 
    
    return `${fm}:${fs}`
}

